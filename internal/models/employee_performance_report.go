package models

import (
	"time"

	"github.com/google/uuid"
)

type EmployeePerformanceReport struct {
	ID               uuid.UUID `db:"id" json:"id"`
	EmployeeUserID   uuid.UUID `db:"employee_user_id" json:"employee_user_id"`
	ReportDate       time.Time `db:"report_date" json:"report_date"`
	PerformanceScore float64   `db:"performance_score" json:"performance_score"`
	CreatedAt        time.Time `db:"created_at" json:"created_at"`
}
