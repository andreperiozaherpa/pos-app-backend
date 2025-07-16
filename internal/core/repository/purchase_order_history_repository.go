package repository

import (
	"context"

	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

// PurchaseOrderHistoryRepository adalah interface untuk operasi CRUD
// pada histori purchase order.
type PurchaseOrderHistoryRepository interface {
	Create(ctx context.Context, history *models.PurchaseOrderHistory) error
	ListByPurchaseOrder(ctx context.Context, purchaseOrderID uuid.UUID) ([]*models.PurchaseOrderHistory, error)
}
