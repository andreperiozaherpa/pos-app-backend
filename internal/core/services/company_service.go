package services

import (
	"context"
	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

// CompanyService menangani manajemen data perusahaan/tenant.
type CompanyService interface {
	// CRUD utama
	CreateCompany(ctx context.Context, company *models.Company) (uuid.UUID, error)
	GetCompanyByID(ctx context.Context, id uuid.UUID) (*models.Company, error)
	UpdateCompany(ctx context.Context, company *models.Company) error
	DeleteCompany(ctx context.Context, id uuid.UUID) error
	ListAllCompanies(ctx context.Context) ([]*models.Company, error)

	// Opsional & Custom
	SearchCompanies(ctx context.Context, query string) ([]*models.Company, error)
	GetCompanyFinancialSummary(ctx context.Context, companyID uuid.UUID) (*models.CompanyFinancialSummary, error)
	ExportCompanies(ctx context.Context) ([]byte, error)
	ArchiveCompany(ctx context.Context, id uuid.UUID) error
	RestoreCompany(ctx context.Context, id uuid.UUID) error
	ListCompanyStores(ctx context.Context, companyID uuid.UUID) ([]*models.Store, error)
}
