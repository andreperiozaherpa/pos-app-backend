package postgres

import (
	"context"
	"pos-app/backend/internal/models"
)

// RolePermissionRepository mendefinisikan interface untuk operasi data terkait RolePermission.
type RolePermissionRepository interface {
	Create(ctx context.Context, rp *models.RolePermission) error
	ListByRoleID(ctx context.Context, roleID int32) ([]*models.RolePermission, error)
	ListByPermissionID(ctx context.Context, permissionID int32) ([]*models.RolePermission, error)
	Delete(ctx context.Context, roleID int32, permissionID int32) error
	DeleteByRoleID(ctx context.Context, roleID int32) error
}

// pgRolePermissionRepository adalah implementasi dari RolePermissionRepository untuk PostgreSQL.
type pgRolePermissionRepository struct {
	db DBExecutor
}

// NewPgRolePermissionRepository adalah constructor untuk membuat instance baru dari pgRolePermissionRepository.
func NewPgRolePermissionRepository(db DBExecutor) RolePermissionRepository {
	return &pgRolePermissionRepository{db: db}
}

// Create menyisipkan hubungan peran-izin baru.
func (r *pgRolePermissionRepository) Create(ctx context.Context, rp *models.RolePermission) error {
	query := `
		INSERT INTO role_permissions (role_id, permission_id)
		VALUES ($1, $2)`
	_, err := r.db.ExecContext(ctx, query, rp.RoleID, rp.PermissionID)
	return err
}

// ListByRoleID mengambil semua izin yang terkait dengan peran tertentu.
func (r *pgRolePermissionRepository) ListByRoleID(ctx context.Context, roleID int32) ([]*models.RolePermission, error) {
	query := `
		SELECT role_id, permission_id
		FROM role_permissions
		WHERE role_id = $1`
	rows, err := r.db.QueryContext(ctx, query, roleID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rolePermissions []*models.RolePermission
	for rows.Next() {
		rp := &models.RolePermission{}
		if err := rows.Scan(&rp.RoleID, &rp.PermissionID); err != nil {
			return nil, err
		}
		rolePermissions = append(rolePermissions, rp)
	}
	return rolePermissions, rows.Err()
}

// ListByPermissionID mengambil semua peran yang memiliki izin tertentu.
func (r *pgRolePermissionRepository) ListByPermissionID(ctx context.Context, permissionID int32) ([]*models.RolePermission, error) {
	query := `
		SELECT role_id, permission_id
		FROM role_permissions
		WHERE permission_id = $1`
	rows, err := r.db.QueryContext(ctx, query, permissionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rolePermissions []*models.RolePermission
	for rows.Next() {
		rp := &models.RolePermission{}
		if err := rows.Scan(&rp.RoleID, &rp.PermissionID); err != nil {
			return nil, err
		}
		rolePermissions = append(rolePermissions, rp)
	}
	return rolePermissions, rows.Err()
}

// Delete menghapus hubungan peran-izin tertentu.
func (r *pgRolePermissionRepository) Delete(ctx context.Context, roleID int32, permissionID int32) error {
	query := `DELETE FROM role_permissions WHERE role_id = $1 AND permission_id = $2`
	_, err := r.db.ExecContext(ctx, query, roleID, permissionID)
	return err
}

// DeleteByRoleID menghapus semua izin yang terkait dengan peran tertentu.
func (r *pgRolePermissionRepository) DeleteByRoleID(ctx context.Context, roleID int32) error {
	query := `DELETE FROM role_permissions WHERE role_id = $1`
	_, err := r.db.ExecContext(ctx, query, roleID)
	return err
}
