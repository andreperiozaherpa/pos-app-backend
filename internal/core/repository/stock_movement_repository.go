package repository

import (
	"context"
	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

// StockMovementRepository mendefinisikan interface untuk operasi data terkait StockMovement.
type StockMovementRepository interface {
	// Create membuat data stock movement baru.
	Create(ctx context.Context, movement *models.StockMovement) error

	// ListByStoreProduct mengambil daftar stock movement berdasarkan store product ID.
	ListByStoreProduct(ctx context.Context, storeProductID uuid.UUID) ([]*models.StockMovement, error)
}
