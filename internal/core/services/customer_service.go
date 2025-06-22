package services

import (
	"context"
	"database/sql"
	"fmt"
	"pos-app/backend/internal/data/postgres"
	"pos-app/backend/internal/models"
	"pos-app/backend/internal/utils"
	"time"

	"github.com/google/uuid"
)

// CustomerService mendefinisikan interface untuk logika bisnis terkait manajemen pelanggan.
type CustomerService interface {
	AddCustomer(ctx context.Context, customer *models.Customer, user *models.User, password string) (*models.Customer, error)
	UpdateCustomer(ctx context.Context, customer *models.Customer, user *models.User) error
	GetCustomerByPhoneNumber(ctx context.Context, phoneNumber string) (*models.Customer, error)
	UpdateCustomerPoints(ctx context.Context, customerUserID uuid.UUID, points int) error
}

// CustomerTransactionFunc adalah tipe fungsi yang berisi logika bisnis untuk dijalankan dalam satu transaksi
// khusus untuk operasi pelanggan.
type CustomerTransactionFunc func(userRepo postgres.UserRepository, customerRepo postgres.CustomerRepository) error

// CustomerTransactionRunner adalah tipe fungsi yang bertanggung jawab untuk mengelola siklus hidup transaksi
// untuk operasi pelanggan.
type CustomerTransactionRunner func(ctx context.Context, fn CustomerTransactionFunc) error

// NewCustomerTransactionRunner membuat implementasi CustomerTransactionRunner yang konkret menggunakan *sql.DB.
func NewCustomerTransactionRunner(db *sql.DB) CustomerTransactionRunner {
	return func(ctx context.Context, fn CustomerTransactionFunc) error {
		tx, err := db.BeginTx(ctx, nil)
		if err != nil {
			return fmt.Errorf("gagal memulai transaksi: %w", err)
		}
		defer tx.Rollback() // Pastikan rollback jika ada error

		// Buat instance repository yang terikat dengan transaksi
		userRepo := postgres.NewPgUserRepository(tx)
		customerRepo := postgres.NewPgCustomerRepository(tx)

		if err := fn(userRepo, customerRepo); err != nil {
			return err // Rollback akan dipanggil oleh defer
		}

		return tx.Commit()
	}
}

// customerService adalah implementasi dari CustomerService.
type customerService struct {
	runInTransaction CustomerTransactionRunner
	userRepo         postgres.UserRepository
	customerRepo     postgres.CustomerRepository
}

// NewCustomerService adalah constructor untuk membuat instance baru dari customerService.
func NewCustomerService(runner CustomerTransactionRunner, userRepo postgres.UserRepository, customerRepo postgres.CustomerRepository) CustomerService {
	return &customerService{
		runInTransaction: runner,
		userRepo:         userRepo,
		customerRepo:     customerRepo,
	}
}

// AddCustomer menambahkan pelanggan baru, termasuk membuat data user dalam satu transaksi.
func (s *customerService) AddCustomer(ctx context.Context, customer *models.Customer, user *models.User, password string) (*models.Customer, error) {
	transactionalLogic := func(userRepo postgres.UserRepository, customerRepo postgres.CustomerRepository) error {
		// 1. Validasi duplikasi username/email/phone
		if user.Username.Valid {
			if _, err := userRepo.GetByUsername(ctx, user.Username.String); err == nil {
				return ErrUsernameExists
			}
		}
		if user.Email.Valid {
			if _, err := userRepo.GetByEmail(ctx, user.Email.String); err == nil {
				return ErrEmailExists
			}
		}
		if user.PhoneNumber.Valid {
			if _, err := userRepo.GetByPhoneNumber(ctx, user.PhoneNumber.String); err == nil {
				return ErrPhoneNumberExists
			}
		}

		// 2. Buat data user
		hashedPassword, err := utils.HashPassword(password)
		if err != nil {
			return err
		}
		user.ID = uuid.New()
		user.PasswordHash = sql.NullString{String: hashedPassword, Valid: true}
		user.IsActive = true
		user.UserType = models.UserTypeCustomer
		user.CreatedAt = time.Now()
		user.UpdatedAt = time.Now()

		if err := userRepo.Create(ctx, user); err != nil {
			return err
		}

		// 3. Buat data customer yang terhubung dengan user
		customer.UserID = user.ID
		customer.CreatedAt = time.Now()
		customer.UpdatedAt = time.Now()
		if err := customerRepo.Create(ctx, customer); err != nil {
			return err
		}
		return nil
	}

	if err := s.runInTransaction(ctx, transactionalLogic); err != nil {
		return nil, err
	}

	return customer, nil
}

// UpdateCustomer memperbarui data spesifik pelanggan dan user terkait.
func (s *customerService) UpdateCustomer(ctx context.Context, customer *models.Customer, user *models.User) error {
	// Update user data first
	if err := s.userRepo.Update(ctx, user); err != nil {
		return err
	}
	// Then update customer data
	customer.UpdatedAt = time.Now()
	return s.customerRepo.Update(ctx, customer)
}

// GetCustomerByPhoneNumber mengambil data pelanggan berdasarkan nomor telepon.
func (s *customerService) GetCustomerByPhoneNumber(ctx context.Context, phoneNumber string) (*models.Customer, error) {
	user, err := s.userRepo.GetByPhoneNumber(ctx, phoneNumber)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrCustomerNotFound
		}
		return nil, err
	}
	customer, err := s.customerRepo.GetByUserID(ctx, user.ID)
	if err != nil {
		if err == sql.ErrNoRows { // Should not happen if user exists, but good to check
			return nil, ErrCustomerNotFound
		}
		return nil, err
	}
	customer.User = user // Attach user info to customer
	return customer, nil
}

// UpdateCustomerPoints memperbarui poin pelanggan.
func (s *customerService) UpdateCustomerPoints(ctx context.Context, customerUserID uuid.UUID, points int) error {
	customer, err := s.customerRepo.GetByUserID(ctx, customerUserID)
	if err != nil {
		if err == sql.ErrNoRows {
			return ErrCustomerNotFound
		}
		return err
	}
	customer.Points = int32(points)
	customer.UpdatedAt = time.Now()
	return s.customerRepo.Update(ctx, customer)
}
