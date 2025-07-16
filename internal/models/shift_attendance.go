package models

import (
	"time"

	"github.com/google/uuid"
)

type ShiftAttendance struct {
	ID         uuid.UUID  `db:"id" json:"id"`
	ShiftID    uuid.UUID  `db:"shift_id" json:"shift_id"`
	EmployeeID uuid.UUID  `db:"employee_id" json:"employee_id"`
	CheckIn    *time.Time `db:"check_in" json:"check_in"`
	CheckOut   *time.Time `db:"check_out" json:"check_out"`
	CreatedAt  time.Time  `db:"created_at" json:"created_at"`
}
