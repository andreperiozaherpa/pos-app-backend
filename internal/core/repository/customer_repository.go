package repository

import (
	"context"
	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

// CustomerRepository mendefinisikan interface untuk operasi data terkait Customer.
type CustomerRepository interface {
	// Create membuat data customer baru.
	Create(ctx context.Context, customer *models.Customer) error

	// GetByUserID mengambil customer berdasarkan user ID.
	GetByUserID(ctx context.Context, userID uuid.UUID) (*models.Customer, error)

	// GetByMembershipNumber mengambil customer berdasarkan nomor membership dan perusahaan.
	GetByMembershipNumber(ctx context.Context, companyID uuid.UUID, membershipNumber string) (*models.Customer, error)

	// ListByCompanyID mengambil semua customer dalam satu perusahaan.
	ListByCompanyID(ctx context.Context, companyID uuid.UUID) ([]*models.Customer, error)

	// Update memperbarui data customer.
	Update(ctx context.Context, customer *models.Customer) error

	// Delete menghapus data customer berdasarkan user ID.
	Delete(ctx context.Context, userID uuid.UUID) error

	// GetLoyaltyPoints mendapatkan poin loyalti customer berdasarkan ID.
	GetLoyaltyPoints(ctx context.Context, customerID uuid.UUID) (int, error)

	// AddLoyaltyPoints menambahkan poin loyalti ke customer.
	AddLoyaltyPoints(ctx context.Context, customerID uuid.UUID, points int) error

	// ListTransactionHistory mengambil histori transaksi customer dengan pagination.
	ListTransactionHistory(ctx context.Context, customerID uuid.UUID, limit, offset int) ([]*models.Transaction, error)

	// SearchCustomers mencari customer berdasarkan nama, email, atau telepon dengan pagination.
	SearchCustomers(ctx context.Context, keyword string, limit, offset int) ([]*models.Customer, error)
}
