package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

// Company merepresentasikan data perusahaan dari tabel 'companies'.
type Company struct {
	ID                   uuid.UUID       `db:"id"` // Menggunakan tag 'db' untuk mapping dengan sqlx (jika digunakan)
	Name                 string          `db:"name"`
	Address              sql.NullString  `db:"address"`      // Menggunakan sql.NullString karena kolom bisa NULL
	ContactInfo          []byte          `db:"contact_info"` // JSONB biasanya dibaca sebagai []byte atau string
	TaxIDNumber          sql.NullString  `db:"tax_id_number"`
	DefaultTaxPercentage sql.NullFloat64 `db:"default_tax_percentage"` // DECIMAL(5,2) yang bisa NULL
	CreatedAt            time.Time       `db:"created_at"`
	UpdatedAt            time.Time       `db:"updated_at"`
}
