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

	// SearchCustomers melakukan pencarian customer secara fleksibel berdasarkan nama, email, atau nomor telepon.
	SearchCustomers(ctx context.Context, companyID uuid.UUID, query string) ([]*models.Customer, error)

	// DeactivateCustomer menonaktifkan (soft delete/blacklist) customer.
	DeactivateCustomer(ctx context.Context, userID uuid.UUID) error

	// ListCustomerTransactions mengambil daftar transaksi yang dilakukan customer.
	ListCustomerTransactions(ctx context.Context, userID uuid.UUID) ([]*models.Transaction, error)

	// GetCustomerLoyaltyPoints mengambil jumlah poin loyalitas customer.
	GetCustomerLoyaltyPoints(ctx context.Context, userID uuid.UUID) (int, error)

	// UpdateCustomerLoyaltyPoints memperbarui (menambah/mengurangi) poin loyalitas customer.
	UpdateCustomerLoyaltyPoints(ctx context.Context, userID uuid.UUID, delta int) error

	// ExportCustomers mengekspor data customer ke file (excel/CSV).
	ExportCustomers(ctx context.Context, companyID uuid.UUID) ([]byte, error)

	// BulkImportCustomers melakukan import massal data customer dari file excel/CSV.
	BulkImportCustomers(ctx context.Context, companyID uuid.UUID, fileData []byte) error

	// GetCustomerContactHistory mengambil riwayat komunikasi dengan customer (call/email/wa).
	GetCustomerContactHistory(ctx context.Context, userID uuid.UUID) ([]*models.ContactHistory, error)
}
