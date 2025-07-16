package repository

import (
	"context"
	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

// TransactionRepository mendefinisikan interface untuk operasi data terkait Transaction.
type TransactionRepository interface {
	// Create membuat data transaksi baru.
	Create(ctx context.Context, transaction *models.Transaction) error

	// GetByID mengambil data transaksi berdasarkan ID.
	GetByID(ctx context.Context, id uuid.UUID) (*models.Transaction, error)

	// GetByTransactionCode mengambil data transaksi berdasarkan kode transaksi.
	GetByTransactionCode(ctx context.Context, code string) (*models.Transaction, error)

	// ListByStoreID mengambil daftar transaksi berdasarkan store ID.
	ListByStoreID(ctx context.Context, storeID uuid.UUID) ([]*models.Transaction, error)
}
