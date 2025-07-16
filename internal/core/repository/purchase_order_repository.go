package repository

import (
	"context"
	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

// PurchaseOrderRepository mendefinisikan interface untuk operasi data terkait PurchaseOrder.
type PurchaseOrderRepository interface {
	// Create membuat data purchase order baru.
	Create(ctx context.Context, po *models.PurchaseOrder) error

	// GetByID mengambil data purchase order berdasarkan ID.
	GetByID(ctx context.Context, id uuid.UUID) (*models.PurchaseOrder, error)

	// ListByStoreID mengambil daftar purchase order berdasarkan store ID.
	ListByStoreID(ctx context.Context, storeID uuid.UUID) ([]*models.PurchaseOrder, error)

	// Update memperbarui data purchase order.
	Update(ctx context.Context, po *models.PurchaseOrder) error

	// Delete menghapus data purchase order berdasarkan ID.
	Delete(ctx context.Context, id uuid.UUID) error
}
