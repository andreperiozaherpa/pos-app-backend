package postgres

import (
	"context"
	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

// EmployeeRoleRepository mendefinisikan interface untuk operasi data terkait EmployeeRole.
type EmployeeRoleRepository interface {
	AssignRoleToEmployee(ctx context.Context, employeeUserID uuid.UUID, roleID int32) error
	RemoveRoleFromEmployee(ctx context.Context, employeeUserID uuid.UUID, roleID int32) error
	GetRolesForEmployee(ctx context.Context, employeeUserID uuid.UUID) ([]*models.Role, error)
}

// pgEmployeeRoleRepository adalah implementasi dari EmployeeRoleRepository untuk PostgreSQL.
type pgEmployeeRoleRepository struct {
	db DBExecutor
}

// NewPgEmployeeRoleRepository adalah constructor untuk membuat instance baru dari pgEmployeeRoleRepository.
func NewPgEmployeeRoleRepository(db DBExecutor) EmployeeRoleRepository {
	return &pgEmployeeRoleRepository{db: db}
}

// Implementasi metode-metode dari interface EmployeeRoleRepository:

func (r *pgEmployeeRoleRepository) AssignRoleToEmployee(ctx context.Context, employeeUserID uuid.UUID, roleID int32) error {
	query := `
		INSERT INTO employee_roles (employee_user_id, role_id)
		VALUES ($1, $2)
		ON CONFLICT (employee_user_id, role_id) DO NOTHING` // Mencegah error jika role sudah di-assign
	_, err := r.db.ExecContext(ctx, query, employeeUserID, roleID)
	return err
}

func (r *pgEmployeeRoleRepository) RemoveRoleFromEmployee(ctx context.Context, employeeUserID uuid.UUID, roleID int32) error {
	query := `DELETE FROM employee_roles WHERE employee_user_id = $1 AND role_id = $2`
	_, err := r.db.ExecContext(ctx, query, employeeUserID, roleID)
	return err
}

func (r *pgEmployeeRoleRepository) GetRolesForEmployee(ctx context.Context, employeeUserID uuid.UUID) ([]*models.Role, error) {
	query := `
		SELECT r.id, r.name, r.description
		FROM roles r
		JOIN employee_roles er ON r.id = er.role_id
		WHERE er.employee_user_id = $1
		ORDER BY r.name ASC`

	rows, err := r.db.QueryContext(ctx, query, employeeUserID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var roles []*models.Role
	for rows.Next() {
		role := &models.Role{}
		if err := rows.Scan(&role.ID, &role.Name, &role.Description); err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}

	return roles, rows.Err()
}
