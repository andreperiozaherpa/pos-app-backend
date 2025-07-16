package postgres

import (
	"context"
	"database/sql"
	"errors"

	"pos-app/backend/internal/core/repository"
	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

// Sesuaikan nama repository jika perlu
type storeProductStockUpdateRepositoryPG struct {
	db *sql.DB
}

func NewStoreProductStockUpdateRepositoryPG(db *sql.DB) repository.StoreProductStockUpdateRepository {
	return &storeProductStockUpdateRepositoryPG{db: db}
}

// Create menambahkan data perubahan stok produk toko
func (r *storeProductStockUpdateRepositoryPG) Create(ctx context.Context, update *models.StoreProductStockUpdate) error {
	query := `
        INSERT INTO store_product_stock_updates 
        (id, store_product_id, adjustment_type, quantity_before, quantity_after, adjusted_by_user_id, reason, adjustment_date, created_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, NOW())`
	_, err := r.db.ExecContext(ctx, query,
		update.ID,
		update.StoreProductID,
		update.AdjustmentType,
		update.QuantityBefore,
		update.QuantityAfter,
		update.AdjustedByUserID,
		update.Reason,
		update.AdjustmentDate,
	)
	return err
}

// GetByID mengambil perubahan stok berdasarkan ID update
func (r *storeProductStockUpdateRepositoryPG) GetByID(ctx context.Context, id uuid.UUID) (*models.StoreProductStockUpdate, error) {
	query := `
        SELECT id, store_product_id, adjustment_type, quantity_before, quantity_after, adjusted_by_user_id, reason, adjustment_date, created_at
        FROM store_product_stock_updates WHERE id = $1`
	row := r.db.QueryRowContext(ctx, query, id)
	update := &models.StoreProductStockUpdate{}
	err := row.Scan(
		&update.ID,
		&update.StoreProductID,
		&update.AdjustmentType,
		&update.QuantityBefore,
		&update.QuantityAfter,
		&update.AdjustedByUserID,
		&update.Reason,
		&update.AdjustmentDate,
		&update.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}
	return update, nil
}

// ListByStoreProductID mengambil list perubahan stok berdasarkan store_product_id
func (r *storeProductStockUpdateRepositoryPG) ListByStoreProductID(ctx context.Context, storeProductID uuid.UUID) ([]*models.StoreProductStockUpdate, error) {
	query := `
        SELECT id, store_product_id, adjustment_type, quantity_before, quantity_after, adjusted_by_user_id, reason, adjustment_date, created_at
        FROM store_product_stock_updates 
        WHERE store_product_id = $1 
        ORDER BY adjustment_date DESC`
	rows, err := r.db.QueryContext(ctx, query, storeProductID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var updates []*models.StoreProductStockUpdate
	for rows.Next() {
		u := &models.StoreProductStockUpdate{}
		if err := rows.Scan(
			&u.ID,
			&u.StoreProductID,
			&u.AdjustmentType,
			&u.QuantityBefore,
			&u.QuantityAfter,
			&u.AdjustedByUserID,
			&u.Reason,
			&u.AdjustmentDate,
			&u.CreatedAt,
		); err != nil {
			return nil, err
		}
		updates = append(updates, u)
	}
	return updates, nil
}
