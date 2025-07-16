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

type shiftSwapRepositoryPG struct {
	db *sql.DB
}

func NewShiftSwapRepositoryPG(db *sql.DB) repository.ShiftSwapRepository {
	return &shiftSwapRepositoryPG{db: db}
}

// Create menambahkan permintaan tukar shift ke database.
func (r *shiftSwapRepositoryPG) Create(ctx context.Context, swap *models.ShiftSwap) error {
	query := `
		INSERT INTO shift_swaps (
			id, requested_by_employee_user_id, requested_to_employee_user_id,
			shift_id, requested_shift_date, reason, status,
			approved_by_user_id, approved_at, created_at, updated_at
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, NOW(), NOW()
		)`
	_, err := r.db.ExecContext(ctx, query,
		swap.ID,
		swap.RequestedByEmployeeUserID,
		swap.RequestedToEmployeeUserID,
		swap.ShiftID,
		swap.RequestedShiftDate,
		swap.Reason,
		swap.Status,
		swap.ApprovedByUserID,
		swap.ApprovedAt,
	)
	return err
}

// GetByID mengambil data shift swap berdasarkan ID.
func (r *shiftSwapRepositoryPG) GetByID(ctx context.Context, id uuid.UUID) (*models.ShiftSwap, error) {
	query := `
		SELECT
			id, requested_by_employee_user_id, requested_to_employee_user_id,
			shift_id, requested_shift_date, reason, status,
			approved_by_user_id, approved_at, created_at, updated_at
		FROM shift_swaps WHERE id = $1`
	row := r.db.QueryRowContext(ctx, query, id)
	swap := &models.ShiftSwap{}
	var approvedByUserID sql.NullString
	var approvedAt sql.NullTime

	err := row.Scan(
		&swap.ID,
		&swap.RequestedByEmployeeUserID,
		&swap.RequestedToEmployeeUserID,
		&swap.ShiftID,
		&swap.RequestedShiftDate,
		&swap.Reason,
		&swap.Status,
		&approvedByUserID,
		&approvedAt,
		&swap.CreatedAt,
		&swap.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}

	if approvedByUserID.Valid {
		id, _ := uuid.Parse(approvedByUserID.String)
		swap.ApprovedByUserID = &id
	}
	if approvedAt.Valid {
		swap.ApprovedAt = &approvedAt.Time
	}

	return swap, nil
}

// ListByEmployee mengambil semua shift swap berdasarkan karyawan.
func (r *shiftSwapRepositoryPG) ListByEmployee(ctx context.Context, employeeUserID uuid.UUID) ([]*models.ShiftSwap, error) {
	query := `
		SELECT
			id, requested_by_employee_user_id, requested_to_employee_user_id,
			shift_id, requested_shift_date, reason, status,
			approved_by_user_id, approved_at, created_at, updated_at
		FROM shift_swaps
		WHERE requested_by_employee_user_id = $1 OR requested_to_employee_user_id = $1
		ORDER BY created_at DESC`
	rows, err := r.db.QueryContext(ctx, query, employeeUserID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var swaps []*models.ShiftSwap
	for rows.Next() {
		s := &models.ShiftSwap{}
		var approvedByUserID sql.NullString
		var approvedAt sql.NullTime

		if err := rows.Scan(
			&s.ID,
			&s.RequestedByEmployeeUserID,
			&s.RequestedToEmployeeUserID,
			&s.ShiftID,
			&s.RequestedShiftDate,
			&s.Reason,
			&s.Status,
			&approvedByUserID,
			&approvedAt,
			&s.CreatedAt,
			&s.UpdatedAt,
		); err != nil {
			return nil, err
		}
		if approvedByUserID.Valid {
			id, _ := uuid.Parse(approvedByUserID.String)
			s.ApprovedByUserID = &id
		}
		if approvedAt.Valid {
			s.ApprovedAt = &approvedAt.Time
		}
		swaps = append(swaps, s)
	}
	return swaps, nil
}

// ListByStatus mengambil semua shift swap berdasarkan status tertentu.
func (r *shiftSwapRepositoryPG) ListByStatus(ctx context.Context, status string) ([]*models.ShiftSwap, error) {
	query := `
		SELECT
			id, requested_by_employee_user_id, requested_to_employee_user_id,
			shift_id, requested_shift_date, reason, status,
			approved_by_user_id, approved_at, created_at, updated_at
		FROM shift_swaps
		WHERE status = $1
		ORDER BY created_at DESC`
	rows, err := r.db.QueryContext(ctx, query, status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var swaps []*models.ShiftSwap
	for rows.Next() {
		s := &models.ShiftSwap{}
		var approvedByUserID sql.NullString
		var approvedAt sql.NullTime

		if err := rows.Scan(
			&s.ID,
			&s.RequestedByEmployeeUserID,
			&s.RequestedToEmployeeUserID,
			&s.ShiftID,
			&s.RequestedShiftDate,
			&s.Reason,
			&s.Status,
			&approvedByUserID,
			&approvedAt,
			&s.CreatedAt,
			&s.UpdatedAt,
		); err != nil {
			return nil, err
		}
		if approvedByUserID.Valid {
			id, _ := uuid.Parse(approvedByUserID.String)
			s.ApprovedByUserID = &id
		}
		if approvedAt.Valid {
			s.ApprovedAt = &approvedAt.Time
		}
		swaps = append(swaps, s)
	}
	return swaps, nil
}

// UpdateStatus memperbarui status permintaan tukar shift (approve/reject/cancel).
func (r *shiftSwapRepositoryPG) UpdateStatus(ctx context.Context, id uuid.UUID, status string, approvedBy *uuid.UUID, approvedAt *time.Time) error {
	query := `
		UPDATE shift_swaps
		SET status = $1,
			approved_by_user_id = $2,
			approved_at = $3,
			updated_at = NOW()
		WHERE id = $4`
	_, err := r.db.ExecContext(ctx, query,
		status,
		approvedBy,
		approvedAt,
		id,
	)
	return err
}
