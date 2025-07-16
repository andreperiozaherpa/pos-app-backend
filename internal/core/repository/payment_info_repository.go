package repository

import (
	"context"

	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

// PaymentInfoRepository adalah interface untuk operasi CRUD dan query
// terkait data pembayaran transaksi.
type PaymentInfoRepository interface {
	Create(ctx context.Context, payment *models.PaymentInfo) error
	GetByID(ctx context.Context, id uuid.UUID) (*models.PaymentInfo, error)
	ListByTransactionID(ctx context.Context, transactionID uuid.UUID) ([]*models.PaymentInfo, error)
}
