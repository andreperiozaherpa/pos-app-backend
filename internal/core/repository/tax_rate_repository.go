package repository

import (
	"context"
	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

// TaxRateRepository mendefinisikan interface untuk operasi data terkait TaxRate.
type TaxRateRepository interface {
	// Create membuat data tax rate baru.
	Create(ctx context.Context, taxRate *models.TaxRate) error

	// GetByID mengambil data tax rate berdasarkan ID.
	GetByID(ctx context.Context, id int32) (*models.TaxRate, error)

	// ListByCompanyID mengambil daftar tax rate berdasarkan company ID.
	ListByCompanyID(ctx context.Context, companyID uuid.UUID) ([]*models.TaxRate, error)

	// Update memperbarui data tax rate.
	Update(ctx context.Context, taxRate *models.TaxRate) error

	// Delete menghapus data tax rate berdasarkan ID.
	Delete(ctx context.Context, id int32) error
}
