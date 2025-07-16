package models

import (
	"time"

	"github.com/google/uuid"
)

type ImportResult struct {
	ID         uuid.UUID `db:"id" json:"id"`
	DataType   string    `db:"data_type" json:"data_type"`
	Success    bool      `db:"success" json:"success"`
	Message    string    `db:"message" json:"message"`
	ImportedAt time.Time `db:"imported_at" json:"imported_at"`
	UploadedBy uuid.UUID `db:"uploaded_by" json:"uploaded_by"`
}
