package repository

import (
	"context"
	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

// BusinessLineRepository mendefinisikan interface untuk operasi data terkait BusinessLine.
type BusinessLineRepository interface {
	// Create membuat data business line baru.
	Create(ctx context.Context, bl *models.BusinessLine) error

	// GetByID mengambil data business line berdasarkan ID.
	GetByID(ctx context.Context, id uuid.UUID) (*models.BusinessLine, error)

	// ListByCompanyID mengambil daftar business line berdasarkan company ID.
	ListByCompanyID(ctx context.Context, companyID uuid.UUID) ([]*models.BusinessLine, error)

	// Update memperbarui data business line.
	Update(ctx context.Context, bl *models.BusinessLine) error

	// Delete menghapus data business line berdasarkan ID.
	Delete(ctx context.Context, id uuid.UUID) error
}
