package services

import (
	"context"
	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

// UserService mendefinisikan kontrak use case terkait user (karyawan & pelanggan/member).
type UserService interface {
	// RegisterUser untuk membuat user baru (employee/customer).
	RegisterUser(ctx context.Context, user *models.User) (uuid.UUID, error)

	// GetUserByID mengambil user berdasarkan ID.
	GetUserByID(ctx context.Context, id uuid.UUID) (*models.User, error)

	// GetUserByUsername mengambil user berdasarkan username.
	GetUserByUsername(ctx context.Context, username string) (*models.User, error)

	// GetUserByEmail mengambil user berdasarkan email.
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)

	// GetUserByPhoneNumber mengambil user berdasarkan nomor telepon.
	GetUserByPhoneNumber(ctx context.Context, phone string) (*models.User, error)

	// UpdateUser memperbarui data user.
	UpdateUser(ctx context.Context, user *models.User) error

	// DeleteUser menghapus user berdasarkan ID.
	DeleteUser(ctx context.Context, id uuid.UUID) error

	// ListUsersByType mengambil semua user berdasarkan tipe (EMPLOYEE/CUSTOMER).
	ListUsersByType(ctx context.Context, userType string) ([]*models.User, error)

	// AuthenticateUser memvalidasi login (username/email/phone + password).
	AuthenticateUser(ctx context.Context, identity string, password string) (*models.User, error)

	// ChangePassword mengganti password user.
	ChangePassword(ctx context.Context, userID uuid.UUID, newPassword string) error

	// ActivateUser & DeactivateUser untuk blokir user tanpa delete.
	ActivateUser(ctx context.Context, userID uuid.UUID) error
	DeactivateUser(ctx context.Context, userID uuid.UUID) error

	// SearchUsers pencarian user fleksibel.
	SearchUsers(ctx context.Context, query string, userType string, limit int) ([]*models.User, error)
}
