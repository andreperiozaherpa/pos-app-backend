package postgres

import (
	"context"
	"database/sql"
	"errors"

	"pos-app/backend/internal/core/repository"
	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

// permissionRepositoryPG implementasi repository Permission menggunakan PostgreSQL
// Sesuai ERD, tabel permissions menggunakan SERIAL id
// dan atribut name, description
// created_at dan updated_at tidak ada di ERD sehingga tidak dipakai
// jadi kita tidak pakai created_at dan updated_at

type permissionRepositoryPG struct {
	db *sql.DB
}

// NewPermissionRepositoryPG membuat instance baru permissionRepositoryPG
func NewPermissionRepositoryPG(db *sql.DB) repository.PermissionRepository {
	return &permissionRepositoryPG{db: db}
}

// Create menambahkan permission baru ke database.
func (r *permissionRepositoryPG) Create(ctx context.Context, p *models.Permission) error {
	query := `INSERT INTO permissions (name, description) VALUES ($1, $2)`
	_, err := r.db.ExecContext(ctx, query, p.Name, p.Description)
	return err
}

// GetByID mengambil permission berdasarkan ID.
func (r *permissionRepositoryPG) GetByID(ctx context.Context, id int32) (*models.Permission, error) {
	query := `SELECT id, name, description FROM permissions WHERE id = $1`
	row := r.db.QueryRowContext(ctx, query, id)

	p := &models.Permission{}
	err := row.Scan(&p.ID, &p.Name, &p.Description)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}
	return p, nil
}

// Update memperbarui permission berdasarkan ID.
func (r *permissionRepositoryPG) Update(ctx context.Context, p *models.Permission) error {
	query := `UPDATE permissions SET name = $1, description = $2 WHERE id = $3`
	result, err := r.db.ExecContext(ctx, query, p.Name, p.Description, p.ID)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return repository.ErrNotFound
	}
	return nil
}

// Delete menghapus permission berdasarkan ID.
func (r *permissionRepositoryPG) Delete(ctx context.Context, id int32) error {
	query := `DELETE FROM permissions WHERE id = $1`
	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return repository.ErrNotFound
	}
	return nil
}

// ListAll mengambil semua permission.
func (r *permissionRepositoryPG) ListAll(ctx context.Context) ([]*models.Permission, error) {
	query := `SELECT id, name, description FROM permissions ORDER BY name ASC`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var permissions []*models.Permission
	for rows.Next() {
		p := &models.Permission{}
		if err := rows.Scan(&p.ID, &p.Name, &p.Description); err != nil {
			return nil, err
		}
		permissions = append(permissions, p)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return permissions, nil
}

// ListPermissionsByRole mendapatkan daftar permission yang dimiliki role tertentu
func (r *permissionRepositoryPG) ListPermissionsByRole(ctx context.Context, roleID uuid.UUID) ([]*models.Permission, error) {
	query := `
		SELECT p.id, p.name, p.description
		FROM permissions p
		INNER JOIN role_permissions rp ON p.id = rp.permission_id
		WHERE rp.role_id = $1
	`
	rows, err := r.db.QueryContext(ctx, query, roleID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var permissions []*models.Permission
	for rows.Next() {
		p := &models.Permission{}
		if err := rows.Scan(&p.ID, &p.Name, &p.Description); err != nil {
			return nil, err
		}
		permissions = append(permissions, p)
	}
	return permissions, nil
}

// AssignPermissionToRole menambahkan permission ke role tertentu
func (r *permissionRepositoryPG) AssignPermissionToRole(ctx context.Context, roleID, permissionID uuid.UUID) error {
	query := `INSERT INTO role_permissions (role_id, permission_id) VALUES ($1, $2) ON CONFLICT DO NOTHING`
	_, err := r.db.ExecContext(ctx, query, roleID, permissionID)
	return err
}

// RemovePermissionFromRole menghapus permission dari role
func (r *permissionRepositoryPG) RemovePermissionFromRole(ctx context.Context, roleID, permissionID uuid.UUID) error {
	query := `DELETE FROM role_permissions WHERE role_id = $1 AND permission_id = $2`
	result, err := r.db.ExecContext(ctx, query, roleID, permissionID)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return repository.ErrNotFound
	}
	return nil
}
