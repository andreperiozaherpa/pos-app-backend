package services

import (
	"context"
	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

// RoleService mendefinisikan kontrak use case untuk role/level user (kasir, manager, admin, dsb).
type RoleService interface {
	// CreateRole membuat role baru.
	CreateRole(ctx context.Context, role *models.Role) (int64, error)

	// GetRoleByID mengambil role berdasarkan ID.
	GetRoleByID(ctx context.Context, id int64) (*models.Role, error)

	// UpdateRole memperbarui data role.
	UpdateRole(ctx context.Context, role *models.Role) error

	// DeleteRole menghapus role berdasarkan ID.
	DeleteRole(ctx context.Context, id int64) error

	// ListRoles mengambil semua role yang tersedia.
	ListRoles(ctx context.Context) ([]*models.Role, error)

	// AssignRoleToEmployee memberikan role ke employee.
	AssignRoleToEmployee(ctx context.Context, userID uuid.UUID, roleID int) error

	// ListRolePermissions menampilkan daftar permission yang dimiliki role.
	ListRolePermissions(ctx context.Context, roleID int) ([]*models.Permission, error)

	// RemoveRoleFromEmployee menghapus role tertentu dari employee.
	RemoveRoleFromEmployee(ctx context.Context, userID uuid.UUID, roleID int) error

	// AssignRoleToMultipleEmployees menetapkan role ke banyak employee sekaligus.
	AssignRoleToMultipleEmployees(ctx context.Context, roleID int, userIDs []uuid.UUID) error

	// ListUsersByRole mengambil daftar user yang memiliki role tertentu.
	ListUsersByRole(ctx context.Context, roleID int) ([]*models.User, error)

	// CloneRole menggandakan role beserta seluruh permissionnya.
	CloneRole(ctx context.Context, sourceRoleID int, newRoleName string) (int, error)

	// ExportRoles mengekspor daftar role ke file excel/CSV.
	ExportRoles(ctx context.Context) ([]byte, error)
}
