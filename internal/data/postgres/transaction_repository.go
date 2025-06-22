package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

// TransactionRepository mendefinisikan interface untuk operasi data terkait Transaction.
type TransactionRepository interface {
	Create(ctx context.Context, transaction *models.Transaction) error
	GetByID(ctx context.Context, id uuid.UUID) (*models.Transaction, error)
	ListByStoreID(ctx context.Context, storeID uuid.UUID) ([]*models.Transaction, error)
}

// pgTransactionRepository adalah implementasi dari TransactionRepository untuk PostgreSQL.
type pgTransactionRepository struct {
	db *sql.DB
}

// NewPgTransactionRepository adalah constructor untuk membuat instance baru dari pgTransactionRepository.
func NewPgTransactionRepository(db *sql.DB) TransactionRepository {
	return &pgTransactionRepository{db: db}
}

// Create menyisipkan transaksi baru dan semua itemnya secara atomik.
func (r *pgTransactionRepository) Create(ctx context.Context, t *models.Transaction) error {
	// Memulai transaksi database
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("gagal memulai transaksi: %w", err)
	}
	// Defer a rollback in case of panic or error
	defer tx.Rollback()

	// 1. Sisipkan data ke tabel 'transactions' (header)
	txQuery := `
		INSERT INTO transactions (id, transaction_code, store_id, cashier_employee_user_id, customer_user_id, 
			active_shift_id, transaction_date, subtotal_amount, total_item_discount_amount, subtotal_after_item_discounts, 
			transaction_level_discount_amount, taxable_amount, total_tax_amount, final_total_amount, 
			received_amount, change_amount, payment_method, notes, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20)`
	_, err = tx.ExecContext(ctx, txQuery,
		t.ID, t.TransactionCode, t.StoreID, t.CashierEmployeeUserID, t.CustomerUserID, t.ActiveShiftID,
		t.TransactionDate, t.SubtotalAmount, t.TotalItemDiscountAmount, t.SubtotalAfterItemDiscounts,
		t.TransactionLevelDiscountAmount, t.TaxableAmount, t.TotalTaxAmount, t.FinalTotalAmount,
		t.ReceivedAmount, t.ChangeAmount, t.PaymentMethod, t.Notes, t.CreatedAt, t.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("gagal menyisipkan header transaksi: %w", err)
	}

	// 2. Sisipkan setiap item ke tabel 'transaction_items'
	itemQuery := `
		INSERT INTO transaction_items (id, transaction_id, store_product_id, quantity, price_per_unit_at_transaction, 
			item_subtotal_before_discount, item_discount_amount, item_subtotal_after_discount, applied_tax_rate_id, 
			applied_tax_rate_percentage, tax_amount_for_item, item_final_total, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)`
	for _, item := range t.Items {
		_, err = tx.ExecContext(ctx, itemQuery,
			item.ID, t.ID, item.StoreProductID, item.Quantity, item.PricePerUnitAtTransaction,
			item.ItemSubtotalBeforeDiscount, item.ItemDiscountAmount, item.ItemSubtotalAfterDiscount,
			item.AppliedTaxRateID, item.AppliedTaxRatePercentage, item.TaxAmountForItem, item.ItemFinalTotal,
			item.CreatedAt, item.UpdatedAt,
		)
		if err != nil {
			return fmt.Errorf("gagal menyisipkan item transaksi (produk ID %s): %w", item.StoreProductID, err)
		}
	}

	// 3. Jika semua berhasil, commit transaksi
	return tx.Commit()
}

// GetByID mengambil data transaksi beserta semua itemnya.
func (r *pgTransactionRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.Transaction, error) {
	t := &models.Transaction{}
	// Query untuk header transaksi
	txQuery := `SELECT id, transaction_code, store_id, cashier_employee_user_id, customer_user_id, final_total_amount, created_at FROM transactions WHERE id = $1`
	err := r.db.QueryRowContext(ctx, txQuery, id).Scan(&t.ID, &t.TransactionCode, &t.StoreID, &t.CashierEmployeeUserID, &t.CustomerUserID, &t.FinalTotalAmount, &t.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, err
	}

	// Query untuk item-item transaksi
	itemQuery := `SELECT id, store_product_id, quantity, price_per_unit_at_transaction, item_final_total FROM transaction_items WHERE transaction_id = $1`
	rows, err := r.db.QueryContext(ctx, itemQuery, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		item := models.TransactionItem{}
		if err := rows.Scan(&item.ID, &item.StoreProductID, &item.Quantity, &item.PricePerUnitAtTransaction, &item.ItemFinalTotal); err != nil {
			return nil, err
		}
		t.Items = append(t.Items, item)
	}

	return t, rows.Err()
}

// ListByStoreID hanya mengambil header transaksi untuk efisiensi.
func (r *pgTransactionRepository) ListByStoreID(ctx context.Context, storeID uuid.UUID) ([]*models.Transaction, error) {
	// Implementasi untuk daftar transaksi
	// Untuk saat ini, kita bisa mengabaikannya atau membuat implementasi sederhana
	return nil, fmt.Errorf("ListByStoreID belum diimplementasikan")
}
