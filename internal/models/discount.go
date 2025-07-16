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
	ID                        uuid.UUID            `db:"id" json:"id"`
	CompanyID                 uuid.UUID            `db:"company_id" json:"company_id"`
	Name                      string               `db:"name" json:"name"`
	Description               sql.NullString       `db:"description" json:"description,omitempty"`
	DiscountType              DiscountType         `db:"discount_type" json:"discount_type"`
	DiscountValue             float64              `db:"discount_value" json:"discount_value"`
	ApplicableTo              DiscountApplicableTo `db:"applicable_to" json:"applicable_to"`
	MasterProductIDApplicable uuid.NullUUID        `db:"master_product_id_applicable" json:"master_product_id_applicable,omitempty"`
	StoreProductIDApplicable  uuid.NullUUID        `db:"store_product_id_applicable" json:"store_product_id_applicable,omitempty"`
	CategoryApplicable        sql.NullString       `db:"category_applicable" json:"category_applicable,omitempty"`
	CustomerTierApplicable    sql.NullString       `db:"customer_tier_applicable" json:"customer_tier_applicable,omitempty"`
	MinPurchaseAmount         sql.NullFloat64      `db:"min_purchase_amount" json:"min_purchase_amount,omitempty"`
	StartDate                 time.Time            `db:"start_date" json:"start_date"`
	EndDate                   time.Time            `db:"end_date" json:"end_date"`
	IsActive                  bool                 `db:"is_active" json:"is_active"`
	CreatedAt                 time.Time            `db:"created_at" json:"created_at"`
	UpdatedAt                 time.Time            `db:"updated_at" json:"updated_at"`
}
