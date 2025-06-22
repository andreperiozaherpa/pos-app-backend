package services

import (
	"context"
	"database/sql"
	"os"
	"pos-app/backend/internal/data/postgres"
	"pos-app/backend/internal/models"
	"pos-app/backend/internal/utils"
	"time"
)

// AuthService mendefinisikan interface untuk logika bisnis terkait autentikasi.
type AuthService interface {
	Login(ctx context.Context, username, password string) (*models.User, string, error)
	ValidateToken(ctx context.Context, tokenString string) (*models.User, error)
}

// authService adalah implementasi dari AuthService.
type authService struct {
	userRepo      postgres.UserRepository
	jwtSecret     string
	jwtExpiration time.Duration
}

// NewAuthService adalah constructor untuk membuat instance baru dari authService.
func NewAuthService(userRepo postgres.UserRepository) AuthService {
	secret := os.Getenv("JWT_SECRET_KEY")
	if secret == "" {
		secret = "default_super_secret_key" // Fallback untuk development
	}

	// Default expiration 24 jam
	expiration := 24 * time.Hour

	return &authService{
		userRepo:      userRepo,
		jwtSecret:     secret,
		jwtExpiration: expiration,
	}
}

// Login mengautentikasi pengguna dan mengembalikan token jika berhasil.
func (s *authService) Login(ctx context.Context, username, password string) (*models.User, string, error) {
	user, err := s.userRepo.GetByUsername(ctx, username)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, "", ErrInvalidCredentials
		}
		return nil, "", err // Error database lainnya
	}

	if !user.IsActive {
		return nil, "", ErrUserInactive
	}

	if !utils.CheckPasswordHash(password, user.PasswordHash.String) {
		return nil, "", ErrInvalidCredentials
	}

	token, err := utils.GenerateToken(user.ID, s.jwtSecret, s.jwtExpiration)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}

// ValidateToken memvalidasi token dan mengembalikan data pengguna jika valid.
func (s *authService) ValidateToken(ctx context.Context, tokenString string) (*models.User, error) {
	userID, err := utils.ValidateToken(tokenString, s.jwtSecret)
	if err != nil {
		return nil, err
	}

	return s.userRepo.GetByID(ctx, userID)
}
