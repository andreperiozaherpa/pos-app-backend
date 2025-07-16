package services

import (
	"context"
	"pos-app/backend/internal/models"
	"time"

	"github.com/google/uuid"
)

// ExpenseService menangani manajemen pengeluaran operasional (operational expense).
type ExpenseService interface {
	// CRUD utama
	CreateExpense(ctx context.Context, expense *models.OperationalExpense) (uuid.UUID, error)
	GetExpenseByID(ctx context.Context, id uuid.UUID) (*models.OperationalExpense, error)
	UpdateExpense(ctx context.Context, expense *models.OperationalExpense) error
	DeleteExpense(ctx context.Context, id uuid.UUID) error
	ListExpensesByCompanyID(ctx context.Context, companyID uuid.UUID) ([]*models.OperationalExpense, error)
	ListExpensesByStoreID(ctx context.Context, storeID uuid.UUID) ([]*models.OperationalExpense, error)

	// Opsional & Custom
	ApproveExpense(ctx context.Context, expenseID uuid.UUID, approverID uuid.UUID) error
	GenerateExpenseReport(ctx context.Context, companyID uuid.UUID, from, to time.Time) ([]byte, error)
	ExportExpenses(ctx context.Context, companyID uuid.UUID, from, to time.Time) ([]byte, error)
	ListExpensesByDateRange(ctx context.Context, companyID uuid.UUID, from, to time.Time) ([]*models.OperationalExpense, error)
	ApproveMultipleExpenses(ctx context.Context, expenseIDs []uuid.UUID, approverID uuid.UUID) error
}
