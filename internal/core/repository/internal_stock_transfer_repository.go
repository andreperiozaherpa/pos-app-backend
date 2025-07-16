package repository

import (
	"context"
	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

// InternalStockTransferRepository mendefinisikan interface untuk operasi data terkait InternalStockTransfer.
type InternalStockTransferRepository interface {
	// Create membuat data internal stock transfer baru.
	Create(ctx context.Context, ist *models.InternalStockTransfer) error

	// GetByID mengambil data internal stock transfer berdasarkan ID.
	GetByID(ctx context.Context, id uuid.UUID) (*models.InternalStockTransfer, error)

	// ListByCompanyID mengambil daftar internal stock transfer berdasarkan company ID.
	ListByCompanyID(ctx context.Context, companyID uuid.UUID) ([]*models.InternalStockTransfer, error)

	// Update memperbarui data internal stock transfer.
	Update(ctx context.Context, ist *models.InternalStockTransfer) error

	// Delete menghapus data internal stock transfer berdasarkan ID.
	Delete(ctx context.Context, id uuid.UUID) error
}
