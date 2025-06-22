package postgres

import (
	"context"
	"database/sql"
	"pos-app/backend/internal/models"
	"time"

	"github.com/google/uuid"
)

// ShiftRepository mendefinisikan interface untuk operasi data terkait Shift.
type ShiftRepository interface {
	Create(ctx context.Context, shift *models.Shift) error
	GetByID(ctx context.Context, id uuid.UUID) (*models.Shift, error)
	ListByEmployeeID(ctx context.Context, employeeUserID uuid.UUID) ([]*models.Shift, error)
	ListByStoreIDAndDateRange(ctx context.Context, storeID uuid.UUID, startDate, endDate time.Time) ([]*models.Shift, error)
	Update(ctx context.Context, shift *models.Shift) error
	Delete(ctx context.Context, id uuid.UUID) error
}

// pgShiftRepository adalah implementasi dari ShiftRepository untuk PostgreSQL.
type pgShiftRepository struct {
	db DBExecutor
}

// NewPgShiftRepository adalah constructor untuk membuat instance baru dari pgShiftRepository.
func NewPgShiftRepository(db DBExecutor) ShiftRepository {
	return &pgShiftRepository{db: db}
}

// Implementasi metode-metode dari interface ShiftRepository:

func (r *pgShiftRepository) Create(ctx context.Context, s *models.Shift) error {
	query := `
		INSERT INTO shifts (id, employee_user_id, store_id, shift_date, start_time, end_time,
			actual_check_in, actual_check_out, notes, created_by_user_id, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`
	_, err := r.db.ExecContext(ctx, query,
		s.ID, s.EmployeeUserID, s.StoreID, s.ShiftDate, s.StartTime, s.EndTime,
		s.ActualCheckIn, s.ActualCheckOut, s.Notes, s.CreatedByUserID, s.CreatedAt, s.UpdatedAt,
	)
	return err
}

func (r *pgShiftRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.Shift, error) {
	s := &models.Shift{}
	query := `
		SELECT id, employee_user_id, store_id, shift_date, start_time, end_time,
			actual_check_in, actual_check_out, notes, created_by_user_id, created_at, updated_at
		FROM shifts
		WHERE id = $1`
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&s.ID, &s.EmployeeUserID, &s.StoreID, &s.ShiftDate, &s.StartTime, &s.EndTime,
		&s.ActualCheckIn, &s.ActualCheckOut, &s.Notes, &s.CreatedByUserID, &s.CreatedAt, &s.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, err
	}
	return s, nil
}

func (r *pgShiftRepository) ListByEmployeeID(ctx context.Context, employeeUserID uuid.UUID) ([]*models.Shift, error) {
	query := `
		SELECT id, employee_user_id, store_id, shift_date, start_time, end_time,
			actual_check_in, actual_check_out, notes, created_by_user_id, created_at, updated_at
		FROM shifts
		WHERE employee_user_id = $1
		ORDER BY shift_date DESC, start_time ASC`
	rows, err := r.db.QueryContext(ctx, query, employeeUserID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var shifts []*models.Shift
	for rows.Next() {
		s := &models.Shift{}
		if err := rows.Scan(
			&s.ID, &s.EmployeeUserID, &s.StoreID, &s.ShiftDate, &s.StartTime, &s.EndTime,
			&s.ActualCheckIn, &s.ActualCheckOut, &s.Notes, &s.CreatedByUserID, &s.CreatedAt, &s.UpdatedAt,
		); err != nil {
			return nil, err
		}
		shifts = append(shifts, s)
	}
	return shifts, rows.Err()
}

func (r *pgShiftRepository) ListByStoreIDAndDateRange(ctx context.Context, storeID uuid.UUID, startDate, endDate time.Time) ([]*models.Shift, error) {
	query := `
		SELECT id, employee_user_id, store_id, shift_date, start_time, end_time,
			actual_check_in, actual_check_out, notes, created_by_user_id, created_at, updated_at
		FROM shifts
		WHERE store_id = $1 AND shift_date BETWEEN $2 AND $3
		ORDER BY shift_date ASC, start_time ASC`
	rows, err := r.db.QueryContext(ctx, query, storeID, startDate, endDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var shifts []*models.Shift
	for rows.Next() {
		s := &models.Shift{}
		if err := rows.Scan(
			&s.ID, &s.EmployeeUserID, &s.StoreID, &s.ShiftDate, &s.StartTime, &s.EndTime,
			&s.ActualCheckIn, &s.ActualCheckOut, &s.Notes, &s.CreatedByUserID, &s.CreatedAt, &s.UpdatedAt,
		); err != nil {
			return nil, err
		}
		shifts = append(shifts, s)
	}
	return shifts, rows.Err()
}

func (r *pgShiftRepository) Update(ctx context.Context, s *models.Shift) error {
	query := `
		UPDATE shifts
		SET employee_user_id = $1, store_id = $2, shift_date = $3, start_time = $4, end_time = $5,
			actual_check_in = $6, actual_check_out = $7, notes = $8, created_by_user_id = $9, updated_at = $10
		WHERE id = $11`
	_, err := r.db.ExecContext(ctx, query,
		s.EmployeeUserID, s.StoreID, s.ShiftDate, s.StartTime, s.EndTime,
		s.ActualCheckIn, s.ActualCheckOut, s.Notes, s.CreatedByUserID, s.UpdatedAt, s.ID,
	)
	return err
}

func (r *pgShiftRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM shifts WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
