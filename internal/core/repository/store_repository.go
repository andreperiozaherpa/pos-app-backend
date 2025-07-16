package repository

import (
	"context"
	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

// StoreRepository mendefinisikan interface untuk operasi data terkait Store.
type StoreRepository interface {
	// Create membuat data store baru.
	Create(ctx context.Context, store *models.Store) error

	// GetByID mengambil data store berdasarkan ID.
	GetByID(ctx context.Context, id uuid.UUID) (*models.Store, error)

	// Update memperbarui data store.
	Update(ctx context.Context, store *models.Store) error

	// Delete menghapus data store berdasarkan ID.
	Delete(ctx context.Context, id uuid.UUID) error

	// ListByBusinessLine mengambil daftar store berdasarkan business line ID.
	ListByBusinessLine(ctx context.Context, businessLineID uuid.UUID) ([]*models.Store, error)
}
