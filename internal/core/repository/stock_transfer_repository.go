package repository

import (
	"context"
	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

type StockTransferRepository interface {
	// StockTransferRepository defines the interface for stock transfer operations.
	// Create adds a new stock transfer record.
	Create(ctx context.Context, transfer *models.InternalStockTransfer) error
	// Create adds a new stock transfer record.
	GetByID(ctx context.Context, id uuid.UUID) (*models.InternalStockTransfer, error)
	// Update modifies an existing stock transfer record.
	Update(ctx context.Context, transfer *models.InternalStockTransfer) error
	// Delete removes a stock transfer record by its ID.
	Delete(ctx context.Context, id uuid.UUID) error
	// ListByStore retrieves stock transfers for a specific store with pagination.
	ListByStore(ctx context.Context, storeID uuid.UUID, limit, offset int) ([]*models.InternalStockTransfer, error)
}
