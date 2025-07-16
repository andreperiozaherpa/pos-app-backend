package repository

import (
	"context"
	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

// AppliedItemDiscountRepository mendefinisikan interface untuk operasi data applied item discount (pivot).
type AppliedItemDiscountRepository interface {
	// Create menambahkan applied item discount baru.
	Create(ctx context.Context, applied *models.AppliedItemDiscount) error

	// Delete menghapus applied item discount berdasarkan composite key.
	Delete(ctx context.Context, transactionItemID uuid.UUID, discountID uuid.UUID) error

	// ListByTransactionItemID mengambil daftar applied item discount berdasarkan transaction item ID.
	ListByTransactionItemID(ctx context.Context, transactionItemID uuid.UUID) ([]*models.AppliedItemDiscount, error)
}
