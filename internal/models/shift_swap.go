package models

import (
	"time"

	"github.com/google/uuid"
)

type ShiftSwap struct {
	ID                        uuid.UUID  `db:"id" json:"id"`
	RequestedByEmployeeUserID uuid.UUID  `db:"requested_by_employee_user_id" json:"requested_by_employee_user_id"`
	RequestedToEmployeeUserID uuid.UUID  `db:"requested_to_employee_user_id" json:"requested_to_employee_user_id"`
	ShiftID                   uuid.UUID  `db:"shift_id" json:"shift_id"`
	RequestedShiftDate        time.Time  `db:"requested_shift_date" json:"requested_shift_date"`
	Reason                    string     `db:"reason" json:"reason"`
	Status                    string     `db:"status" json:"status"`
	ApprovedByUserID          *uuid.UUID `db:"approved_by_user_id" json:"approved_by_user_id"`
	ApprovedAt                *time.Time `db:"approved_at" json:"approved_at"`
	CreatedAt                 time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt                 time.Time  `db:"updated_at" json:"updated_at"`
}
