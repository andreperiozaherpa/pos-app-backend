package models

import (
	"time"

	"github.com/google/uuid"
)

type MasterProductHistory struct {
	ID              uuid.UUID `db:"id" json:"id"`
	MasterProductID uuid.UUID `db:"master_product_id" json:"master_product_id"`
	ChangedBy       uuid.UUID `db:"changed_by" json:"changed_by"`
	ChangeType      string    `db:"change_type" json:"change_type"`
	ChangedAt       time.Time `db:"changed_at" json:"changed_at"`
	Notes           string    `db:"notes" json:"notes"`
}
