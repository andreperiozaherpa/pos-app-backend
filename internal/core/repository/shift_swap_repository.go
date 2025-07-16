package repository

import (
	"context"
	"time"

	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

// ShiftSwapRepository mendefinisikan kontrak operasi CRUD dan query
// untuk entitas ShiftSwap.
type ShiftSwapRepository interface {
	Create(ctx context.Context, swap *models.ShiftSwap) error
	GetByID(ctx context.Context, id uuid.UUID) (*models.ShiftSwap, error)
	ListByEmployee(ctx context.Context, employeeUserID uuid.UUID) ([]*models.ShiftSwap, error)
	ListByStatus(ctx context.Context, status string) ([]*models.ShiftSwap, error)
	UpdateStatus(ctx context.Context, id uuid.UUID, status string, approvedBy *uuid.UUID, approvedAt *time.Time) error
}
