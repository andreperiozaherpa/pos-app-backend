package models

import (
	"time"

	"github.com/google/uuid"
)

// MasterProductVariant merepresentasikan varian produk pusat (misal: warna, ukuran, dll).
type MasterProductVariant struct {
	ID              uuid.UUID `db:"id" json:"id"`
	MasterProductID uuid.UUID `db:"master_product_id" json:"master_product_id"` // FK ke produk pusat
	VariantCode     string    `db:"variant_code" json:"variant_code"`           // Kode varian, unik per produk pusat
	Name            string    `db:"name" json:"name"`                           // Nama varian (misal: Merah, XL, Pedas)
	Description     string    `db:"description" json:"description"`
	CreatedAt       time.Time `db:"created_at" json:"created_at"`
	UpdatedAt       time.Time `db:"updated_at" json:"updated_at"`
}
