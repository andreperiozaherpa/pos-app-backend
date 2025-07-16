package repository

import (
	"context"
	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

// EmployeeRoleRepository mendefinisikan interface untuk operasi data terkait EmployeeRole.
type EmployeeRoleRepository interface {
	// AssignRoleToEmployee menetapkan role kepada karyawan.
	AssignRoleToEmployee(ctx context.Context, employeeUserID uuid.UUID, roleID int32) error

	// RemoveRoleFromEmployee menghapus role dari karyawan.
	RemoveRoleFromEmployee(ctx context.Context, employeeUserID uuid.UUID, roleID int32) error

	// GetRolesForEmployee mengambil daftar role yang dimiliki oleh karyawan.
	GetRolesForEmployee(ctx context.Context, employeeUserID uuid.UUID) ([]*models.Role, error)
}
