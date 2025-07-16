package models

import (
	"time"

	"github.com/google/uuid"
)

// PurchaseOrderHistory menyimpan histori perubahan untuk Purchase Order.
type PurchaseOrderHistory struct {
	ID              uuid.UUID `db:"id" json:"id"`                               // ID unik untuk histori
	PurchaseOrderID uuid.UUID `db:"purchase_order_id" json:"purchase_order_id"` // ID dari Purchase Order yang terkait
	Status          string    `db:"status" json:"status"`                       // Status dari Purchase Order (misalnya: "pending", "approved", "rejected")
	ChangedBy       uuid.UUID `db:"changed_by" json:"changed_by"`               // ID pengguna yang melakukan perubahan status
	ChangedAt       time.Time `db:"changed_at" json:"changed_at"`               // Waktu perubahan status
	Notes           string    `db:"notes" json:"notes"`                         // Catatan tambahan terkait perubahan
}
