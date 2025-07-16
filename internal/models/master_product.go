package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

// MasterProduct merepresentasikan data produk master dari tabel 'master_products'.
type MasterProduct struct {
	ID                uuid.UUID      `db:"id" json:"id"`
	CompanyID         uuid.UUID      `db:"company_id" json:"company_id"`
	MasterProductCode string         `db:"master_product_code" json:"master_product_code"`
	Name              string         `db:"name" json:"name"`
	Description       sql.NullString `db:"description" json:"description,omitempty"`
	Category          sql.NullString `db:"category" json:"category,omitempty"`
	UnitOfMeasure     sql.NullString `db:"unit_of_measure" json:"unit_of_measure,omitempty"`
	Barcode           sql.NullString `db:"barcode" json:"barcode,omitempty"`
	DefaultTaxRateID  sql.NullInt32  `db:"default_tax_rate_id" json:"default_tax_rate_id,omitempty"`
	ImageURL          sql.NullString `db:"image_url" json:"image_url,omitempty"`
	CreatedAt         time.Time      `db:"created_at" json:"created_at"`
	UpdatedAt         time.Time      `db:"updated_at" json:"updated_at"`
}
