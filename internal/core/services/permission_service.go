package services

import (
	"context"
	"pos-app/backend/internal/models"
)

// PermissionService mendefinisikan kontrak use case untuk permission (hak akses).
type PermissionService interface {
	// CreatePermission membuat permission baru.
	CreatePermission(ctx context.Context, permission *models.Permission) (int64, error)

	// GetPermissionByID mengambil permission berdasarkan ID.
	GetPermissionByID(ctx context.Context, id int64) (*models.Permission, error)

	// UpdatePermission memperbarui data permission.
	UpdatePermission(ctx context.Context, permission *models.Permission) error

	// DeletePermission menghapus permission berdasarkan ID.
	DeletePermission(ctx context.Context, id int64) error

	// ListPermissions mengambil semua permission yang tersedia.
	ListPermissions(ctx context.Context) ([]*models.Permission, error)
}
