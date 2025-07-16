package repository

import (
	"context"
	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

// UserRepository adalah kontrak akses data untuk entitas User.
type UserRepository interface {
	// Create membuat pengguna baru di database.
	Create(ctx context.Context, user *models.User) error

	// GetByID mengambil pengguna berdasarkan ID.
	GetByID(ctx context.Context, id uuid.UUID) (*models.User, error)

	// GetByUsername mengambil pengguna berdasarkan username.
	GetByUsername(ctx context.Context, username string) (*models.User, error)

	// GetByEmail mengambil pengguna berdasarkan email.
	GetByEmail(ctx context.Context, email string) (*models.User, error)

	// GetByPhoneNumber mengambil pengguna berdasarkan nomor telepon.
	GetByPhoneNumber(ctx context.Context, phoneNumber string) (*models.User, error)

	// Update memperbarui data pengguna yang ada.
	Update(ctx context.Context, user *models.User) error

	// Delete menghapus pengguna berdasarkan ID (hard delete atau soft delete).
	Delete(ctx context.Context, id uuid.UUID) error
}
