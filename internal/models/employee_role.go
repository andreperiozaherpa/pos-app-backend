package models

import "github.com/google/uuid"

// EmployeeRole merepresentasikan hubungan antara karyawan dan peran (role).
// Ini adalah model untuk tabel pivot 'employee_roles'.
type EmployeeRole struct {
	EmployeeUserID uuid.UUID `db:"employee_user_id" json:"employee_user_id"` // FK ke employees.user_id
	RoleID         int32     `db:"role_id" json:"role_id"`                   // FK ke roles.id (SERIAL)
}
