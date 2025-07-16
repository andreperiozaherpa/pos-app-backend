package postgres

import (
	"context"
	"database/sql"
	"errors"

	"pos-app/backend/internal/core/repository"
	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

type purchaseOrderRepositoryPG struct {
	db *sql.DB
}

func NewPurchaseOrderRepositoryPG(db *sql.DB) repository.PurchaseOrderRepository {
	return &purchaseOrderRepositoryPG{db: db}
}

func (r *purchaseOrderRepositoryPG) Create(ctx context.Context, po *models.PurchaseOrder) error {
	query := `
        INSERT INTO purchase_orders
        (id, store_id, supplier_id, order_date, expected_delivery_date, status, total_amount, notes, created_by_user_id, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, NOW(), NOW())`
	_, err := r.db.ExecContext(ctx, query,
		po.ID, po.StoreID, po.SupplierID, po.OrderDate, po.ExpectedDeliveryDate, po.Status, po.TotalAmount, po.Notes, po.CreatedByUserID,
	)
	return err
}

func (r *purchaseOrderRepositoryPG) GetByID(ctx context.Context, id uuid.UUID) (*models.PurchaseOrder, error) {
	query := `
        SELECT id, store_id, supplier_id, order_date, expected_delivery_date, status, total_amount, notes, created_by_user_id, created_at, updated_at
        FROM purchase_orders WHERE id = $1`
	row := r.db.QueryRowContext(ctx, query, id)
	po := &models.PurchaseOrder{}
	err := row.Scan(
		&po.ID, &po.StoreID, &po.SupplierID, &po.OrderDate, &po.ExpectedDeliveryDate,
		&po.Status, &po.TotalAmount, &po.Notes, &po.CreatedByUserID, &po.CreatedAt, &po.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}
	return po, nil
}

func (r *purchaseOrderRepositoryPG) ListByStoreID(ctx context.Context, storeID uuid.UUID) ([]*models.PurchaseOrder, error) {
	query := `
		SELECT id, store_id, supplier_id, order_date, expected_delivery_date, status, total_amount, notes, created_by_user_id, created_at, updated_at
		FROM purchase_orders
		WHERE store_id = $1
		ORDER BY order_date DESC`
	rows, err := r.db.QueryContext(ctx, query, storeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []*models.PurchaseOrder
	for rows.Next() {
		po := &models.PurchaseOrder{}
		if err := rows.Scan(
			&po.ID, &po.StoreID, &po.SupplierID, &po.OrderDate, &po.ExpectedDeliveryDate,
			&po.Status, &po.TotalAmount, &po.Notes, &po.CreatedByUserID, &po.CreatedAt, &po.UpdatedAt,
		); err != nil {
			return nil, err
		}
		orders = append(orders, po)
	}
	return orders, nil
}

func (r *purchaseOrderRepositoryPG) Update(ctx context.Context, po *models.PurchaseOrder) error {
	query := `
        UPDATE purchase_orders SET
            store_id = $1,
            supplier_id = $2,
            order_date = $3,
            expected_delivery_date = $4,
            status = $5,
            total_amount = $6,
            notes = $7,
            created_by_user_id = $8,
            updated_at = NOW()
        WHERE id = $9`
	result, err := r.db.ExecContext(ctx, query,
		po.StoreID, po.SupplierID, po.OrderDate, po.ExpectedDeliveryDate, po.Status, po.TotalAmount, po.Notes, po.CreatedByUserID, po.ID,
	)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return repository.ErrNotFound
	}
	return nil
}

func (r *purchaseOrderRepositoryPG) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM purchase_orders WHERE id = $1`
	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return repository.ErrNotFound
	}
	return nil
}

// ListPurchaseOrdersBySupplier mengambil daftar purchase order berdasarkan supplier dengan pagination
func (r *purchaseOrderRepositoryPG) ListPurchaseOrdersBySupplier(ctx context.Context, supplierID uuid.UUID, limit, offset int) ([]*models.PurchaseOrder, error) {
	query := `
        SELECT id, store_id, supplier_id, order_date, expected_delivery_date, status, total_amount, notes, created_by_user_id, created_at, updated_at
        FROM purchase_orders
        WHERE supplier_id = $1
        ORDER BY order_date DESC
        LIMIT $2 OFFSET $3`
	rows, err := r.db.QueryContext(ctx, query, supplierID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []*models.PurchaseOrder
	for rows.Next() {
		po := &models.PurchaseOrder{}
		if err := rows.Scan(
			&po.ID, &po.StoreID, &po.SupplierID, &po.OrderDate, &po.ExpectedDeliveryDate,
			&po.Status, &po.TotalAmount, &po.Notes, &po.CreatedByUserID, &po.CreatedAt, &po.UpdatedAt,
		); err != nil {
			return nil, err
		}
		orders = append(orders, po)
	}
	return orders, nil
}
