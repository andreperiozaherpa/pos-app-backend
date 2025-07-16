package repository

import (
	"context"
	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

// SupplierRepository mendefinisikan interface untuk operasi data terkait Supplier.
type SupplierRepository interface {
	// Create membuat data supplier baru.
	Create(ctx context.Context, supplier *models.Supplier) error

	// GetByID mengambil data supplier berdasarkan ID.
	GetByID(ctx context.Context, id uuid.UUID) (*models.Supplier, error)

	// ListByCompanyID mengambil daftar supplier berdasarkan ID perusahaan.
	ListByCompanyID(ctx context.Context, companyID uuid.UUID, limit, offset int) ([]*models.Supplier, error)

	// Update memperbarui data supplier.
	Update(ctx context.Context, supplier *models.Supplier) error

	// Delete menghapus data supplier berdasarkan ID.
	Delete(ctx context.Context, id uuid.UUID) error
}
