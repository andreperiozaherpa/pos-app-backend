package models

import (
	"time"

	"github.com/google/uuid"
)

// Notification merepresentasikan pesan/notifikasi ke user (push/email/SMS/dll).
type Notification struct {
	ID        uuid.UUID  `db:"id" json:"id"`
	UserID    uuid.UUID  `db:"user_id" json:"user_id"`     // Penerima notifikasi (FK ke user)
	Title     string     `db:"title" json:"title"`         // Judul notifikasi
	Message   string     `db:"message" json:"message"`     // Isi notifikasi
	Channel   string     `db:"channel" json:"channel"`     // email, sms, whatsapp, in-app, dll
	Status    string     `db:"status" json:"status"`       // Sent, Failed, Pending, Read, Unread
	SentAt    *time.Time `db:"sent_at" json:"sent_at"`     // Waktu terkirim (nullable)
	ReadAt    *time.Time `db:"read_at" json:"read_at"`     // Waktu dibaca (nullable)
	Reference string     `db:"reference" json:"reference"` // Referensi terkait transaksi/order, dll
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt time.Time  `db:"updated_at" json:"updated_at"`
}
