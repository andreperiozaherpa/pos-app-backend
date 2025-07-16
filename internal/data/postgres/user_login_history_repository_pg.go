package postgres

import (
	"context"
	"database/sql"
	"errors"

	"pos-app/backend/internal/core/repository"
	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

type userLoginHistoryRepositoryPG struct {
	db *sql.DB
}

func NewUserLoginHistoryRepositoryPG(db *sql.DB) repository.UserLoginHistoryRepository {
	return &userLoginHistoryRepositoryPG{db: db}
}

func (r *userLoginHistoryRepositoryPG) Create(ctx context.Context, history *models.UserLoginHistory) error {
	query := `
        INSERT INTO user_login_histories (id, user_id, login_time, ip_address, device_info)
        VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.ExecContext(ctx, query,
		history.ID, history.UserID, history.LoginTime, history.IPAddress, history.DeviceInfo)
	return err
}

func (r *userLoginHistoryRepositoryPG) GetByID(ctx context.Context, id uuid.UUID) (*models.UserLoginHistory, error) {
	query := `
        SELECT id, user_id, login_time, ip_address, device_info
        FROM user_login_histories WHERE id=$1`
	row := r.db.QueryRowContext(ctx, query, id)
	history := &models.UserLoginHistory{}
	err := row.Scan(&history.ID, &history.UserID, &history.LoginTime, &history.IPAddress, &history.DeviceInfo)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}
	return history, nil
}

func (r *userLoginHistoryRepositoryPG) ListByUserID(ctx context.Context, userID uuid.UUID) ([]*models.UserLoginHistory, error) {
	query := `
        SELECT id, user_id, login_time, ip_address, device_info
        FROM user_login_histories WHERE user_id=$1 ORDER BY login_time DESC`
	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var histories []*models.UserLoginHistory
	for rows.Next() {
		h := &models.UserLoginHistory{}
		if err := rows.Scan(&h.ID, &h.UserID, &h.LoginTime, &h.IPAddress, &h.DeviceInfo); err != nil {
			return nil, err
		}
		histories = append(histories, h)
	}
	return histories, nil
}
