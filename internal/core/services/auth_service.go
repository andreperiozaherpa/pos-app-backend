package services

import (
	"context"
	"pos-app/backend/internal/models"
)

// AuthService menangani autentikasi user secara terpisah (JWT, sesi, refresh token, dll).
type AuthService interface {
	Login(ctx context.Context, username, password string) (*models.AuthSession, error)
	Logout(ctx context.Context, sessionID string) error
	RefreshToken(ctx context.Context, refreshToken string) (*models.AuthSession, error)
	ValidateToken(ctx context.Context, token string) (bool, error)
	ChangePassword(ctx context.Context, userID string, oldPassword, newPassword string) error
	SendPasswordResetLink(ctx context.Context, email string) error
	ValidatePasswordResetToken(ctx context.Context, token string) (bool, error)
	GetAuthSessionInfo(ctx context.Context, sessionID string) (*models.AuthSession, error)
}
