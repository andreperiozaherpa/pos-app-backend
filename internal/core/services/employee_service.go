package services

import (
	"context"
	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

// EmployeeService mendefinisikan kontrak use case terkait karyawan (employee).
type EmployeeService interface {
	// RegisterEmployee membuat employee baru dan relasi ke user.
	RegisterEmployee(ctx context.Context, employee *models.Employee) (uuid.UUID, error)

	// GetEmployeeByID mengambil data employee berdasarkan user ID.
	GetEmployeeByID(ctx context.Context, userID uuid.UUID) (*models.Employee, error)

	// UpdateEmployee memperbarui data employee.
	UpdateEmployee(ctx context.Context, employee *models.Employee) error

	// DeleteEmployee menghapus data employee (soft delete recommended).
	DeleteEmployee(ctx context.Context, userID uuid.UUID) error

	// ListEmployeesByCompanyID mengambil daftar employee berdasarkan company ID.
	ListEmployeesByCompanyID(ctx context.Context, companyID uuid.UUID) ([]*models.Employee, error)

	// ListEmployeesByStoreID mengambil daftar employee berdasarkan store ID.
	ListEmployeesByStoreID(ctx context.Context, storeID uuid.UUID) ([]*models.Employee, error)
}
