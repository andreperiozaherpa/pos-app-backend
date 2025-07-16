package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

// UserType mendefinisikan tipe pengguna (Employee atau Customer).
type UserType string

// Konstanta UserType untuk membedakan jenis user di aplikasi.
const (
	UserTypeEmployee UserType = "EMPLOYEE"
	UserTypeCustomer UserType = "CUSTOMER"
)

// User merepresentasikan data pengguna dari tabel 'users'.
//
// Catatan:
// - Field PasswordHash tidak pernah di-expose ke JSON/API (json:"-")
// - Field nullable menggunakan sql.NullString (untuk integrasi DB yang mendukung nilai NULL)
type User struct {
	ID           uuid.UUID      `db:"id" json:"id"`                               // Primary key (UUID)
	UserType     UserType       `db:"user_type" json:"user_type"`                 // EMPLOYEE / CUSTOMER
	Username     sql.NullString `db:"username" json:"username,omitempty"`         // Username login (boleh kosong)
	PasswordHash sql.NullString `db:"password_hash" json:"-"`                     // Hash password, tidak pernah keluar di API
	FullName     sql.NullString `db:"full_name" json:"full_name,omitempty"`       // Nama lengkap user
	Email        sql.NullString `db:"email" json:"email,omitempty"`               // Email user
	PhoneNumber  sql.NullString `db:"phone_number" json:"phone_number,omitempty"` // Nomor HP user
	IsActive     bool           `db:"is_active" json:"is_active"`                 // Status aktif/nonaktif
	CreatedAt    time.Time      `db:"created_at" json:"created_at"`               // Timestamp buat
	UpdatedAt    time.Time      `db:"updated_at" json:"updated_at"`               // Timestamp update terakhir
}
