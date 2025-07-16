package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

// Transaction merepresentasikan data transaksi dari tabel 'transactions'.
type Transaction struct {
	ID                             uuid.UUID         `db:"id" json:"id"`
	TransactionCode                string            `db:"transaction_code" json:"transaction_code"`
	StoreID                        uuid.UUID         `db:"store_id" json:"store_id"`
	CashierEmployeeUserID          uuid.UUID         `db:"cashier_employee_user_id" json:"cashier_employee_user_id"`
	CustomerUserID                 uuid.NullUUID     `db:"customer_user_id" json:"customer_user_id,omitempty"` // Bisa NULL
	ActiveShiftID                  uuid.NullUUID     `db:"active_shift_id" json:"active_shift_id,omitempty"`   // Bisa NULL
	TransactionDate                time.Time         `db:"transaction_date" json:"transaction_date"`
	SubtotalAmount                 float64           `db:"subtotal_amount" json:"subtotal_amount"`
	TotalItemDiscountAmount        float64           `db:"total_item_discount_amount" json:"total_item_discount_amount"`
	SubtotalAfterItemDiscounts     float64           `db:"subtotal_after_item_discounts" json:"subtotal_after_item_discounts"`
	TransactionLevelDiscountAmount float64           `db:"transaction_level_discount_amount" json:"transaction_level_discount_amount"`
	TaxableAmount                  float64           `db:"taxable_amount" json:"taxable_amount"`
	TotalTaxAmount                 float64           `db:"total_tax_amount" json:"total_tax_amount"`
	FinalTotalAmount               float64           `db:"final_total_amount" json:"final_total_amount"`
	ReceivedAmount                 float64           `db:"received_amount" json:"received_amount"`
	ChangeAmount                   float64           `db:"change_amount" json:"change_amount"`
	PaymentMethod                  sql.NullString    `db:"payment_method" json:"payment_method,omitempty"` // Bisa NULL
	Notes                          sql.NullString    `db:"notes" json:"notes,omitempty"`                   // Bisa NULL
	CreatedAt                      time.Time         `db:"created_at" json:"created_at"`
	UpdatedAt                      time.Time         `db:"updated_at" json:"updated_at"`
	Items                          []TransactionItem `db:"-" json:"items,omitempty"` // Digunakan untuk menampung item saat mengambil data, tidak ada di tabel transactions
}

// TransactionItem merepresentasikan data item dalam transaksi dari tabel 'transaction_items'.
type TransactionItem struct {
	ID                         uuid.UUID       `db:"id" json:"id"`
	TransactionID              uuid.UUID       `db:"transaction_id" json:"transaction_id"`
	StoreProductID             uuid.UUID       `db:"store_product_id" json:"store_product_id"`
	Quantity                   int32           `db:"quantity" json:"quantity"`
	PricePerUnitAtTransaction  float64         `db:"price_per_unit_at_transaction" json:"price_per_unit_at_transaction"`
	ItemSubtotalBeforeDiscount float64         `db:"item_subtotal_before_discount" json:"item_subtotal_before_discount"`
	ItemDiscountAmount         float64         `db:"item_discount_amount" json:"item_discount_amount"`
	ItemSubtotalAfterDiscount  float64         `db:"item_subtotal_after_discount" json:"item_subtotal_after_discount"`
	AppliedTaxRateID           sql.NullInt32   `db:"applied_tax_rate_id" json:"applied_tax_rate_id,omitempty"`                 // Bisa NULL
	AppliedTaxRatePercentage   sql.NullFloat64 `db:"applied_tax_rate_percentage" json:"applied_tax_rate_percentage,omitempty"` // Bisa NULL
	TaxAmountForItem           float64         `db:"tax_amount_for_item" json:"tax_amount_for_item"`
	ItemFinalTotal             float64         `db:"item_final_total" json:"item_final_total"`
	CreatedAt                  time.Time       `db:"created_at" json:"created_at"`
	UpdatedAt                  time.Time       `db:"updated_at" json:"updated_at"`
}
