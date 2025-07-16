package models

import (
	"time"

	"github.com/google/uuid"
)

type UserLoginHistory struct {
	ID         uuid.UUID `db:"id" json:"id"`
	UserID     uuid.UUID `db:"user_id" json:"user_id"`
	LoginTime  time.Time `db:"login_time" json:"login_time"`
	IPAddress  string    `db:"ip_address" json:"ip_address"`
	DeviceInfo string    `db:"device_info" json:"device_info"`
}
