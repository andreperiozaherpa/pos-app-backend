package services

import (
	"context"
	"pos-app/backend/internal/models"
	"time"

	"github.com/google/uuid"
)

// ShiftService menangani logika bisnis manajemen shift dan absensi pegawai.
type ShiftService interface {
	// CRUD utama
	CreateShift(ctx context.Context, shift *models.Shift) (uuid.UUID, error)
	GetShiftByID(ctx context.Context, id uuid.UUID) (*models.Shift, error)
	ListShiftsByEmployeeID(ctx context.Context, employeeUserID uuid.UUID) ([]*models.Shift, error)
	ListShiftsByStoreAndDateRange(ctx context.Context, storeID uuid.UUID, from, to time.Time) ([]*models.Shift, error)
	UpdateShift(ctx context.Context, shift *models.Shift) error
	DeleteShift(ctx context.Context, id uuid.UUID) error

	// Opsional & Custom
	GetShiftAttendance(ctx context.Context, shiftID uuid.UUID) (*models.ShiftAttendance, error)
	ApproveShiftSwap(ctx context.Context, swapRequestID uuid.UUID) error
	ExportShifts(ctx context.Context, storeID uuid.UUID, from, to time.Time) ([]byte, error)
	RecordCheckIn(ctx context.Context, shiftID, userID uuid.UUID, checkInTime time.Time) error
	RecordCheckOut(ctx context.Context, shiftID, userID uuid.UUID, checkOutTime time.Time) error
	RequestShiftSwap(ctx context.Context, shiftID, fromUserID, toUserID uuid.UUID) error
	CancelShift(ctx context.Context, shiftID uuid.UUID) error
	ListShiftsByDateRange(ctx context.Context, storeID uuid.UUID, from, to time.Time) ([]*models.Shift, error)
	ExportShiftAttendance(ctx context.Context, storeID uuid.UUID, from, to time.Time) ([]byte, error)
	BulkUpdateShifts(ctx context.Context, shifts []*models.Shift) error
	ListShiftSwaps(ctx context.Context, storeID uuid.UUID, from, to time.Time) ([]*models.ShiftSwap, error)
}
