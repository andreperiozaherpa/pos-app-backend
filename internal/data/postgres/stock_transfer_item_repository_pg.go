package postgres

import (
	"context"
	"database/sql"
	"errors"

	"pos-app/backend/internal/core/repository"
	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

type stockTransferItemRepositoryPG struct {
	db *sql.DB
}

func NewStockTransferItemRepositoryPG(db *sql.DB) repository.StockTransferItemRepository {
	return &stockTransferItemRepositoryPG{db: db}
}

func (r *stockTransferItemRepositoryPG) Create(ctx context.Context, item *models.StockTransferItem) error {
	query := `
        INSERT INTO stock_transfer_items (id, stock_transfer_id, product_id, quantity, created_at, updated_at)
        VALUES ($1, $2, $3, $4, NOW(), NOW())`
	_, err := r.db.ExecContext(ctx, query,
		item.ID, item.StockTransferID, item.ProductID, item.Quantity)
	return err
}

func (r *stockTransferItemRepositoryPG) GetByID(ctx context.Context, id uuid.UUID) (*models.StockTransferItem, error) {
	query := `
        SELECT id, stock_transfer_id, product_id, quantity, created_at, updated_at
        FROM stock_transfer_items WHERE id=$1`
	row := r.db.QueryRowContext(ctx, query, id)
	item := &models.StockTransferItem{}
	err := row.Scan(
		&item.ID, &item.StockTransferID, &item.ProductID,
		&item.Quantity, &item.CreatedAt, &item.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}
	return item, nil
}

func (r *stockTransferItemRepositoryPG) Update(ctx context.Context, item *models.StockTransferItem) error {
	query := `
        UPDATE stock_transfer_items SET stock_transfer_id=$1, product_id=$2, quantity=$3, updated_at=NOW()
        WHERE id=$4`
	res, err := r.db.ExecContext(ctx, query,
		item.StockTransferID, item.ProductID, item.Quantity, item.ID)
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

func (r *stockTransferItemRepositoryPG) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM stock_transfer_items WHERE id=$1`
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

func (r *stockTransferItemRepositoryPG) ListByStockTransfer(ctx context.Context, stockTransferID uuid.UUID) ([]*models.StockTransferItem, error) {
	query := `
        SELECT id, stock_transfer_id, product_id, quantity, created_at, updated_at
        FROM stock_transfer_items WHERE stock_transfer_id=$1`
	rows, err := r.db.QueryContext(ctx, query, stockTransferID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []*models.StockTransferItem
	for rows.Next() {
		item := &models.StockTransferItem{}
		if err := rows.Scan(
			&item.ID, &item.StockTransferID, &item.ProductID,
			&item.Quantity, &item.CreatedAt, &item.UpdatedAt); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}
