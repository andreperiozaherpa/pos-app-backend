package repository

import (
	"context"
	"time"

	"pos-app/backend/internal/models"
)

// ShiftAttendanceRepository adalah interface untuk operasi CRUD
// terkait data absensi shift.
type ShiftAttendanceRepository interface {
	Create(ctx context.Context, attendance *models.ShiftAttendance) error
	GetByID(ctx context.Context, id string) (*models.ShiftAttendance, error)
	ListByShiftID(ctx context.Context, shiftID string) ([]*models.ShiftAttendance, error)
	ListByDateRange(ctx context.Context, startDate, endDate time.Time) ([]*models.ShiftAttendance, error)
}
