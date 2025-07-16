package services

import (
	"context"
	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

// EmployeeService mendefinisikan kontrak use case terkait karyawan (employee).
type EmployeeService interface {
	// RegisterEmployee membuat employee baru dan relasi ke user.
	RegisterEmployee(ctx context.Context, employee *models.Employee) (uuid.UUID, error)

	// GetEmployeeByID mengambil data employee berdasarkan user ID.
	GetEmployeeByID(ctx context.Context, userID uuid.UUID) (*models.Employee, error)

	// UpdateEmployee memperbarui data employee.
	UpdateEmployee(ctx context.Context, employee *models.Employee) error

	// DeleteEmployee menghapus data employee (soft delete recommended).
	DeleteEmployee(ctx context.Context, userID uuid.UUID) error

	// ListEmployeesByCompanyID mengambil daftar employee berdasarkan company ID.
	ListEmployeesByCompanyID(ctx context.Context, companyID uuid.UUID) ([]*models.Employee, error)

	// ListEmployeesByStoreID mengambil daftar employee berdasarkan store ID.
	ListEmployeesByStoreID(ctx context.Context, storeID uuid.UUID) ([]*models.Employee, error)
	// AssignEmployeeToStore menetapkan employee ke store tertentu.
	AssignEmployeeToStore(ctx context.Context, userID uuid.UUID, storeID uuid.UUID) error

	// UpdateEmployeeStatus memperbarui status kerja employee (aktif/nonaktif/cuti).
	UpdateEmployeeStatus(ctx context.Context, userID uuid.UUID, status string) error

	// GetEmployeeAttendance mengambil data absensi employee pada periode tertentu.
	GetEmployeeAttendance(ctx context.Context, userID uuid.UUID, from, to string) ([]*models.EmployeeAttendance, error)

	// ListEmployeeRoles menampilkan daftar role yang dimiliki employee.
	ListEmployeeRoles(ctx context.Context, userID uuid.UUID) ([]*models.Role, error)

	// ListEmployeeAttendanceByDateRange mengambil absensi employee dalam rentang tanggal.
	ListEmployeeAttendanceByDateRange(ctx context.Context, userID uuid.UUID, from, to string) ([]*models.EmployeeAttendance, error)

	// GetEmployeeLeaveHistory menampilkan riwayat cuti employee.
	GetEmployeeLeaveHistory(ctx context.Context, userID uuid.UUID) ([]*models.EmployeeLeave, error)

	// AssignEmployeeRole memberikan role ke employee.
	AssignEmployeeRole(ctx context.Context, userID uuid.UUID, roleID int) error

	// RemoveEmployeeFromStore menghapus penugasan employee dari store.
	RemoveEmployeeFromStore(ctx context.Context, userID uuid.UUID, storeID uuid.UUID) error

	// GetEmployeePerformanceSummary menampilkan ringkasan kinerja employee.
	GetEmployeePerformanceSummary(ctx context.Context, userID uuid.UUID, from, to string) (*models.EmployeePerformanceSummary, error)

	// ExportEmployeeData mengekspor data employee ke format excel/CSV.
	ExportEmployeeData(ctx context.Context, companyID uuid.UUID) ([]byte, error)
}
