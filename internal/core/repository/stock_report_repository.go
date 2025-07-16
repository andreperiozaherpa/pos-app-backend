package repository

import (
	"context"
	"time"

	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

type StockReportRepository interface {
	Create(ctx context.Context, report *models.StockReport) error
	GetByID(ctx context.Context, id uuid.UUID) (*models.StockReport, error)
	ListByStore(ctx context.Context, storeID uuid.UUID, fromDate, toDate time.Time) ([]*models.StockReport, error)
}
