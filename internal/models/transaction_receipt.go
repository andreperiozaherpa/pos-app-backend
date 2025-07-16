package models

import (
	"time"

	"github.com/google/uuid"
)

// TransactionReceipt merepresentasikan data struk transaksi (print/export), termasuk metadata print/akses.
type TransactionReceipt struct {
	ID            uuid.UUID `db:"id" json:"id"`
	TransactionID uuid.UUID `db:"transaction_id" json:"transaction_id"` // FK ke transaksi utama
	ReceiptNumber string    `db:"receipt_number" json:"receipt_number"` // Nomor/serial struk (bisa sama dengan kode transaksi, atau unik)
	PrintedBy     uuid.UUID `db:"printed_by" json:"printed_by"`         // User yang mencetak/ekspor
	PrintDate     time.Time `db:"print_date" json:"print_date"`
	PrintCount    int       `db:"print_count" json:"print_count"`     // Berapa kali dicetak
	PrintType     string    `db:"print_type" json:"print_type"`       // Asli, copy, reprint
	ExportFormat  string    `db:"export_format" json:"export_format"` // PDF, Excel, PNG, dsb
	ReceiptURL    string    `db:"receipt_url" json:"receipt_url"`     // Jika ada file/URL hasil ekspor
	Notes         string    `db:"notes" json:"notes"`
	CreatedAt     time.Time `db:"created_at" json:"created_at"`
	UpdatedAt     time.Time `db:"updated_at" json:"updated_at"`
}
