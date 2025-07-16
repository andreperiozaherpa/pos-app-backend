package models

import (
	"time"

	"github.com/google/uuid"
)

// AuthSession merepresentasikan sesi autentikasi user, termasuk token, device, dan status sesi.
type AuthSession struct {
	ID           uuid.UUID  `db:"id" json:"id"`
	UserID       uuid.UUID  `db:"user_id" json:"user_id"`             // FK ke User
	AccessToken  string     `db:"access_token" json:"access_token"`   // Token aktif
	RefreshToken string     `db:"refresh_token" json:"refresh_token"` // Token refresh
	DeviceInfo   string     `db:"device_info" json:"device_info"`     // Info perangkat/browser
	IPAddress    string     `db:"ip_address" json:"ip_address"`       // Alamat IP login
	LoginTime    time.Time  `db:"login_time" json:"login_time"`       // Waktu login
	LogoutTime   *time.Time `db:"logout_time" json:"logout_time"`     // Waktu logout, nullable
	ExpiredAt    time.Time  `db:"expired_at" json:"expired_at"`       // Kapan sesi/tokens expired
	Status       string     `db:"status" json:"status"`               // Aktif, logout, expired, revoked
	CreatedAt    time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time  `db:"updated_at" json:"updated_at"`
}
