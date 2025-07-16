package services

import (
	"context"
	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

// TaxRateService menangani logika bisnis manajemen pajak (tax rate) per company.
type TaxRateService interface {
	// CRUD Utama
	CreateTaxRate(ctx context.Context, taxRate *models.TaxRate) (int64, error)
	GetTaxRateByID(ctx context.Context, id int64) (*models.TaxRate, error)
	UpdateTaxRate(ctx context.Context, taxRate *models.TaxRate) error
	DeleteTaxRate(ctx context.Context, id int64) error
	ListTaxRatesByCompanyID(ctx context.Context, companyID uuid.UUID) ([]*models.TaxRate, error)

	// Opsional & Custom
	CalculateTaxForTransaction(ctx context.Context, transactionID uuid.UUID) (float64, error)
	ArchiveTaxRate(ctx context.Context, id int64) error
	RestoreTaxRate(ctx context.Context, id int64) error
	ExportTaxRates(ctx context.Context, companyID uuid.UUID) ([]byte, error)
	ListTaxRatesByDateRange(ctx context.Context, companyID uuid.UUID, start, end string) ([]*models.TaxRate, error)
	SetTaxRateActive(ctx context.Context, id int64, active bool) error
}
