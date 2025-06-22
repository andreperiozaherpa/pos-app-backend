package models

import "github.com/google/uuid"

// EmployeeRole merepresentasikan hubungan antara karyawan dan peran.
// Ini adalah model untuk tabel pivot 'employee_roles'.
type EmployeeRole struct {
	EmployeeUserID uuid.UUID `db:"employee_user_id"`
	RoleID         int32     `db:"role_id"` // Sesuai dengan tipe SERIAL di roles.id
}
