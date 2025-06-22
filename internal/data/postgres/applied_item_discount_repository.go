package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

// AppliedItemDiscountRepository mendefinisikan interface untuk operasi data terkait AppliedItemDiscount.
type AppliedItemDiscountRepository interface {
	Create(ctx context.Context, itemDiscounts []models.AppliedItemDiscount) error
	ListByTransactionItemID(ctx context.Context, transactionItemID uuid.UUID) ([]*models.AppliedItemDiscount, error)
	DeleteByTransactionItemID(ctx context.Context, transactionItemID uuid.UUID) error
}

// pgAppliedItemDiscountRepository adalah implementasi dari AppliedItemDiscountRepository untuk PostgreSQL.
type pgAppliedItemDiscountRepository struct {
	db DBExecutor
}

// NewPgAppliedItemDiscountRepository adalah constructor untuk membuat instance baru dari pgAppliedItemDiscountRepository.
func NewPgAppliedItemDiscountRepository(db DBExecutor) AppliedItemDiscountRepository {
	return &pgAppliedItemDiscountRepository{db: db}
}

// Create menyisipkan satu atau lebih diskon yang diterapkan pada sebuah item transaksi.
func (r *pgAppliedItemDiscountRepository) Create(ctx context.Context, itemDiscounts []models.AppliedItemDiscount) error {
	if len(itemDiscounts) == 0 {
		return nil // Tidak ada yang perlu disisipkan
	}

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
		return fmt.Errorf("gagal memulai transaksi untuk membuat applied item discounts: %w", err)
	}
	defer tx.Rollback()

	stmt, err := tx.PrepareContext(ctx, `
		INSERT INTO applied_item_discounts (transaction_item_id, discount_id, applied_discount_amount_on_item)
		VALUES ($1, $2, $3)`)
	if err != nil {
		return fmt.Errorf("gagal menyiapkan statement: %w", err)
	}
	defer stmt.Close()

	for _, id := range itemDiscounts {
		_, err := stmt.ExecContext(ctx, id.TransactionItemID, id.DiscountID, id.AppliedDiscountAmountOnItem)
		if err != nil {
			return fmt.Errorf("gagal menyisipkan applied item discount (transaction_item_id: %s, discount_id: %s): %w", id.TransactionItemID, id.DiscountID, err)
		}
	}

	if !isTx {
		return tx.Commit()
	}
	return nil
}

// ListByTransactionItemID mengambil semua diskon yang diterapkan pada item transaksi tertentu.
func (r *pgAppliedItemDiscountRepository) ListByTransactionItemID(ctx context.Context, transactionItemID uuid.UUID) ([]*models.AppliedItemDiscount, error) {
	query := `
		SELECT transaction_item_id, discount_id, applied_discount_amount_on_item
		FROM applied_item_discounts
		WHERE transaction_item_id = $1`
	rows, err := r.db.QueryContext(ctx, query, transactionItemID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var discounts []*models.AppliedItemDiscount
	for rows.Next() {
		d := &models.AppliedItemDiscount{}
		if err := rows.Scan(&d.TransactionItemID, &d.DiscountID, &d.AppliedDiscountAmountOnItem); err != nil {
			return nil, err
		}
		discounts = append(discounts, d)
	}
	return discounts, rows.Err()
}

// DeleteByTransactionItemID menghapus semua diskon yang terkait dengan item transaksi tertentu.
func (r *pgAppliedItemDiscountRepository) DeleteByTransactionItemID(ctx context.Context, transactionItemID uuid.UUID) error {
	query := `DELETE FROM applied_item_discounts WHERE transaction_item_id = $1`
	_, err := r.db.ExecContext(ctx, query, transactionItemID)
	return err
}
