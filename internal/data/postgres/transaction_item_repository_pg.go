package postgres

import (
	"context"
	"database/sql"
	"errors"

	"pos-app/backend/internal/core/repository"
	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

type transactionItemRepositoryPG struct {
	db *sql.DB
}

func NewTransactionItemRepositoryPG(db *sql.DB) repository.TransactionItemRepository {
	return &transactionItemRepositoryPG{db: db}
}

// Create menambahkan item transaksi baru ke database sesuai ERD
func (r *transactionItemRepositoryPG) Create(ctx context.Context, item *models.TransactionItem) error {
	query := `
        INSERT INTO transaction_items (
			id, transaction_id, store_product_id, quantity,
			price_per_unit_at_transaction, item_subtotal_before_discount,
			item_discount_amount, item_subtotal_after_discount,
			applied_tax_rate_id, applied_tax_rate_percentage,
			tax_amount_for_item, item_final_total,
			created_at, updated_at
		) VALUES (
			$1, $2, $3, $4,
			$5, $6, $7, $8, $9, $10,
			$11, $12, NOW(), NOW()
		)`
	_, err := r.db.ExecContext(ctx, query,
		item.ID, item.TransactionID, item.StoreProductID, item.Quantity,
		item.PricePerUnitAtTransaction, item.ItemSubtotalBeforeDiscount,
		item.ItemDiscountAmount, item.ItemSubtotalAfterDiscount,
		item.AppliedTaxRateID, item.AppliedTaxRatePercentage,
		item.TaxAmountForItem, item.ItemFinalTotal,
	)
	return err
}

func (r *transactionItemRepositoryPG) GetByID(ctx context.Context, id uuid.UUID) (*models.TransactionItem, error) {
	query := `
        SELECT id, transaction_id, store_product_id, quantity,
			price_per_unit_at_transaction, item_subtotal_before_discount,
			item_discount_amount, item_subtotal_after_discount,
			applied_tax_rate_id, applied_tax_rate_percentage,
			tax_amount_for_item, item_final_total,
			created_at, updated_at
        FROM transaction_items WHERE id = $1`
	row := r.db.QueryRowContext(ctx, query, id)
	item := &models.TransactionItem{}
	err := row.Scan(
		&item.ID, &item.TransactionID, &item.StoreProductID, &item.Quantity,
		&item.PricePerUnitAtTransaction, &item.ItemSubtotalBeforeDiscount,
		&item.ItemDiscountAmount, &item.ItemSubtotalAfterDiscount,
		&item.AppliedTaxRateID, &item.AppliedTaxRatePercentage,
		&item.TaxAmountForItem, &item.ItemFinalTotal,
		&item.CreatedAt, &item.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}
	return item, nil
}

// ListByTransactionID mengambil semua item berdasarkan ID transaksi sesuai ERD
func (r *transactionItemRepositoryPG) ListByTransactionID(ctx context.Context, transactionID uuid.UUID) ([]*models.TransactionItem, error) {
	query := `
        SELECT id, transaction_id, store_product_id, quantity,
			price_per_unit_at_transaction, item_subtotal_before_discount,
			item_discount_amount, item_subtotal_after_discount,
			applied_tax_rate_id, applied_tax_rate_percentage,
			tax_amount_for_item, item_final_total,
			created_at, updated_at
        FROM transaction_items
        WHERE transaction_id = $1
        ORDER BY created_at`
	rows, err := r.db.QueryContext(ctx, query, transactionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []*models.TransactionItem
	for rows.Next() {
		item := &models.TransactionItem{}
		if err := rows.Scan(
			&item.ID, &item.TransactionID, &item.StoreProductID, &item.Quantity,
			&item.PricePerUnitAtTransaction, &item.ItemSubtotalBeforeDiscount,
			&item.ItemDiscountAmount, &item.ItemSubtotalAfterDiscount,
			&item.AppliedTaxRateID, &item.AppliedTaxRatePercentage,
			&item.TaxAmountForItem, &item.ItemFinalTotal,
			&item.CreatedAt, &item.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func (r *transactionItemRepositoryPG) Update(ctx context.Context, item *models.TransactionItem) error {
	query := `
        UPDATE transaction_items SET
			transaction_id = $1,
			store_product_id = $2,
			quantity = $3,
			price_per_unit_at_transaction = $4,
			item_subtotal_before_discount = $5,
			item_discount_amount = $6,
			item_subtotal_after_discount = $7,
			applied_tax_rate_id = $8,
			applied_tax_rate_percentage = $9,
			tax_amount_for_item = $10,
			item_final_total = $11,
			updated_at = NOW()
        WHERE id = $12`
	result, err := r.db.ExecContext(ctx, query,
		item.TransactionID, item.StoreProductID, item.Quantity,
		item.PricePerUnitAtTransaction, item.ItemSubtotalBeforeDiscount,
		item.ItemDiscountAmount, item.ItemSubtotalAfterDiscount,
		item.AppliedTaxRateID, item.AppliedTaxRatePercentage,
		item.TaxAmountForItem, item.ItemFinalTotal, item.ID,
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

func (r *transactionItemRepositoryPG) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM transaction_items WHERE id = $1`
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
