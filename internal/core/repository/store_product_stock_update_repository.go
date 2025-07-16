package repository

import (
	"context"
	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

type StoreProductStockUpdateRepository interface {
	Create(ctx context.Context, update *models.StoreProductStockUpdate) error
	GetByID(ctx context.Context, id uuid.UUID) (*models.StoreProductStockUpdate, error)
	ListByStoreProductID(ctx context.Context, storeProductID uuid.UUID) ([]*models.StoreProductStockUpdate, error)
}
