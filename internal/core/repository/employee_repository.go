package repository

import (
	"context"
	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

// EmployeeRepository adalah kontrak akses data untuk entitas Employee.
type EmployeeRepository interface {
	// Create membuat data karyawan baru.
	Create(ctx context.Context, employee *models.Employee) error

	// GetByID mengambil data karyawan berdasarkan user ID.
	GetByID(ctx context.Context, userID uuid.UUID) (*models.Employee, error)

	// GetByEmployeeIDNumber mengambil data karyawan berdasarkan nomor identitas karyawan.
	GetByEmployeeIDNumber(ctx context.Context, employeeIDNumber string) (*models.Employee, error)

	// Update memperbarui data karyawan.
	Update(ctx context.Context, employee *models.Employee) error

	// Delete menghapus data karyawan berdasarkan user ID (hard atau soft delete).
	Delete(ctx context.Context, userID uuid.UUID) error

	// ListByCompanyID mengambil daftar employee berdasarkan companyID dengan pagination.
	ListByCompanyID(ctx context.Context, companyID uuid.UUID, limit, offset int) ([]*models.Employee, error)

	// ListByStoreID mengambil semua karyawan pada toko tertentu.
	ListByStoreID(ctx context.Context, storeID uuid.UUID) ([]*models.Employee, error)
}
