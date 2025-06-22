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

// EmployeeService mendefinisikan interface untuk logika bisnis terkait manajemen karyawan.
type EmployeeService interface {
	AddEmployee(ctx context.Context, employee *models.Employee, user *models.User, password string, roleIDs []int32) (*models.Employee, error)
	UpdateEmployee(ctx context.Context, employee *models.Employee) error
	DeactivateEmployee(ctx context.Context, employeeUserID uuid.UUID) error
	AssignRoleToEmployee(ctx context.Context, employeeUserID uuid.UUID, roleID int32) error
	RemoveRoleFromEmployee(ctx context.Context, employeeUserID uuid.UUID, roleID int32) error
	ListEmployees(ctx context.Context, companyID uuid.UUID) ([]*models.Employee, error)
}

// TransactionFunc adalah tipe fungsi yang berisi logika bisnis untuk dijalankan dalam satu transaksi.
// Ia menerima repository yang sudah terikat dengan transaksi tersebut.
type TransactionFunc func(userRepo postgres.UserRepository, employeeRepo postgres.EmployeeRepository, employeeRoleRepo postgres.EmployeeRoleRepository, roleRepo postgres.RoleRepository) error

// TransactionRunner adalah tipe fungsi yang bertanggung jawab untuk mengelola siklus hidup transaksi.
type TransactionRunner func(ctx context.Context, fn TransactionFunc) error

// NewTransactionRunner membuat implementasi TransactionRunner yang konkret menggunakan *sql.DB.
func NewTransactionRunner(db *sql.DB) TransactionRunner {
	return func(ctx context.Context, fn TransactionFunc) error {
		tx, err := db.BeginTx(ctx, nil)
		if err != nil {
			return fmt.Errorf("gagal memulai transaksi: %w", err)
		}
		defer tx.Rollback()

		// Buat instance repository yang terikat dengan transaksi
		userRepo := postgres.NewPgUserRepository(tx)
		employeeRepo := postgres.NewPgEmployeeRepository(tx)
		employeeRoleRepo := postgres.NewPgEmployeeRoleRepository(tx)
		roleRepo := postgres.NewPgRoleRepository(tx)

		if err := fn(userRepo, employeeRepo, employeeRoleRepo, roleRepo); err != nil {
			return err // Rollback akan dipanggil oleh defer
		}

		return tx.Commit()
	}
}

// employeeService adalah implementasi dari EmployeeService.
type employeeService struct {
	runInTransaction TransactionRunner // Mengelola transaksi
	userRepo         postgres.UserRepository
	employeeRepo     postgres.EmployeeRepository
	employeeRoleRepo postgres.EmployeeRoleRepository
	roleRepo         postgres.RoleRepository
}

// NewEmployeeService adalah constructor untuk membuat instance baru dari employeeService.
func NewEmployeeService(runner TransactionRunner, userRepo postgres.UserRepository, employeeRepo postgres.EmployeeRepository, employeeRoleRepo postgres.EmployeeRoleRepository, roleRepo postgres.RoleRepository) EmployeeService {
	return &employeeService{
		runInTransaction: runner,
		userRepo:         userRepo,
		employeeRepo:     employeeRepo,
		employeeRoleRepo: employeeRoleRepo,
		roleRepo:         roleRepo,
	}
}

// AddEmployee menambahkan karyawan baru, termasuk membuat data user dan menetapkan peran dalam satu transaksi.
func (s *employeeService) AddEmployee(ctx context.Context, employee *models.Employee, user *models.User, password string, roleIDs []int32) (*models.Employee, error) {
	transactionalLogic := func(userRepo postgres.UserRepository, employeeRepo postgres.EmployeeRepository, employeeRoleRepo postgres.EmployeeRoleRepository, roleRepo postgres.RoleRepository) error {
		// 1. Validasi bahwa peran yang diberikan ada
		for _, roleID := range roleIDs {
			if _, err := roleRepo.GetByID(ctx, roleID); err != nil {
				if err == sql.ErrNoRows {
					return fmt.Errorf("%w: dengan id %d", ErrRoleNotFound, roleID)
				}
				return err
			}
		}

		// 2. Buat data user
		if _, err := userRepo.GetByUsername(ctx, user.Username.String); err == nil {
			return ErrUsernameExists
		}
		if user.Email.Valid {
			if _, err := userRepo.GetByEmail(ctx, user.Email.String); err == nil {
				return ErrEmailExists
			}
		}

		hashedPassword, err := utils.HashPassword(password)
		if err != nil {
			return err
		}
		user.ID = uuid.New()
		user.PasswordHash = sql.NullString{String: hashedPassword, Valid: true}
		user.IsActive = true
		user.UserType = models.UserTypeEmployee
		user.CreatedAt = time.Now()
		user.UpdatedAt = time.Now()

		if err := userRepo.Create(ctx, user); err != nil {
			return err
		}

		// 3. Buat data employee yang terhubung dengan user
		employee.UserID = user.ID
		if err := employeeRepo.Create(ctx, employee); err != nil {
			return err
		}

		// 4. Tetapkan peran ke karyawan
		for _, roleID := range roleIDs {
			if err := employeeRoleRepo.AssignRoleToEmployee(ctx, user.ID, roleID); err != nil {
				return err
			}
		}
		return nil
	}

	if err := s.runInTransaction(ctx, transactionalLogic); err != nil {
		return nil, err
	}

	return employee, nil
}

// UpdateEmployee memperbarui data spesifik karyawan.
func (s *employeeService) UpdateEmployee(ctx context.Context, employee *models.Employee) error {
	return s.employeeRepo.Update(ctx, employee)
}

// DeactivateEmployee menonaktifkan akun karyawan (soft delete).
func (s *employeeService) DeactivateEmployee(ctx context.Context, employeeUserID uuid.UUID) error {
	return s.userRepo.Delete(employeeUserID)
}

// AssignRoleToEmployee menetapkan peran ke karyawan.
func (s *employeeService) AssignRoleToEmployee(ctx context.Context, employeeUserID uuid.UUID, roleID int32) error {
	return s.employeeRoleRepo.AssignRoleToEmployee(ctx, employeeUserID, roleID)
}

// RemoveRoleFromEmployee mencabut peran dari karyawan.
func (s *employeeService) RemoveRoleFromEmployee(ctx context.Context, employeeUserID uuid.UUID, roleID int32) error {
	return s.employeeRoleRepo.RemoveRoleFromEmployee(ctx, employeeUserID, roleID)
}

// ListEmployees mengambil daftar karyawan untuk sebuah perusahaan.
func (s *employeeService) ListEmployees(ctx context.Context, companyID uuid.UUID) ([]*models.Employee, error) {
	return s.employeeRepo.ListByCompanyID(ctx, companyID)
}
