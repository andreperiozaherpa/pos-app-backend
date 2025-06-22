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
	ID             uuid.UUID      `db:"id"`
	BusinessLineID uuid.UUID      `db:"business_line_id"`
	ParentStoreID  uuid.NullUUID  `db:"parent_store_id"` // Menggunakan uuid.NullUUID karena bisa NULL
	Name           string         `db:"name"`
	StoreCode      sql.NullString `db:"store_code"`   // Bisa NULL
	StoreType      StoreType      `db:"store_type"`   // Menggunakan tipe kustom StoreType
	Address        sql.NullString `db:"address"`      // Bisa NULL
	PhoneNumber    sql.NullString `db:"phone_number"` // Bisa NULL
	CreatedAt      time.Time      `db:"created_at"`
	UpdatedAt      time.Time      `db:"updated_at"`
}
