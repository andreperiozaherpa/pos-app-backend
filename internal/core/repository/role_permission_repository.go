package repository

import (
	"context"
	"pos-app/backend/internal/models"
)

// RolePermissionRepository mendefinisikan interface untuk operasi data terkait RolePermission.
type RolePermissionRepository interface {
	// Create membuat data role-permission baru.
	Create(ctx context.Context, rp *models.RolePermission) error

	// ListByRoleID mengambil daftar role-permission berdasarkan role ID.
	ListByRoleID(ctx context.Context, roleID int32) ([]*models.RolePermission, error)

	// ListByPermissionID mengambil daftar role-permission berdasarkan permission ID.
	ListByPermissionID(ctx context.Context, permissionID int32) ([]*models.RolePermission, error)

	// Delete menghapus data role-permission berdasarkan role ID dan permission ID.
	Delete(ctx context.Context, roleID int32, permissionID int32) error

	// DeleteByRoleID menghapus semua data role-permission berdasarkan role ID.
	DeleteByRoleID(ctx context.Context, roleID int32) error
}
