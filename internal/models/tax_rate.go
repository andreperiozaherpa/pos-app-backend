package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

// TaxRate merepresentasikan data tarif pajak dari tabel 'tax_rates'.
type TaxRate struct {
	ID             int32          `db:"id"` // SERIAL di PostgreSQL
	CompanyID      uuid.UUID      `db:"company_id"`
	Name           string         `db:"name"`
	RatePercentage float64        `db:"rate_percentage"` // DECIMAL(5,2)
	Description    sql.NullString `db:"description"`     // Bisa NULL
	IsActive       bool           `db:"is_active"`
	CreatedAt      time.Time      `db:"created_at"`
	UpdatedAt      time.Time      `db:"updated_at"`
}
