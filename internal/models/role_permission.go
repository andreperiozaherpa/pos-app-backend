package models

// RolePermission merepresentasikan hubungan antara peran dan izin dari tabel 'role_permissions'.
// Ini adalah tabel pivot antara roles dan permissions.
type RolePermission struct {
	// ID peran, merujuk ke roles.id (SERIAL di DB, int32 di Go)
	RoleID int32 `db:"role_id"`
	// ID izin, merujuk ke permissions.id (SERIAL di DB, int32 di Go)
	PermissionID int32 `db:"permission_id"`
}
