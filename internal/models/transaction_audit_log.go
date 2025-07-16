package models

import (
	"time"

	"github.com/google/uuid"
)

type TransactionAuditLog struct {
	ID                uuid.UUID `db:"id" json:"id"`
	TransactionID     uuid.UUID `db:"transaction_id" json:"transaction_id"`
	ActionType        string    `db:"action_type" json:"action_type"`
	PerformedByUserID uuid.UUID `db:"performed_by_user_id" json:"performed_by_user_id"`
	PerformedAt       time.Time `db:"performed_at" json:"performed_at"`
	Note              string    `db:"note" json:"note"`
}
