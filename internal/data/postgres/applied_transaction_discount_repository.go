package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

// AppliedTransactionDiscountRepository mendefinisikan interface untuk operasi data terkait AppliedTransactionDiscount.
type AppliedTransactionDiscountRepository interface {
	Create(ctx context.Context, transactionDiscounts []models.AppliedTransactionDiscount) error
	ListByTransactionID(ctx context.Context, transactionID uuid.UUID) ([]*models.AppliedTransactionDiscount, error)
	DeleteByTransactionID(ctx context.Context, transactionID uuid.UUID) error
}

// pgAppliedTransactionDiscountRepository adalah implementasi dari AppliedTransactionDiscountRepository untuk PostgreSQL.
type pgAppliedTransactionDiscountRepository struct {
	db DBExecutor
}

// NewPgAppliedTransactionDiscountRepository adalah constructor untuk membuat instance baru dari pgAppliedTransactionDiscountRepository.
func NewPgAppliedTransactionDiscountRepository(db DBExecutor) AppliedTransactionDiscountRepository {
	return &pgAppliedTransactionDiscountRepository{db: db}
}

// Create menyisipkan satu atau lebih diskon yang diterapkan pada sebuah transaksi.
// Ini harus dipanggil dalam transaksi yang lebih besar jika terkait dengan pembuatan transaksi.
func (r *pgAppliedTransactionDiscountRepository) Create(ctx context.Context, tds []models.AppliedTransactionDiscount) error {
	if len(tds) == 0 {
		return nil // Tidak ada yang perlu disisipkan
	}

	// Memulai transaksi database untuk memastikan semua insert atomik.
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
		return fmt.Errorf("gagal memulai transaksi untuk membuat transaction discounts: %w", err)
	}
	defer tx.Rollback() // Pastikan rollback jika ada error

	stmt, err := tx.PrepareContext(ctx, `
		INSERT INTO applied_transaction_discounts (transaction_id, discount_id, applied_discount_amount_on_transaction)
		VALUES ($1, $2, $3)`)
	if err != nil {
		return fmt.Errorf("gagal menyiapkan statement: %w", err)
	}
	defer stmt.Close()

	for _, td := range tds {
		_, err := stmt.ExecContext(ctx, td.TransactionID, td.DiscountID, td.AppliedDiscountAmountOnTransaction)
		if err != nil {
			return fmt.Errorf("gagal menyisipkan transaction discount (transaction_id: %s, discount_id: %s): %w", td.TransactionID, td.DiscountID, err)
		}
	}

	if !isTx {
		return tx.Commit()
	}
	return nil
}

// ListByTransactionID mengambil semua diskon yang diterapkan pada transaksi tertentu.
func (r *pgAppliedTransactionDiscountRepository) ListByTransactionID(ctx context.Context, transactionID uuid.UUID) ([]*models.AppliedTransactionDiscount, error) {
	query := `
		SELECT transaction_id, discount_id, applied_discount_amount_on_transaction
		FROM applied_transaction_discounts
		WHERE transaction_id = $1`
	rows, err := r.db.QueryContext(ctx, query, transactionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var discounts []*models.AppliedTransactionDiscount
	for rows.Next() {
		td := &models.AppliedTransactionDiscount{}
		if err := rows.Scan(&td.TransactionID, &td.DiscountID, &td.AppliedDiscountAmountOnTransaction); err != nil {
			return nil, err
		}
		discounts = append(discounts, td)
	}
	return discounts, rows.Err()
}

// DeleteByTransactionID menghapus semua diskon yang terkait dengan transaksi tertentu.
func (r *pgAppliedTransactionDiscountRepository) DeleteByTransactionID(ctx context.Context, transactionID uuid.UUID) error {
	query := `DELETE FROM applied_transaction_discounts WHERE transaction_id = $1`
	_, err := r.db.ExecContext(ctx, query, transactionID)
	return err
}
