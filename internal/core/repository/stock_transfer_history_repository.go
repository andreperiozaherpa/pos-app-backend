package repository

import (
	"context"

	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

type StockTransferHistoryRepository interface {
	Create(ctx context.Context, history *models.StockTransferHistory) error
	GetByID(ctx context.Context, id uuid.UUID) (*models.StockTransferHistory, error)
	ListByStockTransfer(ctx context.Context, stockTransferID uuid.UUID) ([]*models.StockTransferHistory, error)
}
