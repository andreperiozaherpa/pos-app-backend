package repository

import (
	"context"
	"pos-app/backend/internal/models" // Import package models yang sudah kita buat

	"github.com/google/uuid"
)

// DiscountRepository mendefinisikan interface untuk operasi data terkait Discount.
type DiscountRepository interface {
	// Create membuat data discount baru.
	Create(ctx context.Context, discount *models.Discount) error

	// GetByID mengambil data discount berdasarkan ID.
	GetByID(ctx context.Context, id uuid.UUID) (*models.Discount, error)

	// ListByCompanyID mengambil daftar discount berdasarkan company ID.
	ListByCompanyID(ctx context.Context, companyID uuid.UUID) ([]*models.Discount, error)

	// Update memperbarui data discount.
	Update(ctx context.Context, discount *models.Discount) error

	// Delete menghapus data discount berdasarkan ID.
	Delete(ctx context.Context, id uuid.UUID) error
}
