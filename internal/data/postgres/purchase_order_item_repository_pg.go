package postgres

import (
	"context"
	"database/sql"
	"errors"

	"pos-app/backend/internal/core/repository"
	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

type purchaseOrderItemRepositoryPG struct {
	db *sql.DB
}

func NewPurchaseOrderItemRepositoryPG(db *sql.DB) repository.PurchaseOrderItemRepository {
	return &purchaseOrderItemRepositoryPG{db: db}
}

func (r *purchaseOrderItemRepositoryPG) Create(ctx context.Context, item *models.PurchaseOrderItem) error {
	query := `
        INSERT INTO purchase_order_items (
			id, purchase_order_id, master_product_id, quantity_ordered, 
			purchase_price_per_unit, quantity_received, subtotal,
			created_at, updated_at
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, NOW(), NOW()
		)`
	_, err := r.db.ExecContext(ctx, query,
		item.ID, item.PurchaseOrderID, item.MasterProductID,
		item.QuantityOrdered, item.PurchasePricePerUnit, item.QuantityReceived,
		item.Subtotal,
	)
	return err
}

func (r *purchaseOrderItemRepositoryPG) GetByID(ctx context.Context, id uuid.UUID) (*models.PurchaseOrderItem, error) {
	query := `
        SELECT id, purchase_order_id, master_product_id, quantity_ordered, 
			purchase_price_per_unit, quantity_received, subtotal,
			created_at, updated_at
        FROM purchase_order_items WHERE id = $1`
	row := r.db.QueryRowContext(ctx, query, id)
	item := &models.PurchaseOrderItem{}
	err := row.Scan(
		&item.ID, &item.PurchaseOrderID, &item.MasterProductID,
		&item.QuantityOrdered, &item.PurchasePricePerUnit, &item.QuantityReceived,
		&item.Subtotal, &item.CreatedAt, &item.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}
	return item, nil
}

// ListByPurchaseOrderID mengambil daftar item berdasarkan purchase order ID.
func (r *purchaseOrderItemRepositoryPG) ListByPurchaseOrderID(ctx context.Context, purchaseOrderID uuid.UUID) ([]*models.PurchaseOrderItem, error) {
	query := `
		SELECT id, purchase_order_id, master_product_id, quantity_ordered, 
			purchase_price_per_unit, quantity_received, subtotal,
			created_at, updated_at
		FROM purchase_order_items WHERE purchase_order_id = $1`
	rows, err := r.db.QueryContext(ctx, query, purchaseOrderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []*models.PurchaseOrderItem
	for rows.Next() {
		item := &models.PurchaseOrderItem{}
		if err := rows.Scan(
			&item.ID, &item.PurchaseOrderID, &item.MasterProductID,
			&item.QuantityOrdered, &item.PurchasePricePerUnit, &item.QuantityReceived,
			&item.Subtotal, &item.CreatedAt, &item.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func (r *purchaseOrderItemRepositoryPG) Update(ctx context.Context, item *models.PurchaseOrderItem) error {
	query := `
        UPDATE purchase_order_items SET
			purchase_order_id = $1,
			master_product_id = $2,
			quantity_ordered = $3,
			purchase_price_per_unit = $4,
			quantity_received = $5,
			subtotal = $6,
			updated_at = NOW()
        WHERE id = $7`
	res, err := r.db.ExecContext(ctx, query,
		item.PurchaseOrderID, item.MasterProductID,
		item.QuantityOrdered, item.PurchasePricePerUnit,
		item.QuantityReceived, item.Subtotal, item.ID,
	)
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

func (r *purchaseOrderItemRepositoryPG) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM purchase_order_items WHERE id = $1`
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

func (r *purchaseOrderItemRepositoryPG) ListByPurchaseOrder(ctx context.Context, purchaseOrderID uuid.UUID) ([]*models.PurchaseOrderItem, error) {
	query := `
		SELECT id, purchase_order_id, master_product_id, quantity_ordered, 
			purchase_price_per_unit, quantity_received, subtotal,
			created_at, updated_at
		FROM purchase_order_items
		WHERE purchase_order_id = $1`
	rows, err := r.db.QueryContext(ctx, query, purchaseOrderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []*models.PurchaseOrderItem
	for rows.Next() {
		item := &models.PurchaseOrderItem{}
		if err := rows.Scan(
			&item.ID, &item.PurchaseOrderID, &item.MasterProductID,
			&item.QuantityOrdered, &item.PurchasePricePerUnit, &item.QuantityReceived,
			&item.Subtotal, &item.CreatedAt, &item.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}
