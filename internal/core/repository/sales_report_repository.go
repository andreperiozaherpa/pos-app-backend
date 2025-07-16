package repository

import (
	"context"
	"time"

	"pos-app/backend/internal/models"
)

// SalesReportRepository adalah interface untuk operasi CRUD
// pada laporan penjualan.
type SalesReportRepository interface {
	Create(ctx context.Context, report *models.SalesReport) error
	GetByID(ctx context.Context, id string) (*models.SalesReport, error)
	ListByStore(ctx context.Context, storeID string, fromDate, toDate time.Time) ([]*models.SalesReport, error)
}
