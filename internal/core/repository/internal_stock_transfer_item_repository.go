package repository

import (
	"context"
	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

// InternalStockTransferItemRepository mendefinisikan interface untuk operasi data terkait InternalStockTransferItem.
type InternalStockTransferItemRepository interface {
	// Create membuat data internal stock transfer item baru.
	Create(ctx context.Context, item *models.InternalStockTransferItem) error

	// GetByID mengambil data internal stock transfer item berdasarkan ID.
	GetByID(ctx context.Context, id uuid.UUID) (*models.InternalStockTransferItem, error)

	// ListByInternalStockTransferID mengambil daftar item berdasarkan internal stock transfer ID.
	ListByInternalStockTransferID(ctx context.Context, internalStockTransferID uuid.UUID) ([]*models.InternalStockTransferItem, error)

	// Update memperbarui data internal stock transfer item.
	Update(ctx context.Context, item *models.InternalStockTransferItem) error

	// Delete menghapus data internal stock transfer item berdasarkan ID.
	Delete(ctx context.Context, id uuid.UUID) error
}
