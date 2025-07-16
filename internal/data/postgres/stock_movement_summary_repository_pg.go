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

type stockMovementSummaryRepositoryPG struct {
	db *sql.DB
}

func NewStockMovementSummaryRepositoryPG(db *sql.DB) repository.StockMovementSummaryRepository {
	return &stockMovementSummaryRepositoryPG{db: db}
}

func (r *stockMovementSummaryRepositoryPG) Create(ctx context.Context, summary *models.StockMovementSummary) error {
	query := `
        INSERT INTO stock_movement_summaries (id, store_product_id, period_start, period_end, total_in, total_out, created_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := r.db.ExecContext(ctx, query,
		summary.ID, summary.StoreProductID, summary.PeriodStart, summary.PeriodEnd, summary.TotalIn, summary.TotalOut, summary.CreatedAt)
	return err
}

func (r *stockMovementSummaryRepositoryPG) GetByID(ctx context.Context, id uuid.UUID) (*models.StockMovementSummary, error) {
	query := `
        SELECT id, store_product_id, period_start, period_end, total_in, total_out, created_at
        FROM stock_movement_summaries WHERE id = $1`
	row := r.db.QueryRowContext(ctx, query, id)
	summary := &models.StockMovementSummary{}
	err := row.Scan(&summary.ID, &summary.StoreProductID, &summary.PeriodStart, &summary.PeriodEnd, &summary.TotalIn, &summary.TotalOut, &summary.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}
	return summary, nil
}

func (r *stockMovementSummaryRepositoryPG) ListByStoreProduct(ctx context.Context, storeProductID uuid.UUID, fromDate, toDate time.Time) ([]*models.StockMovementSummary, error) {
	query := `
        SELECT id, store_product_id, period_start, period_end, total_in, total_out, created_at
        FROM stock_movement_summaries
        WHERE store_product_id = $1 AND period_start >= $2 AND period_end <= $3
        ORDER BY period_start DESC`
	rows, err := r.db.QueryContext(ctx, query, storeProductID, fromDate, toDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var summaries []*models.StockMovementSummary
	for rows.Next() {
		s := &models.StockMovementSummary{}
		if err := rows.Scan(&s.ID, &s.StoreProductID, &s.PeriodStart, &s.PeriodEnd, &s.TotalIn, &s.TotalOut, &s.CreatedAt); err != nil {
			return nil, err
		}
		summaries = append(summaries, s)
	}
	return summaries, nil
}
