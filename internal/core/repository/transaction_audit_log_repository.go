package repository

import (
	"context"
	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

type TransactionAuditLogRepository interface {
	Create(ctx context.Context, log *models.TransactionAuditLog) error
	ListByTransactionID(ctx context.Context, transactionID uuid.UUID) ([]*models.TransactionAuditLog, error)
}
