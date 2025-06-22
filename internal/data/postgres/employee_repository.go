package postgres

import (
	"context"
	"database/sql"
	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

// EmployeeRepository mendefinisikan interface untuk operasi data terkait Employee.
type EmployeeRepository interface {
	Create(ctx context.Context, employee *models.Employee) error
	// Primary key di tabel employees adalah user_id, jadi kita Get berdasarkan itu.
	GetByUserID(ctx context.Context, userID uuid.UUID) (*models.Employee, error)
	Update(ctx context.Context, employee *models.Employee) error
	Delete(ctx context.Context, userID uuid.UUID) error
	// Metode ini akan berguna untuk menampilkan semua karyawan di satu toko.
	ListByStoreID(ctx context.Context, storeID uuid.UUID) ([]*models.Employee, error)
	ListByCompanyID(ctx context.Context, companyID uuid.UUID) ([]*models.Employee, error)
}

// pgEmployeeRepository adalah implementasi dari EmployeeRepository untuk PostgreSQL.
type pgEmployeeRepository struct {
	db DBExecutor
}

// NewPgEmployeeRepository adalah constructor untuk membuat instance baru dari pgEmployeeRepository.
func NewPgEmployeeRepository(db DBExecutor) EmployeeRepository {
	return &pgEmployeeRepository{db: db}
}

// Implementasi metode-metode dari interface EmployeeRepository:

func (r *pgEmployeeRepository) Create(ctx context.Context, employee *models.Employee) error {
	query := `
		INSERT INTO employees (user_id, company_id, store_id, employee_id_number, join_date, position, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	_, err := r.db.ExecContext(ctx, query,
		employee.UserID,
		employee.CompanyID,
		employee.StoreID,
		employee.EmployeeIDNumber,
		employee.JoinDate,
		employee.Position,
		employee.CreatedAt,
		employee.UpdatedAt,
	)
	return err
}

func (r *pgEmployeeRepository) GetByUserID(ctx context.Context, userID uuid.UUID) (*models.Employee, error) {
	employee := &models.Employee{}
	query := `
		SELECT user_id, company_id, store_id, employee_id_number, join_date, position, created_at, updated_at
		FROM employees
		WHERE user_id = $1`

	err := r.db.QueryRowContext(ctx, query, userID).Scan(
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
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, err
	}
	return employee, nil
}

func (r *pgEmployeeRepository) Update(ctx context.Context, employee *models.Employee) error {
	query := `
		UPDATE employees
		SET company_id = $1, store_id = $2, employee_id_number = $3, join_date = $4, position = $5, updated_at = $6
		WHERE user_id = $7`

	_, err := r.db.ExecContext(ctx, query,
		employee.CompanyID,
		employee.StoreID,
		employee.EmployeeIDNumber,
		employee.JoinDate,
		employee.Position,
		employee.UpdatedAt,
		employee.UserID,
	)
	return err
}

func (r *pgEmployeeRepository) Delete(ctx context.Context, userID uuid.UUID) error {
	// Ini adalah hard delete. Soft delete akan dilakukan pada tabel 'users' (is_active = false).
	query := `DELETE FROM employees WHERE user_id = $1`
	_, err := r.db.ExecContext(ctx, query, userID)
	return err
}

func (r *pgEmployeeRepository) ListByStoreID(ctx context.Context, storeID uuid.UUID) ([]*models.Employee, error) {
	query := `
		SELECT user_id, company_id, store_id, employee_id_number, join_date, position, created_at, updated_at
		FROM employees
		WHERE store_id = $1
		ORDER BY created_at DESC`

	rows, err := r.db.QueryContext(ctx, query, storeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var employees []*models.Employee
	for rows.Next() {
		employee := &models.Employee{}
		if err := rows.Scan(
			&employee.UserID,
			&employee.CompanyID,
			&employee.StoreID,
			&employee.EmployeeIDNumber,
			&employee.JoinDate,
			&employee.Position,
			&employee.CreatedAt,
			&employee.UpdatedAt,
		); err != nil {
			return nil, err
		}
		employees = append(employees, employee)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return employees, nil
}

func (r *pgEmployeeRepository) ListByCompanyID(ctx context.Context, companyID uuid.UUID) ([]*models.Employee, error) {
	query := `
		SELECT user_id, company_id, store_id, employee_id_number, join_date, position, created_at, updated_at
		FROM employees
		WHERE company_id = $1
		ORDER BY created_at DESC`

	rows, err := r.db.QueryContext(ctx, query, companyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var employees []*models.Employee
	for rows.Next() {
		employee := &models.Employee{}
		if err := rows.Scan(
			&employee.UserID,
			&employee.CompanyID,
			&employee.StoreID,
			&employee.EmployeeIDNumber,
			&employee.JoinDate,
			&employee.Position,
			&employee.CreatedAt,
			&employee.UpdatedAt,
		); err != nil {
			return nil, err
		}
		employees = append(employees, employee)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return employees, nil
}
