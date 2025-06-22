package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

// Transaction merepresentasikan data transaksi dari tabel 'transactions'.
type Transaction struct {
	ID                             uuid.UUID         `db:"id"`
	TransactionCode                string            `db:"transaction_code"`
	StoreID                        uuid.UUID         `db:"store_id"`
	CashierEmployeeUserID          uuid.UUID         `db:"cashier_employee_user_id"`
	CustomerUserID                 uuid.NullUUID     `db:"customer_user_id"` // Bisa NULL
	ActiveShiftID                  uuid.NullUUID     `db:"active_shift_id"`  // Bisa NULL
	TransactionDate                time.Time         `db:"transaction_date"` // TIMESTAMPTZ
	SubtotalAmount                 float64           `db:"subtotal_amount"`
	TotalItemDiscountAmount        float64           `db:"total_item_discount_amount"`
	SubtotalAfterItemDiscounts     float64           `db:"subtotal_after_item_discounts"`
	TransactionLevelDiscountAmount float64           `db:"transaction_level_discount_amount"`
	TaxableAmount                  float64           `db:"taxable_amount"`
	TotalTaxAmount                 float64           `db:"total_tax_amount"`
	FinalTotalAmount               float64           `db:"final_total_amount"`
	ReceivedAmount                 float64           `db:"received_amount"`
	ChangeAmount                   float64           `db:"change_amount"`
	PaymentMethod                  sql.NullString    `db:"payment_method"` // Bisa NULL
	Notes                          sql.NullString    `db:"notes"`          // Bisa NULL
	CreatedAt                      time.Time         `db:"created_at"`
	UpdatedAt                      time.Time         `db:"updated_at"`
	Items                          []TransactionItem `db:"-"` // Digunakan untuk menampung item saat mengambil data, tidak ada di tabel transactions
}

// TransactionItem merepresentasikan data item dalam transaksi dari tabel 'transaction_items'.
type TransactionItem struct {
	ID                         uuid.UUID       `db:"id"`
	TransactionID              uuid.UUID       `db:"transaction_id"`
	StoreProductID             uuid.UUID       `db:"store_product_id"`
	Quantity                   int32           `db:"quantity"`
	PricePerUnitAtTransaction  float64         `db:"price_per_unit_at_transaction"`
	ItemSubtotalBeforeDiscount float64         `db:"item_subtotal_before_discount"`
	ItemDiscountAmount         float64         `db:"item_discount_amount"`
	ItemSubtotalAfterDiscount  float64         `db:"item_subtotal_after_discount"`
	AppliedTaxRateID           sql.NullInt32   `db:"applied_tax_rate_id"`         // Bisa NULL
	AppliedTaxRatePercentage   sql.NullFloat64 `db:"applied_tax_rate_percentage"` // Bisa NULL
	TaxAmountForItem           float64         `db:"tax_amount_for_item"`
	ItemFinalTotal             float64         `db:"item_final_total"`
	CreatedAt                  time.Time       `db:"created_at"`
	UpdatedAt                  time.Time       `db:"updated_at"`
}
