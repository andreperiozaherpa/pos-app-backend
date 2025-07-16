package repository

import (
	"context"
	"pos-app/backend/internal/models"
)

// PermissionRepository mendefinisikan interface untuk operasi data terkait Permission.
type PermissionRepository interface {
	// Create membuat data permission baru.
	Create(ctx context.Context, permission *models.Permission) error

	// GetByID mengambil data permission berdasarkan ID.
	GetByID(ctx context.Context, id int32) (*models.Permission, error)

	// ListAll mengambil semua data permission.
	ListAll(ctx context.Context) ([]*models.Permission, error)

	// Update memperbarui data permission.
	Update(ctx context.Context, permission *models.Permission) error

	// Delete menghapus data permission berdasarkan ID.
	Delete(ctx context.Context, id int32) error
}
