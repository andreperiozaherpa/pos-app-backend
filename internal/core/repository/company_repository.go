package repository

import (
	"context"
	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

// CompanyRepository mendefinisikan interface untuk operasi data terkait Company.
type CompanyRepository interface {
	// Create membuat data company baru.
	Create(ctx context.Context, company *models.Company) error

	// GetByID mengambil data company berdasarkan ID.
	GetByID(ctx context.Context, id uuid.UUID) (*models.Company, error)

	// Update memperbarui data company.
	Update(ctx context.Context, company *models.Company) error

	// Delete menghapus data company berdasarkan ID.
	Delete(ctx context.Context, id uuid.UUID) error

	// ListAll mengambil semua data company.
	ListAll(ctx context.Context) ([]*models.Company, error)
}
