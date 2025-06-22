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

// employeeService adalah implementasi dari EmployeeService.
type employeeService struct {
	db               *sql.DB // Diperlukan untuk memulai transaksi
	userRepo         postgres.UserRepository
	employeeRepo     postgres.EmployeeRepository
	employeeRoleRepo postgres.EmployeeRoleRepository
	roleRepo         postgres.RoleRepository
}

// NewEmployeeService adalah constructor untuk membuat instance baru dari employeeService.
func NewEmployeeService(db *sql.DB, userRepo postgres.UserRepository, employeeRepo postgres.EmployeeRepository, employeeRoleRepo postgres.EmployeeRoleRepository, roleRepo postgres.RoleRepository) EmployeeService {
	return &employeeService{
		db:               db,
		userRepo:         userRepo,
		employeeRepo:     employeeRepo,
		employeeRoleRepo: employeeRoleRepo,
		roleRepo:         roleRepo,
	}
}

// AddEmployee menambahkan karyawan baru, termasuk membuat data user dan menetapkan peran dalam satu transaksi.
func (s *employeeService) AddEmployee(ctx context.Context, employee *models.Employee, user *models.User, password string, roleIDs []int32) (*models.Employee, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("gagal memulai transaksi: %w", err)
	}
	defer tx.Rollback() // Rollback jika terjadi error

	// Buat instance repository yang terikat dengan transaksi
	txUserRepo := postgres.NewPgUserRepository(tx)
	txEmployeeRepo := postgres.NewPgEmployeeRepository(tx)
	txEmployeeRoleRepo := postgres.NewPgEmployeeRoleRepository(tx)
	txRoleRepo := postgres.NewPgRoleRepository(tx)

	// 1. Validasi bahwa peran yang diberikan ada
	for _, roleID := range roleIDs {
		if _, err := txRoleRepo.GetByID(ctx, roleID); err != nil {
			if err == sql.ErrNoRows {
				return nil, fmt.Errorf("%w: dengan id %d", ErrRoleNotFound, roleID)
			}
			return nil, err
		}
	}

	// 2. Buat data user
	// Validasi duplikasi username/email
	if _, err := txUserRepo.GetByUsername(ctx, user.Username.String); err == nil {
		return nil, ErrUsernameExists
	}
	if user.Email.Valid {
		if _, err := txUserRepo.GetByEmail(ctx, user.Email.String); err == nil {
			return nil, ErrEmailExists
		}
	}

	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}
	user.ID = uuid.New()
	user.PasswordHash = sql.NullString{String: hashedPassword, Valid: true}
	user.IsActive = true
	user.UserType = models.UserTypeEmployee
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	if err := txUserRepo.Create(ctx, user); err != nil {
		return nil, err
	}

	// 3. Buat data employee yang terhubung dengan user
	employee.UserID = user.ID
	if err := txEmployeeRepo.Create(ctx, employee); err != nil {
		return nil, err
	}

	// 4. Tetapkan peran ke karyawan
	for _, roleID := range roleIDs {
		empRole := &models.EmployeeRole{EmployeeUserID: user.ID, RoleID: roleID}
		if err := txEmployeeRoleRepo.Create(ctx, empRole); err != nil {
			return nil, err
		}
	}

	// Commit transaksi jika semua operasi berhasil
	return employee, tx.Commit()
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
	return s.employeeRoleRepo.Create(ctx, &models.EmployeeRole{EmployeeUserID: employeeUserID, RoleID: roleID})
}

// RemoveRoleFromEmployee mencabut peran dari karyawan.
func (s *employeeService) RemoveRoleFromEmployee(ctx context.Context, employeeUserID uuid.UUID, roleID int32) error {
	return s.employeeRoleRepo.Delete(ctx, employeeUserID, roleID)
}

// ListEmployees mengambil daftar karyawan untuk sebuah perusahaan.
func (s *employeeService) ListEmployees(ctx context.Context, companyID uuid.UUID) ([]*models.Employee, error) {
	return s.employeeRepo.ListByCompanyID(ctx, companyID)
}
