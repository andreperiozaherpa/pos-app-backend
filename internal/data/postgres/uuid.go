package postgres

import (
	"github.com/google/uuid"
)

// NewUUID mengembalikan UUID baru versi 4.
func NewUUID() uuid.UUID {
	return uuid.New()
}

// ParseUUID mengubah string menjadi UUID, mengembalikan error jika gagal.
func ParseUUID(s string) (uuid.UUID, error) {
	return uuid.Parse(s)
}

// IsValidUUID memeriksa apakah string valid sebagai UUID.
func IsValidUUID(s string) bool {
	_, err := uuid.Parse(s)
	return err == nil
}
