package services

import (
	"context"
	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

// CustomerService mendefinisikan kontrak use case pelanggan/member.
type CustomerService interface {
	// RegisterCustomer membuat pelanggan/member baru.
	RegisterCustomer(ctx context.Context, customer *models.Customer) (uuid.UUID, error)

	// GetCustomerByID mengambil data customer berdasarkan user ID.
	GetCustomerByID(ctx context.Context, userID uuid.UUID) (*models.Customer, error)

	// UpdateCustomer memperbarui data customer/member.
	UpdateCustomer(ctx context.Context, customer *models.Customer) error

	// DeleteCustomer menghapus data customer/member (soft delete recommended).
	DeleteCustomer(ctx context.Context, userID uuid.UUID) error

	// ListCustomersByCompanyID mengambil semua customer pada company tertentu.
	ListCustomersByCompanyID(ctx context.Context, companyID uuid.UUID) ([]*models.Customer, error)
}
