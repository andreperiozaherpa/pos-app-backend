package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

// PurchaseOrderRepository mendefinisikan interface untuk operasi data terkait PurchaseOrder.
type PurchaseOrderRepository interface {
	Create(ctx context.Context, po *models.PurchaseOrder) error
	GetByID(ctx context.Context, id uuid.UUID) (*models.PurchaseOrder, error)
	ListByStoreID(ctx context.Context, storeID uuid.UUID) ([]*models.PurchaseOrder, error)
	Update(ctx context.Context, po *models.PurchaseOrder) error
	Delete(ctx context.Context, id uuid.UUID) error
}

// pgPurchaseOrderRepository adalah implementasi dari PurchaseOrderRepository untuk PostgreSQL.
type pgPurchaseOrderRepository struct {
	db DBExecutor
}

// NewPgPurchaseOrderRepository adalah constructor untuk membuat instance baru dari pgPurchaseOrderRepository.
func NewPgPurchaseOrderRepository(db DBExecutor) PurchaseOrderRepository {
	return &pgPurchaseOrderRepository{db: db}
}

// Implementasi metode-metode dari interface PurchaseOrderRepository:

func (r *pgPurchaseOrderRepository) Create(ctx context.Context, po *models.PurchaseOrder) error {
	tx, isTx := r.db.(*sql.Tx)
	var err error
	if !isTx {
		db, ok := r.db.(*sql.DB)
		if !ok {
			return fmt.Errorf("unexpected DBExecutor type; expected *sql.DB or *sql.Tx")
		}
		tx, err = db.BeginTx(ctx, nil)
	}

	if err != nil {
		return fmt.Errorf("gagal memulai transaksi: %w", err)
	}
	defer tx.Rollback() // Rollback jika ada error atau panic

	// 1. Sisipkan data ke tabel 'purchase_orders' (header)
	poQuery := `
		INSERT INTO purchase_orders (id, store_id, supplier_id, order_date, expected_delivery_date, 
			status, total_amount, notes, created_by_user_id, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`
	_, err = tx.ExecContext(ctx, poQuery,
		po.ID, po.StoreID, po.SupplierID, po.OrderDate, po.ExpectedDeliveryDate,
		po.Status, po.TotalAmount, po.Notes, po.CreatedByUserID, po.CreatedAt, po.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("gagal menyisipkan header purchase order: %w", err)
	}

	// 2. Sisipkan setiap item ke tabel 'purchase_order_items'
	itemQuery := `
		INSERT INTO purchase_order_items (id, purchase_order_id, master_product_id, quantity_ordered, 
			purchase_price_per_unit, quantity_received, subtotal, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	for _, item := range po.Items {
		_, err = tx.ExecContext(ctx, itemQuery,
			item.ID, po.ID, item.MasterProductID, item.QuantityOrdered,
			item.PurchasePricePerUnit, item.QuantityReceived, item.Subtotal, item.CreatedAt, item.UpdatedAt,
		)
		if err != nil {
			return fmt.Errorf("gagal menyisipkan item purchase order (produk ID %s): %w", item.MasterProductID, err)
		}
	}

	// 3. Commit transaksi
	if !isTx {
		return tx.Commit()
	}
	return nil
}

func (r *pgPurchaseOrderRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.PurchaseOrder, error) {
	po := &models.PurchaseOrder{}
	// Query untuk header purchase order
	poQuery := `
		SELECT id, store_id, supplier_id, order_date, expected_delivery_date, 
			status, total_amount, notes, created_by_user_id, created_at, updated_at
		FROM purchase_orders
		WHERE id = $1`
	err := r.db.QueryRowContext(ctx, poQuery, id).Scan(
		&po.ID, &po.StoreID, &po.SupplierID, &po.OrderDate, &po.ExpectedDeliveryDate,
		&po.Status, &po.TotalAmount, &po.Notes, &po.CreatedByUserID, &po.CreatedAt, &po.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, err
	}

	// Query untuk item-item purchase order
	itemQuery := `
		SELECT id, purchase_order_id, master_product_id, quantity_ordered, 
			purchase_price_per_unit, quantity_received, subtotal, created_at, updated_at
		FROM purchase_order_items
		WHERE purchase_order_id = $1`
	rows, err := r.db.QueryContext(ctx, itemQuery, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []models.PurchaseOrderItem
	for rows.Next() {
		item := models.PurchaseOrderItem{}
		if err := rows.Scan(
			&item.ID, &item.PurchaseOrderID, &item.MasterProductID, &item.QuantityOrdered,
			&item.PurchasePricePerUnit, &item.QuantityReceived, &item.Subtotal, &item.CreatedAt, &item.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	po.Items = items

	return po, rows.Err()
}

func (r *pgPurchaseOrderRepository) ListByStoreID(ctx context.Context, storeID uuid.UUID) ([]*models.PurchaseOrder, error) {
	query := `
		SELECT id, store_id, supplier_id, order_date, status, total_amount, created_by_user_id, created_at
		FROM purchase_orders
		WHERE store_id = $1
		ORDER BY order_date DESC`
	rows, err := r.db.QueryContext(ctx, query, storeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var purchaseOrders []*models.PurchaseOrder
	for rows.Next() {
		po := &models.PurchaseOrder{}
		if err := rows.Scan(&po.ID, &po.StoreID, &po.SupplierID, &po.OrderDate, &po.Status, &po.TotalAmount, &po.CreatedByUserID, &po.CreatedAt); err != nil {
			return nil, err
		}
		purchaseOrders = append(purchaseOrders, po)
	}
	return purchaseOrders, rows.Err()
}

func (r *pgPurchaseOrderRepository) Update(ctx context.Context, po *models.PurchaseOrder) error {
	query := `
		UPDATE purchase_orders
		SET store_id = $1, supplier_id = $2, order_date = $3, expected_delivery_date = $4, 
			status = $5, total_amount = $6, notes = $7, created_by_user_id = $8, updated_at = $9
		WHERE id = $10`
	_, err := r.db.ExecContext(ctx, query,
		po.StoreID, po.SupplierID, po.OrderDate, po.ExpectedDeliveryDate,
		po.Status, po.TotalAmount, po.Notes, po.CreatedByUserID, po.UpdatedAt, po.ID,
	)
	return err
}

func (r *pgPurchaseOrderRepository) Delete(ctx context.Context, id uuid.UUID) error {
	tx, isTx := r.db.(*sql.Tx)
	var err error
	if !isTx {
		db, ok := r.db.(*sql.DB)
		if !ok {
			return fmt.Errorf("unexpected DBExecutor type; expected *sql.DB or *sql.Tx")
		}
		tx, err = db.BeginTx(ctx, nil)
	}

	if err != nil {
		return fmt.Errorf("gagal memulai transaksi delete: %w", err)
	}
	defer tx.Rollback() // Rollback jika ada error

	// 1. Hapus item-item terkait dari purchase_order_items
	deleteItemsQuery := `DELETE FROM purchase_order_items WHERE purchase_order_id = $1`
	_, err = tx.ExecContext(ctx, deleteItemsQuery, id)
	if err != nil {
		return fmt.Errorf("gagal menghapus item purchase order: %w", err)
	}

	// 2. Hapus purchase_order itu sendiri
	deletePOQuery := `DELETE FROM purchase_orders WHERE id = $1`
	_, err = tx.ExecContext(ctx, deletePOQuery, id)
	if err != nil {
		return fmt.Errorf("gagal menghapus purchase order: %w", err)
	}

	if !isTx {
		return tx.Commit()
	}
	return nil

}
