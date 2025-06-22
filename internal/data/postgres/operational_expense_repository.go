package postgres

import (
	"context"
	"database/sql"
	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

// OperationalExpenseRepository mendefinisikan interface untuk operasi data terkait OperationalExpense.
type OperationalExpenseRepository interface {
	Create(ctx context.Context, expense *models.OperationalExpense) error
	GetByID(ctx context.Context, id uuid.UUID) (*models.OperationalExpense, error)
	ListByCompanyID(ctx context.Context, companyID uuid.UUID) ([]*models.OperationalExpense, error)
	ListByStoreID(ctx context.Context, storeID uuid.UUID) ([]*models.OperationalExpense, error)
	Update(ctx context.Context, expense *models.OperationalExpense) error
	Delete(ctx context.Context, id uuid.UUID) error
}

// pgOperationalExpenseRepository adalah implementasi dari OperationalExpenseRepository untuk PostgreSQL.
type pgOperationalExpenseRepository struct {
	db DBExecutor
}

// NewPgOperationalExpenseRepository adalah constructor untuk membuat instance baru dari pgOperationalExpenseRepository.
func NewPgOperationalExpenseRepository(db DBExecutor) OperationalExpenseRepository {
	return &pgOperationalExpenseRepository{db: db}
}

// Create menyisipkan pengeluaran operasional baru.
func (r *pgOperationalExpenseRepository) Create(ctx context.Context, oe *models.OperationalExpense) error {
	query := `
		INSERT INTO operational_expenses (id, company_id, store_id, expense_date, category,
			description, amount, created_by_user_id, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`
	_, err := r.db.ExecContext(ctx, query,
		oe.ID, oe.CompanyID, oe.StoreID, oe.ExpenseDate, oe.Category,
		oe.Description, oe.Amount, oe.CreatedByUserID, oe.CreatedAt, oe.UpdatedAt,
	)
	return err
}

// GetByID mengambil pengeluaran operasional berdasarkan ID.
func (r *pgOperationalExpenseRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.OperationalExpense, error) {
	oe := &models.OperationalExpense{}
	query := `
		SELECT id, company_id, store_id, expense_date, category,
			description, amount, created_by_user_id, created_at, updated_at
		FROM operational_expenses
		WHERE id = $1`
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&oe.ID, &oe.CompanyID, &oe.StoreID, &oe.ExpenseDate, &oe.Category,
		&oe.Description, &oe.Amount, &oe.CreatedByUserID, &oe.CreatedAt, &oe.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, err
	}
	return oe, nil
}

// ListByCompanyID mengambil daftar pengeluaran operasional untuk perusahaan tertentu.
func (r *pgOperationalExpenseRepository) ListByCompanyID(ctx context.Context, companyID uuid.UUID) ([]*models.OperationalExpense, error) {
	query := `
		SELECT id, company_id, store_id, expense_date, category,
			description, amount, created_by_user_id, created_at, updated_at
		FROM operational_expenses
		WHERE company_id = $1
		ORDER BY expense_date DESC`
	rows, err := r.db.QueryContext(ctx, query, companyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var expenses []*models.OperationalExpense
	for rows.Next() {
		oe := &models.OperationalExpense{}
		if err := rows.Scan(
			&oe.ID, &oe.CompanyID, &oe.StoreID, &oe.ExpenseDate, &oe.Category,
			&oe.Description, &oe.Amount, &oe.CreatedByUserID, &oe.CreatedAt, &oe.UpdatedAt,
		); err != nil {
			return nil, err
		}
		expenses = append(expenses, oe)
	}
	return expenses, rows.Err()
}

// ListByStoreID mengambil daftar pengeluaran operasional untuk toko tertentu.
func (r *pgOperationalExpenseRepository) ListByStoreID(ctx context.Context, storeID uuid.UUID) ([]*models.OperationalExpense, error) {
	query := `
		SELECT id, company_id, store_id, expense_date, category,
			description, amount, created_by_user_id, created_at, updated_at
		FROM operational_expenses
		WHERE store_id = $1
		ORDER BY expense_date DESC`
	rows, err := r.db.QueryContext(ctx, query, storeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var expenses []*models.OperationalExpense
	for rows.Next() {
		oe := &models.OperationalExpense{}
		if err := rows.Scan(
			&oe.ID, &oe.CompanyID, &oe.StoreID, &oe.ExpenseDate, &oe.Category,
			&oe.Description, &oe.Amount, &oe.CreatedByUserID, &oe.CreatedAt, &oe.UpdatedAt,
		); err != nil {
			return nil, err
		}
		expenses = append(expenses, oe)
	}
	return expenses, rows.Err()
}

// Update memperbarui data pengeluaran operasional.
func (r *pgOperationalExpenseRepository) Update(ctx context.Context, oe *models.OperationalExpense) error {
	query := `
		UPDATE operational_expenses
		SET company_id = $1, store_id = $2, expense_date = $3, category = $4,
			description = $5, amount = $6, created_by_user_id = $7, updated_at = $8
		WHERE id = $9`
	_, err := r.db.ExecContext(ctx, query,
		oe.CompanyID, oe.StoreID, oe.ExpenseDate, oe.Category,
		oe.Description, oe.Amount, oe.CreatedByUserID, oe.UpdatedAt, oe.ID,
	)
	return err
}

// Delete menghapus pengeluaran operasional dari database.
func (r *pgOperationalExpenseRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM operational_expenses WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
