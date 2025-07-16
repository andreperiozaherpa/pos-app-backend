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

type transactionRepositoryPG struct {
	db *sql.DB
}

func NewTransactionRepositoryPG(db *sql.DB) repository.TransactionRepository {
	return &transactionRepositoryPG{db: db}
}

// Create menyimpan transaksi baru ke database
func (r *transactionRepositoryPG) Create(ctx context.Context, t *models.Transaction) error {
	query := `
        INSERT INTO transactions
        (id, transaction_code, store_id, cashier_employee_user_id, customer_user_id, active_shift_id,
         transaction_date, subtotal_amount, total_item_discount_amount, subtotal_after_item_discounts,
         transaction_level_discount_amount, taxable_amount, total_tax_amount, final_total_amount,
         received_amount, change_amount, payment_method, notes, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, NOW(), NOW())`
	_, err := r.db.ExecContext(ctx, query,
		t.ID, t.TransactionCode, t.StoreID, t.CashierEmployeeUserID, t.CustomerUserID, t.ActiveShiftID,
		t.TransactionDate, t.SubtotalAmount, t.TotalItemDiscountAmount, t.SubtotalAfterItemDiscounts,
		t.TransactionLevelDiscountAmount, t.TaxableAmount, t.TotalTaxAmount, t.FinalTotalAmount,
		t.ReceivedAmount, t.ChangeAmount, t.PaymentMethod, t.Notes)
	return err
}

// GetByID mengambil transaksi berdasarkan ID
func (r *transactionRepositoryPG) GetByID(ctx context.Context, id uuid.UUID) (*models.Transaction, error) {
	query := `
        SELECT id, transaction_code, store_id, cashier_employee_user_id, customer_user_id, active_shift_id,
               transaction_date, subtotal_amount, total_item_discount_amount, subtotal_after_item_discounts,
               transaction_level_discount_amount, taxable_amount, total_tax_amount, final_total_amount,
               received_amount, change_amount, payment_method, notes, created_at, updated_at
        FROM transactions WHERE id = $1`
	row := r.db.QueryRowContext(ctx, query, id)
	t := &models.Transaction{}
	err := row.Scan(
		&t.ID, &t.TransactionCode, &t.StoreID, &t.CashierEmployeeUserID, &t.CustomerUserID, &t.ActiveShiftID,
		&t.TransactionDate, &t.SubtotalAmount, &t.TotalItemDiscountAmount, &t.SubtotalAfterItemDiscounts,
		&t.TransactionLevelDiscountAmount, &t.TaxableAmount, &t.TotalTaxAmount, &t.FinalTotalAmount,
		&t.ReceivedAmount, &t.ChangeAmount, &t.PaymentMethod, &t.Notes, &t.CreatedAt, &t.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}
	return t, nil
}

// GetByTransactionCode mengambil transaksi berdasarkan kode transaksi
func (r *transactionRepositoryPG) GetByTransactionCode(ctx context.Context, code string) (*models.Transaction, error) {
	query := `
        SELECT id, transaction_code, store_id, cashier_employee_user_id, customer_user_id, active_shift_id,
               transaction_date, subtotal_amount, total_item_discount_amount, subtotal_after_item_discounts,
               transaction_level_discount_amount, taxable_amount, total_tax_amount, final_total_amount,
               received_amount, change_amount, payment_method, notes, created_at, updated_at
        FROM transactions WHERE transaction_code = $1`
	row := r.db.QueryRowContext(ctx, query, code)
	t := &models.Transaction{}
	err := row.Scan(
		&t.ID, &t.TransactionCode, &t.StoreID, &t.CashierEmployeeUserID, &t.CustomerUserID, &t.ActiveShiftID,
		&t.TransactionDate, &t.SubtotalAmount, &t.TotalItemDiscountAmount, &t.SubtotalAfterItemDiscounts,
		&t.TransactionLevelDiscountAmount, &t.TaxableAmount, &t.TotalTaxAmount, &t.FinalTotalAmount,
		&t.ReceivedAmount, &t.ChangeAmount, &t.PaymentMethod, &t.Notes, &t.CreatedAt, &t.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}
	return t, nil
}

// ListByStoreID mengambil daftar transaksi berdasarkan Store ID.
func (r *transactionRepositoryPG) ListByStoreID(ctx context.Context, storeID uuid.UUID) ([]*models.Transaction, error) {
	query := `
        SELECT id, transaction_code, store_id, cashier_employee_user_id, customer_user_id, active_shift_id,
               transaction_date, subtotal_amount, total_item_discount_amount, subtotal_after_item_discounts,
               transaction_level_discount_amount, taxable_amount, total_tax_amount, final_total_amount,
               received_amount, change_amount, payment_method, notes, created_at, updated_at
        FROM transactions
        WHERE store_id = $1
        ORDER BY transaction_date DESC`
	rows, err := r.db.QueryContext(ctx, query, storeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []*models.Transaction
	for rows.Next() {
		var t models.Transaction
		err := rows.Scan(
			&t.ID, &t.TransactionCode, &t.StoreID, &t.CashierEmployeeUserID, &t.CustomerUserID, &t.ActiveShiftID,
			&t.TransactionDate, &t.SubtotalAmount, &t.TotalItemDiscountAmount, &t.SubtotalAfterItemDiscounts,
			&t.TransactionLevelDiscountAmount, &t.TaxableAmount, &t.TotalTaxAmount, &t.FinalTotalAmount,
			&t.ReceivedAmount, &t.ChangeAmount, &t.PaymentMethod, &t.Notes, &t.CreatedAt, &t.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, &t)
	}
	return transactions, nil
}

// Update mengupdate data transaksi yang sudah ada
func (r *transactionRepositoryPG) Update(ctx context.Context, t *models.Transaction) error {
	query := `
        UPDATE transactions SET
            transaction_code = $1,
            store_id = $2,
            cashier_employee_user_id = $3,
            customer_user_id = $4,
            active_shift_id = $5,
            transaction_date = $6,
            subtotal_amount = $7,
            total_item_discount_amount = $8,
            subtotal_after_item_discounts = $9,
            transaction_level_discount_amount = $10,
            taxable_amount = $11,
            total_tax_amount = $12,
            final_total_amount = $13,
            received_amount = $14,
            change_amount = $15,
            payment_method = $16,
            notes = $17,
            updated_at = NOW()
        WHERE id = $18`
	result, err := r.db.ExecContext(ctx, query,
		t.TransactionCode, t.StoreID, t.CashierEmployeeUserID, t.CustomerUserID, t.ActiveShiftID,
		t.TransactionDate, t.SubtotalAmount, t.TotalItemDiscountAmount, t.SubtotalAfterItemDiscounts,
		t.TransactionLevelDiscountAmount, t.TaxableAmount, t.TotalTaxAmount, t.FinalTotalAmount,
		t.ReceivedAmount, t.ChangeAmount, t.PaymentMethod, t.Notes, t.ID)
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

// Delete menghapus transaksi berdasarkan ID
func (r *transactionRepositoryPG) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM transactions WHERE id = $1`
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

// ListTransactionsByDateRange mengambil transaksi dalam rentang tanggal
func (r *transactionRepositoryPG) ListTransactionsByDateRange(ctx context.Context, startDate, endDate time.Time) ([]*models.Transaction, error) {
	query := `
        SELECT id, transaction_code, store_id, cashier_employee_user_id, customer_user_id, active_shift_id,
               transaction_date, subtotal_amount, total_item_discount_amount, subtotal_after_item_discounts,
               transaction_level_discount_amount, taxable_amount, total_tax_amount, final_total_amount,
               received_amount, change_amount, payment_method, notes, created_at, updated_at
        FROM transactions
        WHERE created_at BETWEEN $1 AND $2
        ORDER BY created_at DESC`
	rows, err := r.db.QueryContext(ctx, query, startDate, endDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []*models.Transaction
	for rows.Next() {
		t := &models.Transaction{}
		if err := rows.Scan(
			&t.ID, &t.TransactionCode, &t.StoreID, &t.CashierEmployeeUserID, &t.CustomerUserID, &t.ActiveShiftID,
			&t.TransactionDate, &t.SubtotalAmount, &t.TotalItemDiscountAmount, &t.SubtotalAfterItemDiscounts,
			&t.TransactionLevelDiscountAmount, &t.TaxableAmount, &t.TotalTaxAmount, &t.FinalTotalAmount,
			&t.ReceivedAmount, &t.ChangeAmount, &t.PaymentMethod, &t.Notes, &t.CreatedAt, &t.UpdatedAt); err != nil {
			return nil, err
		}
		transactions = append(transactions, t)
	}
	return transactions, nil
}

// GetAuditTrail mengambil histori audit log transaksi dari transaction_audit_logs
func (r *transactionRepositoryPG) GetAuditTrail(ctx context.Context, transactionID uuid.UUID) ([]*models.TransactionAuditLog, error) {
	query := `
        SELECT id, transaction_id, action_type, performed_by_user_id, performed_at, note
        FROM transaction_audit_logs
        WHERE transaction_id = $1
        ORDER BY performed_at DESC`
	rows, err := r.db.QueryContext(ctx, query, transactionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var logs []*models.TransactionAuditLog
	for rows.Next() {
		log := &models.TransactionAuditLog{}
		err := rows.Scan(
			&log.ID,
			&log.TransactionID,
			&log.ActionType,
			&log.PerformedByUserID,
			&log.PerformedAt,
			&log.Note,
		)
		if err != nil {
			return nil, err
		}
		logs = append(logs, log)
	}
	return logs, nil
}
