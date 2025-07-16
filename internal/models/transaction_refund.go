package models

import (
	"time"

	"github.com/google/uuid"
)

// TransactionRefund merepresentasikan data refund (pengembalian dana) pada transaksi POS.
type TransactionRefund struct {
	ID            uuid.UUID  `db:"id" json:"id"`
	TransactionID uuid.UUID  `db:"transaction_id" json:"transaction_id"` // FK ke transaksi utama
	RefundedBy    uuid.UUID  `db:"refunded_by" json:"refunded_by"`       // User/kasir yang memproses refund
	RefundDate    time.Time  `db:"refund_date" json:"refund_date"`
	RefundAmount  float64    `db:"refund_amount" json:"refund_amount"` // Jumlah uang yang dikembalikan
	RefundReason  string     `db:"refund_reason" json:"refund_reason"` // Alasan/refund note
	ApprovedBy    *uuid.UUID `db:"approved_by" json:"approved_by"`     // User yang menyetujui, nullable
	RefundMethod  string     `db:"refund_method" json:"refund_method"` // Tunai, transfer, e-wallet, dsb
	Status        string     `db:"status" json:"status"`               // Pending, Approved, Rejected, Selesai
	Notes         string     `db:"notes" json:"notes"`
	CreatedAt     time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt     time.Time  `db:"updated_at" json:"updated_at"`
}
