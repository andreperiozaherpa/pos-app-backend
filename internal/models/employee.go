package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

// Employee merepresentasikan data karyawan dari tabel 'employees'.
// Satu karyawan = satu user (UserID adalah PK dan FK ke users.id).
type Employee struct {
	UserID           uuid.UUID      `db:"user_id" json:"user_id"`                                 // Primary Key, Foreign Key ke users.id
	CompanyID        uuid.UUID      `db:"company_id" json:"company_id"`                           // ID perusahaan tempat karyawan bekerja
	StoreID          uuid.NullUUID  `db:"store_id" json:"store_id,omitempty"`                     // Cabang/toko, boleh null untuk peran sentral
	EmployeeIDNumber sql.NullString `db:"employee_id_number" json:"employee_id_number,omitempty"` // NIK atau nomor pegawai, opsional
	JoinDate         sql.NullTime   `db:"join_date" json:"join_date,omitempty"`                   // Tanggal gabung (bisa null)
	Position         sql.NullString `db:"position" json:"position,omitempty"`                     // Jabatan/posisi, opsional
	CreatedAt        time.Time      `db:"created_at" json:"created_at"`                           // Timestamp pembuatan
	UpdatedAt        time.Time      `db:"updated_at" json:"updated_at"`                           // Timestamp update terakhir
}
