package repository

import (
	"context"
	"pos-app/backend/internal/models"
)

// RoleRepository mendefinisikan interface untuk operasi data terkait Role.
type RoleRepository interface {
	// Create membuat data role baru.
	Create(ctx context.Context, role *models.Role) error

	// GetByID mengambil data role berdasarkan ID.
	GetByID(ctx context.Context, id int32) (*models.Role, error)

	// GetByName mengambil data role berdasarkan nama.
	GetByName(ctx context.Context, name string) (*models.Role, error)

	// ListAll mengambil semua data role.
	ListAll(ctx context.Context) ([]*models.Role, error)

	// Update memperbarui data role.
	Update(ctx context.Context, role *models.Role) error

	// Delete menghapus data role berdasarkan ID.
	Delete(ctx context.Context, id int32) error
}
