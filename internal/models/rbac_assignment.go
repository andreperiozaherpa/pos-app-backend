package models

import (
	"time"

	"github.com/google/uuid"
)

// RBACAssignment merepresentasikan assignment role dan permission ke user pada sistem RBAC.
type RBACAssignment struct {
	ID           uuid.UUID  `db:"id" json:"id"`
	UserID       uuid.UUID  `db:"user_id" json:"user_id"`             // User yang mendapat assignment
	RoleID       *int       `db:"role_id" json:"role_id"`             // Nullable (jika hanya permission)
	PermissionID *int       `db:"permission_id" json:"permission_id"` // Nullable (jika hanya role)
	AssignedBy   uuid.UUID  `db:"assigned_by" json:"assigned_by"`     // Siapa yang assign
	AssignedAt   time.Time  `db:"assigned_at" json:"assigned_at"`
	RevokedBy    *uuid.UUID `db:"revoked_by" json:"revoked_by"` // Nullable, jika assignment dicabut
	RevokedAt    *time.Time `db:"revoked_at" json:"revoked_at"` // Nullable
	Status       string     `db:"status" json:"status"`         // Active, Revoked
	Notes        string     `db:"notes" json:"notes"`
	CreatedAt    time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time  `db:"updated_at" json:"updated_at"`
}
