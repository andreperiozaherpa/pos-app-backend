package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

// StoreType mendefinisikan tipe toko.
type StoreType string

// Konstanta untuk StoreType
const (
	StoreTypePusat   StoreType = "PUSAT"
	StoreTypeCabang  StoreType = "CABANG"
	StoreTypeRanting StoreType = "RANTING"
)

// Store merepresentasikan data toko dari tabel 'stores'.
type Store struct {
	ID             uuid.UUID      `db:"id" json:"id"`
	BusinessLineID uuid.UUID      `db:"business_line_id" json:"business_line_id"`
	ParentStoreID  uuid.NullUUID  `db:"parent_store_id" json:"parent_store_id,omitempty"` // Bisa NULL
	Name           string         `db:"name" json:"name"`
	StoreCode      sql.NullString `db:"store_code" json:"store_code,omitempty"`     // Bisa NULL
	StoreType      StoreType      `db:"store_type" json:"store_type"`               // Enum: PUSAT, CABANG, RANTING
	Address        sql.NullString `db:"address" json:"address,omitempty"`           // Bisa NULL
	PhoneNumber    sql.NullString `db:"phone_number" json:"phone_number,omitempty"` // Bisa NULL
	CreatedAt      time.Time      `db:"created_at" json:"created_at"`
	UpdatedAt      time.Time      `db:"updated_at" json:"updated_at"`
}
