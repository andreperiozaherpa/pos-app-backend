package models

import (
	"github.com/google/uuid"
)

// AppliedItemDiscount merepresentasikan data diskon yang diterapkan pada item transaksi.
// Ini adalah tabel join antara transaction_items dan discounts.
type AppliedItemDiscount struct {
	TransactionItemID           uuid.UUID `db:"transaction_item_id" json:"transaction_item_id"`
	DiscountID                  uuid.UUID `db:"discount_id" json:"discount_id"`
	AppliedDiscountAmountOnItem float64   `db:"applied_discount_amount_on_item" json:"applied_discount_amount_on_item"`
}

// AppliedTransactionDiscount merepresentasikan data diskon yang diterapkan pada keseluruhan transaksi.
// Ini adalah tabel join antara transactions dan discounts.
type AppliedTransactionDiscount struct {
	TransactionID                      uuid.UUID `db:"transaction_id" json:"transaction_id"`
	DiscountID                         uuid.UUID `db:"discount_id" json:"discount_id"`
	AppliedDiscountAmountOnTransaction float64   `db:"applied_discount_amount_on_transaction" json:"applied_discount_amount_on_transaction"`
}
