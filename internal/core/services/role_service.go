package services

import (
	"context"
	"pos-app/backend/internal/models"
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

	// AssignRoleToEmployee menetapkan role ke employee (opsional, untuk multi-role).
	AssignRoleToEmployee(ctx context.Context, employeeUserID string, roleID int64) error
}
