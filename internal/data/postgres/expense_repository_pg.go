package postgres

import (
	"context"
	"database/sql"
	"errors"

	"pos-app/backend/internal/core/repository"
	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

type expenseRepositoryPG struct {
	db *sql.DB
}

func NewExpenseRepositoryPG(db *sql.DB) repository.ExpenseRepository {
	return &expenseRepositoryPG{db: db}
}

func (r *expenseRepositoryPG) Create(ctx context.Context, expense *models.OperationalExpense) error {
	query := `
        INSERT INTO operational_expenses (id, company_id, store_id, description, amount, expense_date, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6, NOW(), NOW())`
	_, err := r.db.ExecContext(ctx, query,
		expense.ID, expense.CompanyID, expense.StoreID,
		expense.Description, expense.Amount, expense.ExpenseDate)
	return err
}

func (r *expenseRepositoryPG) GetByID(ctx context.Context, id uuid.UUID) (*models.OperationalExpense, error) {
	query := `
        SELECT id, company_id, store_id, description, amount, expense_date, created_at, updated_at
        FROM operational_expenses WHERE id=$1`
	row := r.db.QueryRowContext(ctx, query, id)
	expense := &models.OperationalExpense{}
	err := row.Scan(
		&expense.ID, &expense.CompanyID, &expense.StoreID, &expense.Description,
		&expense.Amount, &expense.ExpenseDate, &expense.CreatedAt, &expense.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}
	return expense, nil
}

func (r *expenseRepositoryPG) Update(ctx context.Context, expense *models.OperationalExpense) error {
	query := `
        UPDATE operational_expenses SET company_id=$1, store_id=$2, description=$3, amount=$4, expense_date=$5, updated_at=NOW()
        WHERE id=$6`
	res, err := r.db.ExecContext(ctx, query,
		expense.CompanyID, expense.StoreID, expense.Description,
		expense.Amount, expense.ExpenseDate, expense.ID)
	if err != nil {
		return err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return repository.ErrNotFound
	}
	return nil
}

func (r *expenseRepositoryPG) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM operational_expenses WHERE id=$1`
	res, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return repository.ErrNotFound
	}
	return nil
}

func (r *expenseRepositoryPG) ListByCompany(ctx context.Context, companyID uuid.UUID, limit, offset int) ([]*models.OperationalExpense, error) {
	query := `
        SELECT id, company_id, store_id, description, amount, expense_date, created_at, updated_at
        FROM operational_expenses
        WHERE company_id=$1
        ORDER BY expense_date DESC
        LIMIT $2 OFFSET $3`
	rows, err := r.db.QueryContext(ctx, query, companyID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var expenses []*models.OperationalExpense
	for rows.Next() {
		expense := &models.OperationalExpense{}
		if err := rows.Scan(
			&expense.ID, &expense.CompanyID, &expense.StoreID, &expense.Description,
			&expense.Amount, &expense.ExpenseDate, &expense.CreatedAt, &expense.UpdatedAt); err != nil {
			return nil, err
		}
		expenses = append(expenses, expense)
	}
	return expenses, nil
}
