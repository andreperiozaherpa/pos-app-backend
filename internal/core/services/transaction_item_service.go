package services

import (
	"context"
	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

// TransactionItemService menangani detail item per transaksi POS.
type TransactionItemService interface {
	// CRUD utama
	CreateTransactionItem(ctx context.Context, item *models.TransactionItem) (uuid.UUID, error)
	GetTransactionItemByID(ctx context.Context, id uuid.UUID) (*models.TransactionItem, error)
	ListTransactionItemsByTransactionID(ctx context.Context, transactionID uuid.UUID) ([]*models.TransactionItem, error)
	UpdateTransactionItem(ctx context.Context, item *models.TransactionItem) error
	DeleteTransactionItem(ctx context.Context, id uuid.UUID) error

	// Opsional/custom
	CalculateItemDiscount(ctx context.Context, itemID uuid.UUID) (float64, error)
	ListTransactionItemsByProductID(ctx context.Context, productID uuid.UUID) ([]*models.TransactionItem, error)
	ExportTransactionItems(ctx context.Context, transactionID uuid.UUID) ([]byte, error)
	BulkUpdateTransactionItems(ctx context.Context, items []*models.TransactionItem) error
}
