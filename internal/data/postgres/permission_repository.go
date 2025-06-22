package postgres

import (
	"context"
	"database/sql"
	"pos-app/backend/internal/models"
)

// PermissionRepository mendefinisikan interface untuk operasi data terkait Permission.
type PermissionRepository interface {
	Create(ctx context.Context, permission *models.Permission) error
	GetByID(ctx context.Context, id int32) (*models.Permission, error)
	ListAll(ctx context.Context) ([]*models.Permission, error)
	Update(ctx context.Context, permission *models.Permission) error
	Delete(ctx context.Context, id int32) error
}

// pgPermissionRepository adalah implementasi dari PermissionRepository untuk PostgreSQL.
type pgPermissionRepository struct {
	db DBExecutor
}

// NewPgPermissionRepository adalah constructor untuk membuat instance baru dari pgPermissionRepository.
func NewPgPermissionRepository(db DBExecutor) PermissionRepository {
	return &pgPermissionRepository{db: db}
}

// Create menyisipkan izin baru dan mengembalikan ID yang di-generate.
func (r *pgPermissionRepository) Create(ctx context.Context, p *models.Permission) error {
	query := `
		INSERT INTO permissions (name, description, group_name)
		VALUES ($1, $2, $3)
		RETURNING id`
	return r.db.QueryRowContext(ctx, query, p.Name, p.Description, p.GroupName).Scan(&p.ID)
}

// GetByID mengambil izin berdasarkan ID.
func (r *pgPermissionRepository) GetByID(ctx context.Context, id int32) (*models.Permission, error) {
	p := &models.Permission{}
	query := `SELECT id, name, description, group_name FROM permissions WHERE id = $1`
	err := r.db.QueryRowContext(ctx, query, id).Scan(&p.ID, &p.Name, &p.Description, &p.GroupName)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, err
	}
	return p, nil
}

// ListAll mengambil semua izin yang ada, diurutkan berdasarkan grup dan nama.
func (r *pgPermissionRepository) ListAll(ctx context.Context) ([]*models.Permission, error) {
	query := `SELECT id, name, description, group_name FROM permissions ORDER BY group_name, name`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var permissions []*models.Permission
	for rows.Next() {
		p := &models.Permission{}
		if err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.GroupName); err != nil {
			return nil, err
		}
		permissions = append(permissions, p)
	}
	return permissions, rows.Err()
}

// Update memperbarui data izin.
func (r *pgPermissionRepository) Update(ctx context.Context, p *models.Permission) error {
	query := `
		UPDATE permissions
		SET name = $1, description = $2, group_name = $3
		WHERE id = $4`
	_, err := r.db.ExecContext(ctx, query, p.Name, p.Description, p.GroupName, p.ID)
	return err
}

// Delete menghapus izin dari database.
func (r *pgPermissionRepository) Delete(ctx context.Context, id int32) error {
	query := `DELETE FROM permissions WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
