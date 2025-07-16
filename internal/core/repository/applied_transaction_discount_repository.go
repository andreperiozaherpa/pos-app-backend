package repository

import (
	"context"
	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

// AppliedTransactionDiscountRepository mendefinisikan interface untuk operasi data terkait AppliedTransactionDiscount.
type AppliedTransactionDiscountRepository interface {
	// Create menambahkan satu atau lebih applied transaction discount ke database.
	Create(ctx context.Context, transactionDiscounts []models.AppliedTransactionDiscount) error

	// ListByTransactionID mengambil daftar applied transaction discount berdasarkan transaction ID.
	ListByTransactionID(ctx context.Context, transactionID uuid.UUID) ([]*models.AppliedTransactionDiscount, error)

	// DeleteByTransactionID menghapus semua applied transaction discount berdasarkan transaction ID.
	DeleteByTransactionID(ctx context.Context, transactionID uuid.UUID) error
}
