package postgres

import (
	"context"
	"database/sql"
	"pos-app/backend/internal/models"
)

// RoleRepository mendefinisikan interface untuk operasi data terkait Role.
type RoleRepository interface {
	Create(ctx context.Context, role *models.Role) error
	GetByID(ctx context.Context, id int32) (*models.Role, error)
	GetByName(ctx context.Context, name string) (*models.Role, error)
	ListAll(ctx context.Context) ([]*models.Role, error)
	Update(ctx context.Context, role *models.Role) error
	Delete(ctx context.Context, id int32) error
}

// pgRoleRepository adalah implementasi dari RoleRepository untuk PostgreSQL.
type pgRoleRepository struct {
	db DBExecutor
}

// NewPgRoleRepository adalah constructor untuk membuat instance baru dari pgRoleRepository.
func NewPgRoleRepository(db DBExecutor) RoleRepository {
	return &pgRoleRepository{db: db}
}

// Implementasi metode-metode dari interface RoleRepository:

func (r *pgRoleRepository) Create(ctx context.Context, role *models.Role) error {
	// ID akan di-generate oleh database (SERIAL), jadi kita gunakan RETURNING id
	query := `
		INSERT INTO roles (name, description)
		VALUES ($1, $2)
		RETURNING id`
	return r.db.QueryRowContext(ctx, query, role.Name, role.Description).Scan(&role.ID)
}

func (r *pgRoleRepository) GetByID(ctx context.Context, id int32) (*models.Role, error) {
	role := &models.Role{}
	query := `
		SELECT id, name, description
		FROM roles
		WHERE id = $1`
	err := r.db.QueryRowContext(ctx, query, id).Scan(&role.ID, &role.Name, &role.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, err
	}
	return role, nil
}

func (r *pgRoleRepository) GetByName(ctx context.Context, name string) (*models.Role, error) {
	role := &models.Role{}
	query := `
		SELECT id, name, description
		FROM roles
		WHERE name = $1`
	err := r.db.QueryRowContext(ctx, query, name).Scan(&role.ID, &role.Name, &role.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, err
	}
	return role, nil
}

func (r *pgRoleRepository) ListAll(ctx context.Context) ([]*models.Role, error) {
	query := `SELECT id, name, description FROM roles ORDER BY name ASC`
	rows, err := r.db.QueryContext(ctx, query)
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

func (r *pgRoleRepository) Update(ctx context.Context, role *models.Role) error {
	query := `UPDATE roles SET name = $1, description = $2 WHERE id = $3`
	_, err := r.db.ExecContext(ctx, query, role.Name, role.Description, role.ID)
	return err
}

func (r *pgRoleRepository) Delete(ctx context.Context, id int32) error {
	query := `DELETE FROM roles WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
