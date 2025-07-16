package repository

import (
	"context"
	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

type StockTransferItemRepository interface {
	Create(ctx context.Context, item *models.StockTransferItem) error
	GetByID(ctx context.Context, id uuid.UUID) (*models.StockTransferItem, error)
	Update(ctx context.Context, item *models.StockTransferItem) error
	Delete(ctx context.Context, id uuid.UUID) error
	ListByStockTransfer(ctx context.Context, stockTransferID uuid.UUID) ([]*models.StockTransferItem, error)
}
