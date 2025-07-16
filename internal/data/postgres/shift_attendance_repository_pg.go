package postgres

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"pos-app/backend/internal/core/repository"
	"pos-app/backend/internal/models"
)

type shiftAttendanceRepositoryPG struct {
	db *sql.DB
}

func NewShiftAttendanceRepositoryPG(db *sql.DB) repository.ShiftAttendanceRepository {
	return &shiftAttendanceRepositoryPG{db: db}
}

func (r *shiftAttendanceRepositoryPG) Create(ctx context.Context, attendance *models.ShiftAttendance) error {
	query := `
        INSERT INTO shift_attendances (id, shift_id, employee_id, check_in, check_out, created_at)
        VALUES ($1, $2, $3, $4, $5, NOW())`
	_, err := r.db.ExecContext(ctx, query,
		attendance.ID, attendance.ShiftID, attendance.EmployeeID, attendance.CheckIn, attendance.CheckOut)
	return err
}

func (r *shiftAttendanceRepositoryPG) GetByID(ctx context.Context, id string) (*models.ShiftAttendance, error) {
	query := `
        SELECT id, shift_id, employee_id, check_in, check_out, created_at
        FROM shift_attendances WHERE id=$1`
	row := r.db.QueryRowContext(ctx, query, id)
	attendance := &models.ShiftAttendance{}
	err := row.Scan(&attendance.ID, &attendance.ShiftID, &attendance.EmployeeID, &attendance.CheckIn, &attendance.CheckOut, &attendance.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}
	return attendance, nil
}

// Ubah nama method agar sesuai interface
func (r *shiftAttendanceRepositoryPG) ListByShiftID(ctx context.Context, shiftID string) ([]*models.ShiftAttendance, error) {
	query := `
        SELECT id, shift_id, employee_id, check_in, check_out, created_at
        FROM shift_attendances
        WHERE shift_id=$1 ORDER BY check_in`
	rows, err := r.db.QueryContext(ctx, query, shiftID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var attendances []*models.ShiftAttendance
	for rows.Next() {
		a := &models.ShiftAttendance{}
		if err := rows.Scan(&a.ID, &a.ShiftID, &a.EmployeeID, &a.CheckIn, &a.CheckOut, &a.CreatedAt); err != nil {
			return nil, err
		}
		attendances = append(attendances, a)
	}
	return attendances, nil
}

// Tambahkan ListByDateRange sesuai interface
func (r *shiftAttendanceRepositoryPG) ListByDateRange(ctx context.Context, startDate, endDate time.Time) ([]*models.ShiftAttendance, error) {
	query := `
        SELECT id, shift_id, employee_id, check_in, check_out, created_at
        FROM shift_attendances
        WHERE check_in >= $1 AND check_in <= $2
        ORDER BY check_in`
	rows, err := r.db.QueryContext(ctx, query, startDate, endDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var attendances []*models.ShiftAttendance
	for rows.Next() {
		a := &models.ShiftAttendance{}
		if err := rows.Scan(&a.ID, &a.ShiftID, &a.EmployeeID, &a.CheckIn, &a.CheckOut, &a.CreatedAt); err != nil {
			return nil, err
		}
		attendances = append(attendances, a)
	}
	return attendances, nil
}
