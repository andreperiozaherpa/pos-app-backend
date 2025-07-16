package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

// ActivityLog merepresentasikan data log aktivitas dari tabel 'activity_logs'.
type ActivityLog struct {
	ID             int64          `db:"id" json:"id"`
	UserID         uuid.NullUUID  `db:"user_id" json:"user_id,omitempty"`
	CompanyID      uuid.NullUUID  `db:"company_id" json:"company_id,omitempty"`
	StoreID        uuid.NullUUID  `db:"store_id" json:"store_id,omitempty"`
	ActionType     string         `db:"action_type" json:"action_type"`
	Description    string         `db:"description" json:"description"`
	TargetEntity   sql.NullString `db:"target_entity" json:"target_entity,omitempty"`
	TargetEntityID uuid.NullUUID  `db:"target_entity_id" json:"target_entity_id,omitempty"`
	IPAddress      sql.NullString `db:"ip_address" json:"ip_address,omitempty"`
	UserAgent      sql.NullString `db:"user_agent" json:"user_agent,omitempty"`
	LogTime        time.Time      `db:"log_time" json:"log_time"`
}
