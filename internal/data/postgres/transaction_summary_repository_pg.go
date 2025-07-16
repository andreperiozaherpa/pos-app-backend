package postgres

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"pos-app/backend/internal/core/repository"
	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

type transactionSummaryRepositoryPG struct {
	db *sql.DB
}

func NewTransactionSummaryRepositoryPG(db *sql.DB) repository.TransactionSummaryRepository {
	return &transactionSummaryRepositoryPG{db: db}
}

func (r *transactionSummaryRepositoryPG) Create(ctx context.Context, summary *models.TransactionSummary) error {
	query := `
        INSERT INTO transaction_summaries (id, store_id, total_transactions, total_revenue, period_start, period_end, created_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := r.db.ExecContext(ctx, query,
		summary.ID, summary.StoreID, summary.TotalTransactions, summary.TotalRevenue,
		summary.PeriodStart, summary.PeriodEnd, time.Now())
	return err
}

func (r *transactionSummaryRepositoryPG) GetByID(ctx context.Context, id uuid.UUID) (*models.TransactionSummary, error) {
	query := `
        SELECT id, store_id, total_transactions, total_revenue, period_start, period_end, created_at
        FROM transaction_summaries WHERE id=$1`
	row := r.db.QueryRowContext(ctx, query, id)
	summary := &models.TransactionSummary{}
	err := row.Scan(
		&summary.ID,
		&summary.StoreID,
		&summary.TotalTransactions,
		&summary.TotalRevenue,
		&summary.PeriodStart,
		&summary.PeriodEnd,
		&summary.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}
	return summary, nil
}

func (r *transactionSummaryRepositoryPG) ListByStoreAndPeriod(ctx context.Context, storeID uuid.UUID, fromDate, toDate time.Time) ([]*models.TransactionSummary, error) {
	query := `
        SELECT id, store_id, total_transactions, total_revenue, period_start, period_end, created_at
        FROM transaction_summaries
        WHERE store_id=$1 AND period_start >= $2 AND period_end <= $3
        ORDER BY period_start DESC`
	rows, err := r.db.QueryContext(ctx, query, storeID, fromDate, toDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var summaries []*models.TransactionSummary
	for rows.Next() {
		s := &models.TransactionSummary{}
		if err := rows.Scan(
			&s.ID, &s.StoreID, &s.TotalTransactions, &s.TotalRevenue,
			&s.PeriodStart, &s.PeriodEnd, &s.CreatedAt,
		); err != nil {
			return nil, err
		}
		summaries = append(summaries, s)
	}
	return summaries, nil
}
