package postgres

import (
	"context"
	"database/sql"

	"pos-app/backend/internal/core/repository"
	"pos-app/backend/internal/models"
)

// roleRepositoryPG adalah implementasi repository Role menggunakan PostgreSQL.
type roleRepositoryPG struct {
	db *sql.DB
}

// NewRoleRepositoryPG membuat instance baru RoleRepository PostgreSQL.
func NewRoleRepositoryPG(db *sql.DB) repository.RoleRepository {
	return &roleRepositoryPG{db: db}
}

// Create menambahkan role baru ke database.
// Karena id adalah SERIAL, kita biarkan DB yang generate dan mengambil kembali id dengan RETURNING.
func (r *roleRepositoryPG) Create(ctx context.Context, role *models.Role) error {
	query := `
		INSERT INTO roles (name, description)
		VALUES ($1, $2)
		RETURNING id`
	err := r.db.QueryRowContext(ctx, query, role.Name, role.Description).Scan(&role.ID)
	return err
}

// GetByID mengambil role berdasarkan ID.
func (r *roleRepositoryPG) GetByID(ctx context.Context, id int32) (*models.Role, error) {
	query := `
		SELECT id, name, description
		FROM roles WHERE id = $1`
	row := r.db.QueryRowContext(ctx, query, id)
	var role models.Role
	err := row.Scan(
		&role.ID, &role.Name, &role.Description,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}
	return &role, nil
}

// Update memperbarui data role yang sudah ada.
func (r *roleRepositoryPG) Update(ctx context.Context, role *models.Role) error {
	query := `
		UPDATE roles
		SET name = $1, description = $2
		WHERE id = $3`
	_, err := r.db.ExecContext(ctx, query,
		role.Name, role.Description, role.ID,
	)
	return err
}

// Delete menghapus role berdasarkan ID.
func (r *roleRepositoryPG) Delete(ctx context.Context, id int32) error {
	query := `DELETE FROM roles WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

// ListRolesByUser mengambil daftar role yang dimiliki oleh user tertentu.
func (r *roleRepositoryPG) ListRolesByUser(ctx context.Context, userID string) ([]*models.Role, error) {
	query := `
		SELECT r.id, r.name, r.description
		FROM roles r
		INNER JOIN employee_role er ON er.role_id = r.id
		WHERE er.employee_id = $1`

	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var roles []*models.Role
	for rows.Next() {
		var role models.Role
		err := rows.Scan(
			&role.ID, &role.Name, &role.Description,
		)
		if err != nil {
			return nil, err
		}
		roles = append(roles, &role)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return roles, nil
}

// AssignRoleToUser menetapkan role ke user (employee).
func (r *roleRepositoryPG) AssignRoleToUser(ctx context.Context, userID string, roleID int32) error {
	query := `
		INSERT INTO employee_role (employee_id, role_id)
		VALUES ($1, $2)
		ON CONFLICT DO NOTHING`
	_, err := r.db.ExecContext(ctx, query, userID, roleID)
	return err
}

// RemoveRoleFromUser menghapus penugasan role dari user (employee).
func (r *roleRepositoryPG) RemoveRoleFromUser(ctx context.Context, userID string, roleID int32) error {
	query := `
		DELETE FROM employee_role
		WHERE employee_id = $1 AND role_id = $2`
	_, err := r.db.ExecContext(ctx, query, userID, roleID)
	return err
}

// GetByName mengambil role berdasarkan nama role.
func (r *roleRepositoryPG) GetByName(ctx context.Context, name string) (*models.Role, error) {
	query := `
		SELECT id, name, description
		FROM roles WHERE name = $1`
	row := r.db.QueryRowContext(ctx, query, name)
	var role models.Role
	err := row.Scan(
		&role.ID, &role.Name, &role.Description,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}
	return &role, nil
}

// ListAll mengambil semua role yang ada di database.
func (r *roleRepositoryPG) ListAll(ctx context.Context) ([]*models.Role, error) {
	query := `
		SELECT id, name, description
		FROM roles
		ORDER BY name ASC`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var roles []*models.Role
	for rows.Next() {
		var role models.Role
		err := rows.Scan(
			&role.ID, &role.Name, &role.Description,
		)
		if err != nil {
			return nil, err
		}
		roles = append(roles, &role)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return roles, nil
}
