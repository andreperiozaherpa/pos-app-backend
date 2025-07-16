package postgres

import (
	"context"
	"database/sql"

	"pos-app/backend/internal/core/repository"
	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

type activityLogRepositoryPG struct {
	db *sql.DB
}

func NewActivityLogRepositoryPG(db *sql.DB) repository.ActivityLogRepository {
	return &activityLogRepositoryPG{db: db}
}

// Create membuat data activity log baru.
func (r *activityLogRepositoryPG) Create(ctx context.Context, log *models.ActivityLog) error {
	query := `
        INSERT INTO activity_logs (
			id, user_id, company_id, store_id, action_type, description, target_entity, target_entity_id, ip_address, user_agent, log_time
		)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, NOW())`
	_, err := r.db.ExecContext(ctx, query,
		log.ID, log.UserID, log.CompanyID, log.StoreID, log.ActionType, log.Description,
		log.TargetEntity, log.TargetEntityID, log.IPAddress, log.UserAgent,
	)
	return err
}

// ListByUserID mengambil daftar activity log berdasarkan user ID.
func (r *activityLogRepositoryPG) ListByUserID(ctx context.Context, userID uuid.UUID) ([]*models.ActivityLog, error) {
	query := `
        SELECT id, user_id, company_id, store_id, action_type, description, target_entity, target_entity_id, ip_address, user_agent, log_time
        FROM activity_logs
        WHERE user_id = $1
        ORDER BY log_time DESC`
	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var logs []*models.ActivityLog
	for rows.Next() {
		log := &models.ActivityLog{}
		if err := rows.Scan(
			&log.ID, &log.UserID, &log.CompanyID, &log.StoreID, &log.ActionType, &log.Description,
			&log.TargetEntity, &log.TargetEntityID, &log.IPAddress, &log.UserAgent, &log.LogTime,
		); err != nil {
			return nil, err
		}
		logs = append(logs, log)
	}
	return logs, nil
}

// ListByCompanyID mengambil daftar activity log berdasarkan company ID.
func (r *activityLogRepositoryPG) ListByCompanyID(ctx context.Context, companyID uuid.UUID) ([]*models.ActivityLog, error) {
	query := `
        SELECT id, user_id, company_id, store_id, action_type, description, target_entity, target_entity_id, ip_address, user_agent, log_time
        FROM activity_logs
        WHERE company_id = $1
        ORDER BY log_time DESC`
	rows, err := r.db.QueryContext(ctx, query, companyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var logs []*models.ActivityLog
	for rows.Next() {
		log := &models.ActivityLog{}
		if err := rows.Scan(
			&log.ID, &log.UserID, &log.CompanyID, &log.StoreID, &log.ActionType, &log.Description,
			&log.TargetEntity, &log.TargetEntityID, &log.IPAddress, &log.UserAgent, &log.LogTime,
		); err != nil {
			return nil, err
		}
		logs = append(logs, log)
	}
	return logs, nil
}

// ListByStoreID mengambil daftar activity log berdasarkan store ID.
func (r *activityLogRepositoryPG) ListByStoreID(ctx context.Context, storeID uuid.UUID) ([]*models.ActivityLog, error) {
	query := `
        SELECT id, user_id, company_id, store_id, action_type, description, target_entity, target_entity_id, ip_address, user_agent, log_time
        FROM activity_logs
        WHERE store_id = $1
        ORDER BY log_time DESC`
	rows, err := r.db.QueryContext(ctx, query, storeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var logs []*models.ActivityLog
	for rows.Next() {
		log := &models.ActivityLog{}
		if err := rows.Scan(
			&log.ID, &log.UserID, &log.CompanyID, &log.StoreID, &log.ActionType, &log.Description,
			&log.TargetEntity, &log.TargetEntityID, &log.IPAddress, &log.UserAgent, &log.LogTime,
		); err != nil {
			return nil, err
		}
		logs = append(logs, log)
	}
	return logs, nil
}
