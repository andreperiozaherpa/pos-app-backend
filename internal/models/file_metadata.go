package models

import (
	"time"

	"github.com/google/uuid"
)

// FileMetadata menyimpan metadata file yang di-upload/disimpan di aplikasi (foto, dokumen, hasil export/import, dll).
type FileMetadata struct {
	ID          uuid.UUID  `db:"id" json:"id"`
	OwnerUserID *uuid.UUID `db:"owner_user_id" json:"owner_user_id"` // User yang upload (nullable jika sistem)
	RelatedType string     `db:"related_type" json:"related_type"`   // Jenis file terkait (produk, transaksi, dsb)
	RelatedID   *uuid.UUID `db:"related_id" json:"related_id"`       // FK ke entitas terkait (opsional/nullable)
	FileName    string     `db:"file_name" json:"file_name"`         // Nama file asli
	FileURL     string     `db:"file_url" json:"file_url"`           // Lokasi file di storage/URL
	MimeType    string     `db:"mime_type" json:"mime_type"`         // Jenis MIME file
	FileSize    int64      `db:"file_size" json:"file_size"`         // Ukuran file (byte)
	Status      string     `db:"status" json:"status"`               // Active, Archived, Deleted
	UploadedAt  time.Time  `db:"uploaded_at" json:"uploaded_at"`     // Waktu upload
	ArchivedAt  *time.Time `db:"archived_at" json:"archived_at"`     // Jika diarsipkan (nullable)
	DeletedAt   *time.Time `db:"deleted_at" json:"deleted_at"`       // Jika dihapus (nullable)
	CreatedAt   time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time  `db:"updated_at" json:"updated_at"`
}
