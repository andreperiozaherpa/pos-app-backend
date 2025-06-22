package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

// UserType mendefinisikan tipe pengguna (Employee atau Customer).
type UserType string

// Konstanta untuk UserType
const (
	UserTypeEmployee UserType = "EMPLOYEE"
	UserTypeCustomer UserType = "CUSTOMER"
)

// User merepresentasikan data pengguna dari tabel 'users'.
type User struct {
	ID           uuid.UUID      `db:"id"`
	UserType     UserType       `db:"user_type"`     // Menggunakan tipe kustom UserType
	Username     sql.NullString `db:"username"`      // Bisa NULL
	PasswordHash sql.NullString `db:"password_hash"` // Bisa NULL (untuk customer tanpa password)
	FullName     sql.NullString `db:"full_name"`     // Bisa NULL
	Email        sql.NullString `db:"email"`         // Bisa NULL
	PhoneNumber  sql.NullString `db:"phone_number"`  // Bisa NULL
	IsActive     bool           `db:"is_active"`
	CreatedAt    time.Time      `db:"created_at"`
	UpdatedAt    time.Time      `db:"updated_at"`
}
