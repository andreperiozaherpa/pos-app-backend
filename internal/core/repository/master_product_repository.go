package repository

import (
	"context"
	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

// MasterProductRepository mendefinisikan interface untuk operasi data terkait MasterProduct.
type MasterProductRepository interface {
	// Create membuat data master product baru.
	Create(ctx context.Context, mp *models.MasterProduct) error

	// GetByID mengambil data master product berdasarkan ID.
	GetByID(ctx context.Context, id uuid.UUID) (*models.MasterProduct, error)

	// GetByCompanyAndCode mengambil data master product berdasarkan company ID dan kode produk master.
	GetByCompanyAndCode(ctx context.Context, companyID uuid.UUID, code string) (*models.MasterProduct, error)

	// ListByCompanyID mengambil daftar master product berdasarkan ID perusahaan.
	ListByCompanyID(ctx context.Context, companyID uuid.UUID) ([]*models.MasterProduct, error)

	// Update memperbarui data master product.
	Update(ctx context.Context, mp *models.MasterProduct) error

	// Delete menghapus data master product berdasarkan ID.
	Delete(ctx context.Context, id uuid.UUID) error
}
