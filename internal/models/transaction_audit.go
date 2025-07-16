package models

import (
	"time"

	"github.com/google/uuid"
)

// TransactionAuditTrail merekam histori perubahan data pada transaksi (audit trail).
type TransactionAuditTrail struct {
	ID            uuid.UUID `db:"id" json:"id"`
	TransactionID uuid.UUID `db:"transaction_id" json:"transaction_id"` // FK ke transaksi utama
	ChangedBy     uuid.UUID `db:"changed_by" json:"changed_by"`         // User yang melakukan perubahan
	ChangeTime    time.Time `db:"change_time" json:"change_time"`       // Waktu perubahan
	ChangeType    string    `db:"change_type" json:"change_type"`       // Jenis aksi: create, update, refund, cancel, dsb
	FieldChanged  string    `db:"field_changed" json:"field_changed"`   // Field yang diubah (optional)
	OldValue      string    `db:"old_value" json:"old_value"`           // Nilai lama (JSON/text)
	NewValue      string    `db:"new_value" json:"new_value"`           // Nilai baru (JSON/text)
	Notes         string    `db:"notes" json:"notes"`                   // Keterangan tambahan
	CreatedAt     time.Time `db:"created_at" json:"created_at"`
}
