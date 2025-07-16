package postgres

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"pos-app/backend/internal/core/repository"
	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

type shiftRepositoryPG struct {
	db *sql.DB
}

func NewShiftRepositoryPG(db *sql.DB) repository.ShiftRepository {
	return &shiftRepositoryPG{db: db}
}

// Create membuat data shift baru.
func (r *shiftRepositoryPG) Create(ctx context.Context, shift *models.Shift) error {
	query := `
        INSERT INTO shifts (
            id, employee_user_id, store_id, shift_date, start_time, end_time,
            actual_check_in, actual_check_out, notes, created_by_user_id, created_at, updated_at
        )
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, NOW(), NOW())`
	_, err := r.db.ExecContext(ctx, query,
		shift.ID, shift.EmployeeUserID, shift.StoreID, shift.ShiftDate, shift.StartTime, shift.EndTime,
		shift.ActualCheckIn, shift.ActualCheckOut, shift.Notes, shift.CreatedByUserID,
	)
	return err
}

// GetByID mengambil data shift berdasarkan ID.
func (r *shiftRepositoryPG) GetByID(ctx context.Context, id uuid.UUID) (*models.Shift, error) {
	query := `
        SELECT id, employee_user_id, store_id, shift_date, start_time, end_time,
               actual_check_in, actual_check_out, notes, created_by_user_id, created_at, updated_at
        FROM shifts WHERE id=$1`
	row := r.db.QueryRowContext(ctx, query, id)
	shift := &models.Shift{}
	err := row.Scan(
		&shift.ID, &shift.EmployeeUserID, &shift.StoreID, &shift.ShiftDate,
		&shift.StartTime, &shift.EndTime,
		&shift.ActualCheckIn, &shift.ActualCheckOut,
		&shift.Notes, &shift.CreatedByUserID,
		&shift.CreatedAt, &shift.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}
	return shift, nil
}

// ListByEmployeeID mengambil daftar shift berdasarkan user ID karyawan.
func (r *shiftRepositoryPG) ListByEmployeeID(ctx context.Context, employeeUserID uuid.UUID) ([]*models.Shift, error) {
	query := `
        SELECT id, employee_user_id, store_id, shift_date, start_time, end_time,
               actual_check_in, actual_check_out, notes, created_by_user_id, created_at, updated_at
        FROM shifts
        WHERE employee_user_id=$1`
	rows, err := r.db.QueryContext(ctx, query, employeeUserID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var shifts []*models.Shift
	for rows.Next() {
		shift := &models.Shift{}
		if err := rows.Scan(
			&shift.ID, &shift.EmployeeUserID, &shift.StoreID, &shift.ShiftDate,
			&shift.StartTime, &shift.EndTime,
			&shift.ActualCheckIn, &shift.ActualCheckOut,
			&shift.Notes, &shift.CreatedByUserID,
			&shift.CreatedAt, &shift.UpdatedAt,
		); err != nil {
			return nil, err
		}
		shifts = append(shifts, shift)
	}
	return shifts, nil
}

// ListByStoreIDAndDateRange mengambil daftar shift berdasarkan store ID dan rentang tanggal.
func (r *shiftRepositoryPG) ListByStoreIDAndDateRange(ctx context.Context, storeID uuid.UUID, start, end time.Time) ([]*models.Shift, error) {
	query := `
        SELECT id, employee_user_id, store_id, shift_date, start_time, end_time,
               actual_check_in, actual_check_out, notes, created_by_user_id, created_at, updated_at
        FROM shifts
        WHERE store_id=$1 AND shift_date >= $2 AND shift_date <= $3
        ORDER BY shift_date, start_time`
	rows, err := r.db.QueryContext(ctx, query, storeID, start, end)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var shifts []*models.Shift
	for rows.Next() {
		shift := &models.Shift{}
		if err := rows.Scan(
			&shift.ID, &shift.EmployeeUserID, &shift.StoreID, &shift.ShiftDate,
			&shift.StartTime, &shift.EndTime,
			&shift.ActualCheckIn, &shift.ActualCheckOut,
			&shift.Notes, &shift.CreatedByUserID,
			&shift.CreatedAt, &shift.UpdatedAt,
		); err != nil {
			return nil, err
		}
		shifts = append(shifts, shift)
	}
	return shifts, nil
}

// Update memperbarui data shift.
func (r *shiftRepositoryPG) Update(ctx context.Context, shift *models.Shift) error {
	query := `
        UPDATE shifts
        SET employee_user_id=$1, store_id=$2, shift_date=$3, start_time=$4, end_time=$5,
            actual_check_in=$6, actual_check_out=$7, notes=$8, created_by_user_id=$9, updated_at=NOW()
        WHERE id=$10`
	res, err := r.db.ExecContext(ctx, query,
		shift.EmployeeUserID, shift.StoreID, shift.ShiftDate, shift.StartTime, shift.EndTime,
		shift.ActualCheckIn, shift.ActualCheckOut, shift.Notes, shift.CreatedByUserID, shift.ID,
	)
	if err != nil {
		return err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return repository.ErrNotFound
	}
	return nil
}

// Delete menghapus data shift berdasarkan ID.
func (r *shiftRepositoryPG) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM shifts WHERE id=$1`
	res, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return repository.ErrNotFound
	}
	return nil
}
