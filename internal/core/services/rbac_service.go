package services

import (
	"context"
	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

// RBACService menangani logika akses Role-Based Access Control (gabungan role, permission, assignment).
type RBACService interface {
	AssignRoleToUser(ctx context.Context, userID uuid.UUID, roleID int) error
	AssignPermissionToRole(ctx context.Context, roleID int, permissionID int) error
	CheckUserPermission(ctx context.Context, userID uuid.UUID, permissionName string) (bool, error)
	RevokeRoleFromUser(ctx context.Context, userID uuid.UUID, roleID int) error
	RevokePermissionFromRole(ctx context.Context, roleID int, permissionID int) error
	ListAllRBACAssignments(ctx context.Context) ([]*models.RBACAssignment, error)
	ExportRBACConfig(ctx context.Context) ([]byte, error)
}
