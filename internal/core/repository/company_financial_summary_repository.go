package repository

import (
	"context"

	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

// CompanyFinancialSummaryRepository adalah interface untuk operasi CRUD dan query pada ringkasan keuangan perusahaan.
type CompanyFinancialSummaryRepository interface {
	// Create menambahkan data ringkasan keuangan baru ke dalam database.
	Create(ctx context.Context, summary *models.CompanyFinancialSummary) error

	// GetByCompanyID mengambil ringkasan keuangan berdasarkan ID perusahaan.
	GetByCompanyID(ctx context.Context, companyID uuid.UUID) (*models.CompanyFinancialSummary, error)

	// Update memperbarui data ringkasan keuangan yang sudah ada.
	Update(ctx context.Context, summary *models.CompanyFinancialSummary) error

	// Delete menghapus data ringkasan keuangan berdasarkan ID.
	Delete(ctx context.Context, id uuid.UUID) error
}
