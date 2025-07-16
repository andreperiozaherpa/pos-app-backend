package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

// Company merepresentasikan data perusahaan dari tabel 'companies'.
type Company struct {
	ID                   uuid.UUID       `db:"id" json:"id"`                                                   // PK, UUID
	Name                 string          `db:"name" json:"name"`                                               // Nama perusahaan
	Address              sql.NullString  `db:"address" json:"address,omitempty"`                               // Alamat, nullable
	ContactInfo          []byte          `db:"contact_info" json:"contact_info,omitempty"`                     // JSONB, nullable
	TaxIDNumber          sql.NullString  `db:"tax_id_number" json:"tax_id_number,omitempty"`                   // NPWP/Tax ID, nullable
	DefaultTaxPercentage sql.NullFloat64 `db:"default_tax_percentage" json:"default_tax_percentage,omitempty"` // Default tax %
	CreatedAt            time.Time       `db:"created_at" json:"created_at"`
	UpdatedAt            time.Time       `db:"updated_at" json:"updated_at"`
}
