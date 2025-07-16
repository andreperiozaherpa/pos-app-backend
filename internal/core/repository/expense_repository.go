package repository

import (
	"context"

	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

// ExpenseRepository adalah interface kontrak untuk operasi CRUD dan query
// terkait data OperationalExpense.
type ExpenseRepository interface {
	Create(ctx context.Context, expense *models.OperationalExpense) error
	GetByID(ctx context.Context, id uuid.UUID) (*models.OperationalExpense, error)
	Update(ctx context.Context, expense *models.OperationalExpense) error
	Delete(ctx context.Context, id uuid.UUID) error
	ListByCompany(ctx context.Context, companyID uuid.UUID, limit, offset int) ([]*models.OperationalExpense, error)
}
