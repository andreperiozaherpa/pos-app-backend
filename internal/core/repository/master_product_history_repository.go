package repository

import (
	"context"

	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

// MasterProductHistoryRepository adalah interface untuk operasi pada
// histori perubahan master product.
type MasterProductHistoryRepository interface {
	Create(ctx context.Context, history *models.MasterProductHistory) error
	ListByMasterProduct(ctx context.Context, masterProductID uuid.UUID) ([]*models.MasterProductHistory, error)
}
