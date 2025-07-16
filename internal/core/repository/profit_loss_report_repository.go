package repository

import (
	"context"
	"time"

	"pos-app/backend/internal/models"
)

// ProfitLossReportRepository mendefinisikan kontrak operasi
// CRUD untuk entitas ProfitLossReport.
type ProfitLossReportRepository interface {
	Create(ctx context.Context, report *models.ProfitLossReport) error
	GetByID(ctx context.Context, id string) (*models.ProfitLossReport, error)
	ListByCompany(ctx context.Context, companyID string, fromDate, toDate time.Time) ([]*models.ProfitLossReport, error)
}
