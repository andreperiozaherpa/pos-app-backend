package repository

import (
	"context"
	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

// OperationalExpenseRepository mendefinisikan interface untuk operasi data terkait OperationalExpense.
type OperationalExpenseRepository interface {
	// Create membuat data operational expense baru.
	Create(ctx context.Context, expense *models.OperationalExpense) error

	// GetByID mengambil data operational expense berdasarkan ID.
	GetByID(ctx context.Context, id uuid.UUID) (*models.OperationalExpense, error)

	// ListByCompanyID mengambil daftar operational expense berdasarkan company ID.
	ListByCompanyID(ctx context.Context, companyID uuid.UUID) ([]*models.OperationalExpense, error)

	// ListByStoreID mengambil daftar operational expense berdasarkan store ID.
	ListByStoreID(ctx context.Context, storeID uuid.UUID) ([]*models.OperationalExpense, error)

	// Update memperbarui data operational expense.
	Update(ctx context.Context, expense *models.OperationalExpense) error

	// Delete menghapus data operational expense berdasarkan ID.
	Delete(ctx context.Context, id uuid.UUID) error
}
