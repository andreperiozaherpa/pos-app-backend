package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

// DiscountType mendefinisikan tipe diskon.
type DiscountType string

// Konstanta untuk DiscountType
const (
	DiscountTypePercentage  DiscountType = "PERCENTAGE"
	DiscountTypeFixedAmount DiscountType = "FIXED_AMOUNT"
)

// DiscountApplicableTo mendefinisikan target penerapan diskon.
type DiscountApplicableTo string

// Konstanta untuk DiscountApplicableTo
const (
	DiscountApplicableToMasterProduct    DiscountApplicableTo = "MASTER_PRODUCT"
	DiscountApplicableToStoreProduct     DiscountApplicableTo = "STORE_PRODUCT"
	DiscountApplicableToCategory         DiscountApplicableTo = "CATEGORY"
	DiscountApplicableToTotalTransaction DiscountApplicableTo = "TOTAL_TRANSACTION"
	DiscountApplicableToCustomerTier     DiscountApplicableTo = "CUSTOMER_TIER"
)

// Discount merepresentasikan data diskon dari tabel 'discounts'.
type Discount struct {
	ID                        uuid.UUID            `db:"id"`
	CompanyID                 uuid.UUID            `db:"company_id"`
	Name                      string               `db:"name"`
	Description               sql.NullString       `db:"description"`
	DiscountType              DiscountType         `db:"discount_type"`
	DiscountValue             float64              `db:"discount_value"` // DECIMAL
	ApplicableTo              DiscountApplicableTo `db:"applicable_to"`
	MasterProductIDApplicable uuid.NullUUID        `db:"master_product_id_applicable"`
	StoreProductIDApplicable  uuid.NullUUID        `db:"store_product_id_applicable"`
	CategoryApplicable        sql.NullString       `db:"category_applicable"`
	CustomerTierApplicable    sql.NullString       `db:"customer_tier_applicable"`
	MinPurchaseAmount         sql.NullFloat64      `db:"min_purchase_amount"` // DECIMAL bisa NULL
	StartDate                 time.Time            `db:"start_date"`          // DATE
	EndDate                   time.Time            `db:"end_date"`            // DATE
	IsActive                  bool                 `db:"is_active"`
	CreatedAt                 time.Time            `db:"created_at"`
	UpdatedAt                 time.Time            `db:"updated_at"`
}
