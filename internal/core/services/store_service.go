package services

import (
	"context"
	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

// StoreService menangani manajemen data store/toko per business line atau company.
type StoreService interface {
	// CRUD utama
	CreateStore(ctx context.Context, store *models.Store) (uuid.UUID, error)
	GetStoreByID(ctx context.Context, id uuid.UUID) (*models.Store, error)
	UpdateStore(ctx context.Context, store *models.Store) error
	DeleteStore(ctx context.Context, id uuid.UUID) error
	ListStoresByBusinessLineID(ctx context.Context, businessLineID uuid.UUID) ([]*models.Store, error)

	// Opsional & Custom
	ListStoresByCompanyID(ctx context.Context, companyID uuid.UUID) ([]*models.Store, error)
	ActivateStore(ctx context.Context, id uuid.UUID) error
	DeactivateStore(ctx context.Context, id uuid.UUID) error
	ExportStores(ctx context.Context, companyID uuid.UUID) ([]byte, error)
	ArchiveStore(ctx context.Context, id uuid.UUID) error
	RestoreStore(ctx context.Context, id uuid.UUID) error
	ListStoreEmployees(ctx context.Context, storeID uuid.UUID) ([]*models.Employee, error)
}
