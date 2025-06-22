package services

import (
	"context"
	"database/sql"
	"pos-app/backend/internal/data/postgres"
	"pos-app/backend/internal/models"
	"pos-app/backend/internal/utils"
	"time"

	"github.com/google/uuid"
)

// UserService mendefinisikan interface untuk logika bisnis terkait manajemen pengguna.
type UserService interface {
	RegisterUser(ctx context.Context, user *models.User, password string) error
	GetUserProfile(ctx context.Context, userID uuid.UUID) (*models.User, error)
	UpdateUserProfile(ctx context.Context, user *models.User) error
}

// userService adalah implementasi dari UserService.
type userService struct {
	userRepo postgres.UserRepository
}

// NewUserService adalah constructor untuk membuat instance baru dari userService.
func NewUserService(userRepo postgres.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

// RegisterUser mendaftarkan pengguna baru.
func (s *userService) RegisterUser(ctx context.Context, user *models.User, password string) error {
	// Validasi apakah username sudah ada
	_, err := s.userRepo.GetByUsername(ctx, user.Username.String)
	if err == nil {
		return ErrUsernameExists
	}
	if err != sql.ErrNoRows {
		return err // Error lain dari database
	}

	// Validasi apakah email sudah ada (jika email disediakan)
	if user.Email.Valid {
		_, err = s.userRepo.GetByEmail(ctx, user.Email.String)
		if err == nil {
			return ErrEmailExists
		}
		if err != sql.ErrNoRows {
			return err // Error lain dari database
		}
	}

	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	user.ID = uuid.New()
	user.PasswordHash = sql.NullString{String: hashedPassword, Valid: true}
	user.IsActive = true
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	return s.userRepo.Create(ctx, user)
}

// GetUserProfile mengambil profil pengguna berdasarkan ID.
func (s *userService) GetUserProfile(ctx context.Context, userID uuid.UUID) (*models.User, error) {
	user, err := s.userRepo.GetByID(ctx, userID)
	if err == sql.ErrNoRows {
		return nil, ErrUserNotFound
	}
	return user, err
}

// UpdateUserProfile memperbarui profil pengguna.
func (s *userService) UpdateUserProfile(ctx context.Context, user *models.User) error {
	user.UpdatedAt = time.Now()
	err := s.userRepo.Update(ctx, user)
	if err == sql.ErrNoRows {
		return ErrUserNotFound
	}
	return err
}
