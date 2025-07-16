package repository

import (
	"context"
	"pos-app/backend/internal/models"
	"time"

	"github.com/google/uuid"
)

type TransactionSummaryRepository interface {
	Create(ctx context.Context, summary *models.TransactionSummary) error
	GetByID(ctx context.Context, id uuid.UUID) (*models.TransactionSummary, error)
	ListByStoreAndPeriod(ctx context.Context, storeID uuid.UUID, fromDate, toDate time.Time) ([]*models.TransactionSummary, error)
}
