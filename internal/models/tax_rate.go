package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

// TaxRate merepresentasikan data tarif pajak dari tabel 'tax_rates'.
type TaxRate struct {
	ID             int32          `db:"id" json:"id"` // SERIAL di PostgreSQL
	CompanyID      uuid.UUID      `db:"company_id" json:"company_id"`
	Name           string         `db:"name" json:"name"`
	RatePercentage float64        `db:"rate_percentage" json:"rate_percentage"`   // DECIMAL(5,2)
	Description    sql.NullString `db:"description" json:"description,omitempty"` // Bisa NULL
	IsActive       bool           `db:"is_active" json:"is_active"`
	CreatedAt      time.Time      `db:"created_at" json:"created_at"`
	UpdatedAt      time.Time      `db:"updated_at" json:"updated_at"`
}
