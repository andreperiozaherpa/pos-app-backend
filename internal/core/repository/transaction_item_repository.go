package repository

import (
	"context"
	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

// TransactionItemRepository mendefinisikan interface untuk operasi data terkait TransactionItem.
type TransactionItemRepository interface {
	// Create membuat data transaction item baru.
	Create(ctx context.Context, item *models.TransactionItem) error

	// GetByID mengambil data transaction item berdasarkan ID.
	GetByID(ctx context.Context, id uuid.UUID) (*models.TransactionItem, error)

	// ListByTransactionID mengambil daftar item berdasarkan transaction ID.
	ListByTransactionID(ctx context.Context, transactionID uuid.UUID) ([]*models.TransactionItem, error)

	// Update memperbarui data transaction item.
	Update(ctx context.Context, item *models.TransactionItem) error

	// Delete menghapus data transaction item berdasarkan ID.
	Delete(ctx context.Context, id uuid.UUID) error
}
