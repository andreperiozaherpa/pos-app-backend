package repository

import (
	"context"
	"time"

	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

type CustomerActivityReportRepository interface {
	Create(ctx context.Context, report *models.CustomerActivityReport) error
	GetByID(ctx context.Context, id uuid.UUID) (*models.CustomerActivityReport, error)
	ListByCustomer(ctx context.Context, customerUserID uuid.UUID, fromDate, toDate time.Time) ([]*models.CustomerActivityReport, error)
	ListByCompany(ctx context.Context, companyID uuid.UUID, fromDate, toDate time.Time) ([]*models.CustomerActivityReport, error)
}
