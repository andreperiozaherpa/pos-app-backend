package models

import (
	"time"

	"github.com/google/uuid"
)

// ImportHistory merepresentasikan histori proses import data (produk, customer, dsb) ke sistem.
type ImportHistory struct {
	ID           uuid.UUID `db:"id" json:"id"`
	ImportType   string    `db:"import_type" json:"import_type"` // Tipe data diimport (produk, customer, supplier, dst.)
	FileName     string    `db:"file_name" json:"file_name"`     // Nama file asal import
	Status       string    `db:"status" json:"status"`           // Sukses, Gagal, Partial
	ImportedBy   uuid.UUID `db:"imported_by" json:"imported_by"` // User yang melakukan import
	ImportedAt   time.Time `db:"imported_at" json:"imported_at"`
	TotalRows    int       `db:"total_rows" json:"total_rows"`         // Jumlah baris dalam file
	SuccessRows  int       `db:"success_rows" json:"success_rows"`     // Jumlah baris sukses
	FailedRows   int       `db:"failed_rows" json:"failed_rows"`       // Jumlah baris gagal
	ErrorFileURL string    `db:"error_file_url" json:"error_file_url"` // Link file error hasil validasi (jika ada)
	Notes        string    `db:"notes" json:"notes"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time `db:"updated_at" json:"updated_at"`
}
