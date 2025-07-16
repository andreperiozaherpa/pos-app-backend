package models

import (
	"time"

	"github.com/google/uuid"
)

// ExportHistory merepresentasikan histori proses export data (produk, transaksi, laporan, dsb) dari sistem.
type ExportHistory struct {
	ID          uuid.UUID `db:"id" json:"id"`
	ExportType  string    `db:"export_type" json:"export_type"`   // Jenis export: produk, transaksi, laporan, dsb
	FileName    string    `db:"file_name" json:"file_name"`       // Nama file hasil export
	Status      string    `db:"status" json:"status"`             // Sukses, Gagal, Partial
	ExportedBy  uuid.UUID `db:"exported_by" json:"exported_by"`   // User yang melakukan export
	ExportedAt  time.Time `db:"exported_at" json:"exported_at"`   // Waktu export
	TotalRows   int       `db:"total_rows" json:"total_rows"`     // Jumlah baris diexport
	SuccessRows int       `db:"success_rows" json:"success_rows"` // Jumlah baris sukses
	FailedRows  int       `db:"failed_rows" json:"failed_rows"`   // Jumlah baris gagal
	FileURL     string    `db:"file_url" json:"file_url"`         // Link file hasil export (jika disimpan di storage)
	Notes       string    `db:"notes" json:"notes"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}
