package postgres

import (
	"context"
	"database/sql"
	"errors"

	"pos-app/backend/internal/core/repository"
	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

type employeeRepositoryPG struct {
	db *sql.DB
}

// NewEmployeeRepositoryPG membuat instance baru EmployeeRepository berbasis PostgreSQL.
func NewEmployeeRepositoryPG(db *sql.DB) repository.EmployeeRepository {
	return &employeeRepositoryPG{db: db}
}

// Create menambahkan employee baru ke database.
func (r *employeeRepositoryPG) Create(ctx context.Context, employee *models.Employee) error {
	query := `
		INSERT INTO employees 
		(user_id, company_id, store_id, employee_id_number, join_date, position, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, NOW(), NOW())
	`
	_, err := r.db.ExecContext(ctx, query,
		employee.UserID, employee.CompanyID, employee.StoreID, employee.EmployeeIDNumber,
		employee.JoinDate, employee.Position,
	)
	return err
}

// GetByID mengambil data employee berdasarkan ID.
func (r *employeeRepositoryPG) GetByID(ctx context.Context, userID uuid.UUID) (*models.Employee, error) {
	query := `
		SELECT user_id, company_id, store_id, employee_id_number, join_date, position, created_at, updated_at
		FROM employees WHERE user_id = $1
	`
	row := r.db.QueryRowContext(ctx, query, userID)
	employee := new(models.Employee)
	err := row.Scan(
		&employee.UserID, &employee.CompanyID, &employee.StoreID, &employee.EmployeeIDNumber,
		&employee.JoinDate, &employee.Position, &employee.CreatedAt, &employee.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}
	return employee, nil
}

// Update memperbarui data employee di database.
func (r *employeeRepositoryPG) Update(ctx context.Context, employee *models.Employee) error {
	query := `
		UPDATE employees
		SET company_id = $2, store_id = $3, name = $4, email = $5, phone = $6, position = $7, status = $8, updated_at = NOW()
		WHERE id = $1
	`
	result, err := r.db.ExecContext(ctx, query,
		employee.CompanyID, employee.StoreID,
		employee.Position)
	if err != nil {
		return err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return repository.ErrNotFound
	}
	return nil
}

// Delete menghapus employee berdasarkan ID.
func (r *employeeRepositoryPG) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM employees WHERE id = $1`
	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return repository.ErrNotFound
	}
	return nil
}

// ListEmployeesByCompanyID mengambil daftar employee berdasarkan company ID dengan pagination.
func (r *employeeRepositoryPG) ListEmployeesByCompanyID(ctx context.Context, companyID uuid.UUID, limit, offset int) ([]*models.Employee, error) {
	query := `
		SELECT id, company_id, store_id, name, email, phone, position, status, created_at, updated_at
		FROM employees WHERE company_id = $1
		ORDER BY name
		LIMIT $2 OFFSET $3
	`
	rows, err := r.db.QueryContext(ctx, query, companyID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var employees []*models.Employee
	for rows.Next() {
		employee := new(models.Employee)
		err := rows.Scan(
			&employee.CompanyID, &employee.StoreID,
			&employee.Position,
			&employee.CreatedAt, &employee.UpdatedAt)
		if err != nil {
			return nil, err
		}
		employees = append(employees, employee)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return employees, nil
}

// GetEmployeeAttendance - contoh stub, implementasi menyesuaikan schema dan kebutuhan
func (r *employeeRepositoryPG) GetEmployeeAttendance(ctx context.Context, employeeID uuid.UUID, date string) ([]*models.EmployeeAttendance, error) {
	// Implementasi query absensi sesuai tabel employee_attendance
	// Ini hanya contoh, silakan sesuaikan sesuai struktur tabel dan kebutuhan
	query := `
		SELECT id, employee_id, attendance_date, check_in, check_out, status, created_at, updated_at
		FROM employee_attendance
		WHERE employee_id = $1 AND attendance_date = $2
	`
	rows, err := r.db.QueryContext(ctx, query, employeeID, date)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var attendances []*models.EmployeeAttendance
	for rows.Next() {
		attendance := new(models.EmployeeAttendance)
		err := rows.Scan(
			&attendance.ID, &attendance.EmployeeUserID, &attendance.AttendanceDate,
			&attendance.CheckInTime, &attendance.CheckOutTime, &attendance.AttendanceStatus,
			&attendance.CreatedAt, &attendance.UpdatedAt)
		if err != nil {
			return nil, err
		}
		attendances = append(attendances, attendance)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return attendances, nil
}

// GetByEmployeeIDNumber mengambil data employee berdasarkan nomor identitas karyawan.
func (r *employeeRepositoryPG) GetByEmployeeIDNumber(ctx context.Context, employeeIDNumber string) (*models.Employee, error) {
	query := `
		SELECT user_id, company_id, store_id, employee_id_number, join_date, position, created_at, updated_at
		FROM employees WHERE employee_id_number = $1
	`
	row := r.db.QueryRowContext(ctx, query, employeeIDNumber)
	employee := new(models.Employee)
	err := row.Scan(
		&employee.UserID, &employee.CompanyID, &employee.StoreID, &employee.EmployeeIDNumber,
		&employee.JoinDate, &employee.Position, &employee.CreatedAt, &employee.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}
	return employee, nil
}

// ListByCompanyID mengambil daftar employee berdasarkan companyID dengan pagination.
func (r *employeeRepositoryPG) ListByCompanyID(ctx context.Context, companyID uuid.UUID, limit, offset int) ([]*models.Employee, error) {
	query := `
		SELECT user_id, company_id, store_id, employee_id_number, join_date, position, created_at, updated_at
		FROM employees
		WHERE company_id = $1
		ORDER BY join_date DESC
		LIMIT $2 OFFSET $3
	`
	rows, err := r.db.QueryContext(ctx, query, companyID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var employees []*models.Employee
	for rows.Next() {
		employee := new(models.Employee)
		err := rows.Scan(
			&employee.UserID,
			&employee.CompanyID,
			&employee.StoreID,
			&employee.EmployeeIDNumber,
			&employee.JoinDate,
			&employee.Position,
			&employee.CreatedAt,
			&employee.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		employees = append(employees, employee)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return employees, nil
}

// ListByStoreID mengambil semua employee pada toko tertentu.
func (r *employeeRepositoryPG) ListByStoreID(ctx context.Context, storeID uuid.UUID) ([]*models.Employee, error) {
	query := `
		SELECT id, company_id, store_id, name, email, phone, position, status, created_at, updated_at
		FROM employees WHERE store_id = $1
		ORDER BY name
	`
	rows, err := r.db.QueryContext(ctx, query, storeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var employees []*models.Employee
	for rows.Next() {
		employee := new(models.Employee)
		err := rows.Scan(
			&employee.CompanyID, &employee.StoreID, &employee.Position,
			&employee.CreatedAt, &employee.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		employees = append(employees, employee)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return employees, nil
}
