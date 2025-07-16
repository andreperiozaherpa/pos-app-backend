package repository

import (
	"context"
	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

// PurchaseOrderItemRepository mendefinisikan interface untuk operasi data terkait PurchaseOrderItem.
type PurchaseOrderItemRepository interface {
	// Create membuat data purchase order item baru.
	Create(ctx context.Context, item *models.PurchaseOrderItem) error

	// GetByID mengambil data purchase order item berdasarkan ID.
	GetByID(ctx context.Context, id uuid.UUID) (*models.PurchaseOrderItem, error)

	// ListByPurchaseOrderID mengambil daftar item berdasarkan purchase order ID.
	ListByPurchaseOrderID(ctx context.Context, purchaseOrderID uuid.UUID) ([]*models.PurchaseOrderItem, error)

	// Update memperbarui data purchase order item.
	Update(ctx context.Context, item *models.PurchaseOrderItem) error

	// Delete menghapus data purchase order item berdasarkan ID.
	Delete(ctx context.Context, id uuid.UUID) error
}
