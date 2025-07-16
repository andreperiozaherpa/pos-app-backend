package models

import (
	"time"

	"github.com/google/uuid"
)

// ContactHistory merepresentasikan riwayat komunikasi dengan customer atau supplier.
type ContactHistory struct {
	ID                uuid.UUID  `db:"id" json:"id"`
	RelatedUserID     *uuid.UUID `db:"related_user_id" json:"related_user_id"`         // FK ke User (Customer/Supplier), nullable
	RelatedSupplierID *uuid.UUID `db:"related_supplier_id" json:"related_supplier_id"` // FK ke Supplier, nullable
	ContactType       string     `db:"contact_type" json:"contact_type"`               // Jenis: phone, email, wa, sms, dsb
	ContactDate       time.Time  `db:"contact_date" json:"contact_date"`
	Subject           string     `db:"subject" json:"subject"`       // Subjek komunikasi
	Content           string     `db:"content" json:"content"`       // Isi pesan/catatan
	HandledBy         *uuid.UUID `db:"handled_by" json:"handled_by"` // User yang handle (employee), nullable
	Notes             string     `db:"notes" json:"notes"`
	CreatedAt         time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt         time.Time  `db:"updated_at" json:"updated_at"`
}
