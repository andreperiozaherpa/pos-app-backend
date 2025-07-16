package repository

import (
	"context"
	"time"

	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

// StockMovementSummaryRepository mendefinisikan kontrak operasi CRUD
// untuk entitas StockMovementSummary.
type StockMovementSummaryRepository interface {
	Create(ctx context.Context, summary *models.StockMovementSummary) error
	GetByID(ctx context.Context, id uuid.UUID) (*models.StockMovementSummary, error)
	ListByStoreProduct(ctx context.Context, storeProductID uuid.UUID, fromDate, toDate time.Time) ([]*models.StockMovementSummary, error)
}
