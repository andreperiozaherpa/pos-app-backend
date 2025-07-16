package repository

import (
	"context"
	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

// StoreProductRepository mendefinisikan interface untuk operasi data terkait Product (store_product).
type StoreProductRepository interface {
	// Create membuat data store product baru.
	Create(ctx context.Context, product *models.StoreProduct) error

	// GetByID mengambil store product berdasarkan ID.
	GetByID(ctx context.Context, id uuid.UUID) (*models.StoreProduct, error)

	// GetByStoreAndMasterProduct mengambil store product berdasarkan store ID dan master product ID.
	GetByStoreAndMasterProduct(ctx context.Context, storeID, masterProductID uuid.UUID) (*models.StoreProduct, error)

	// ListByStore mengambil daftar store product berdasarkan store ID.
	ListByStore(ctx context.Context, storeID uuid.UUID) ([]*models.StoreProduct, error)

	// Update memperbarui data store product.
	Update(ctx context.Context, product *models.StoreProduct) error

	// Delete menghapus data store product berdasarkan ID.
	Delete(ctx context.Context, id uuid.UUID) error
}
