package postgres

import (
	"context"
	"database/sql"
	"errors"

	"pos-app/backend/internal/core/repository"
	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

type stockTransferHistoryRepositoryPG struct {
	db *sql.DB
}

func NewStockTransferHistoryRepositoryPG(db *sql.DB) repository.StockTransferHistoryRepository {
	return &stockTransferHistoryRepositoryPG{db: db}
}

func (r *stockTransferHistoryRepositoryPG) Create(ctx context.Context, history *models.StockTransferHistory) error {
	query := `
        INSERT INTO stock_transfer_histories (id, stock_transfer_id, action, action_date, created_at)
        VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.ExecContext(ctx, query,
		history.ID, history.StockTransferID, history.Action, history.ActionDate, history.CreatedAt)
	return err
}

func (r *stockTransferHistoryRepositoryPG) GetByID(ctx context.Context, id uuid.UUID) (*models.StockTransferHistory, error) {
	query := `
        SELECT id, stock_transfer_id, action, action_date, created_at
        FROM stock_transfer_histories WHERE id=$1`
	row := r.db.QueryRowContext(ctx, query, id)
	history := &models.StockTransferHistory{}
	err := row.Scan(&history.ID, &history.StockTransferID, &history.Action, &history.ActionDate, &history.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}
	return history, nil
}

func (r *stockTransferHistoryRepositoryPG) ListByStockTransfer(ctx context.Context, stockTransferID uuid.UUID) ([]*models.StockTransferHistory, error) {
	query := `
        SELECT id, stock_transfer_id, action, action_date, created_at
        FROM stock_transfer_histories
        WHERE stock_transfer_id=$1
        ORDER BY action_date DESC`
	rows, err := r.db.QueryContext(ctx, query, stockTransferID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var histories []*models.StockTransferHistory
	for rows.Next() {
		h := &models.StockTransferHistory{}
		if err := rows.Scan(&h.ID, &h.StockTransferID, &h.Action, &h.ActionDate, &h.CreatedAt); err != nil {
			return nil, err
		}
		histories = append(histories, h)
	}
	return histories, nil
}
