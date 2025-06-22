package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

// ActivityLog merepresentasikan data log aktivitas dari tabel 'activity_logs'.
type ActivityLog struct {
	ID             int64          `db:"id"` // BIGSERIAL di PostgreSQL
	UserID         uuid.NullUUID  `db:"user_id"`
	CompanyID      uuid.NullUUID  `db:"company_id"`
	StoreID        uuid.NullUUID  `db:"store_id"`
	ActionType     string         `db:"action_type"`
	Description    string         `db:"description"`
	TargetEntity   sql.NullString `db:"target_entity"`
	TargetEntityID uuid.NullUUID  `db:"target_entity_id"` // Menggunakan uuid.NullUUID
	IPAddress      sql.NullString `db:"ip_address"`
	UserAgent      sql.NullString `db:"user_agent"`
	LogTime        time.Time      `db:"log_time"` // TIMESTAMPTZ
}
