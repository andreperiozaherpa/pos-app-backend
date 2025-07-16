package services

import (
	"context"
	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

// BusinessLineService menangani manajemen lini usaha dalam satu perusahaan/tenant.
type BusinessLineService interface {
	// CRUD utama
	CreateBusinessLine(ctx context.Context, businessLine *models.BusinessLine) (uuid.UUID, error)
	GetBusinessLineByID(ctx context.Context, id uuid.UUID) (*models.BusinessLine, error)
	UpdateBusinessLine(ctx context.Context, businessLine *models.BusinessLine) error
	DeleteBusinessLine(ctx context.Context, id uuid.UUID) error
	ListBusinessLinesByCompanyID(ctx context.Context, companyID uuid.UUID) ([]*models.BusinessLine, error)

	// Opsional & Custom
	SearchBusinessLines(ctx context.Context, companyID uuid.UUID, query string) ([]*models.BusinessLine, error)
	ExportBusinessLines(ctx context.Context, companyID uuid.UUID) ([]byte, error)
	ArchiveBusinessLine(ctx context.Context, id uuid.UUID) error
	RestoreBusinessLine(ctx context.Context, id uuid.UUID) error
}
