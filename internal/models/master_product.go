package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

// MasterProduct merepresentasikan data produk master dari tabel 'master_products'.
type MasterProduct struct {
	ID                uuid.UUID      `db:"id"`
	CompanyID         uuid.UUID      `db:"company_id"`
	MasterProductCode string         `db:"master_product_code"`
	Name              string         `db:"name"`
	Description       sql.NullString `db:"description"`
	Category          sql.NullString `db:"category"`
	UnitOfMeasure     sql.NullString `db:"unit_of_measure"`
	Barcode           sql.NullString `db:"barcode"`
	DefaultTaxRateID  sql.NullInt32  `db:"default_tax_rate_id"` // INTEGER bisa NULL, jadi pakai sql.NullInt32
	ImageURL          sql.NullString `db:"image_url"`
	CreatedAt         time.Time      `db:"created_at"`
	UpdatedAt         time.Time      `db:"updated_at"`
}
