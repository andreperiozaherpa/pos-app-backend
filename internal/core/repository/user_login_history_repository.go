package repository

import (
	"context"
	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

type UserLoginHistoryRepository interface {
	Create(ctx context.Context, history *models.UserLoginHistory) error
	GetByID(ctx context.Context, id uuid.UUID) (*models.UserLoginHistory, error)
	ListByUserID(ctx context.Context, userID uuid.UUID) ([]*models.UserLoginHistory, error)
}
