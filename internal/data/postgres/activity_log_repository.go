package postgres

import (
	"context"
	"fmt"
	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

// ActivityLogRepository mendefinisikan interface untuk operasi data terkait ActivityLog.
type ActivityLogRepository interface {
	Create(ctx context.Context, log *models.ActivityLog) error
	ListByUserID(ctx context.Context, userID uuid.UUID) ([]*models.ActivityLog, error)
	ListByCompanyID(ctx context.Context, companyID uuid.UUID) ([]*models.ActivityLog, error)
	ListByStoreID(ctx context.Context, storeID uuid.UUID) ([]*models.ActivityLog, error)
}

// pgActivityLogRepository adalah implementasi dari ActivityLogRepository untuk PostgreSQL.
type pgActivityLogRepository struct {
	db DBExecutor
}

// NewPgActivityLogRepository adalah constructor untuk membuat instance baru dari pgActivityLogRepository.
func NewPgActivityLogRepository(db DBExecutor) ActivityLogRepository {
	return &pgActivityLogRepository{db: db}
}

// Create menyisipkan catatan aktivitas baru.
func (r *pgActivityLogRepository) Create(ctx context.Context, log *models.ActivityLog) error {
	query := `
		INSERT INTO activity_logs (user_id, company_id, store_id, action_type, description,
			target_entity, target_entity_id, ip_address, user_agent, log_time)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		RETURNING id` // Mengembalikan ID yang di-generate otomatis

	err := r.db.QueryRowContext(ctx, query,
		log.UserID, log.CompanyID, log.StoreID, log.ActionType, log.Description,
		log.TargetEntity, log.TargetEntityID, log.IPAddress, log.UserAgent, log.LogTime,
	).Scan(&log.ID) // Memindai ID yang dikembalikan ke struct log

	if err != nil {
		return fmt.Errorf("gagal menyisipkan activity log: %w", err)
	}
	return nil
}

// ListByUserID mengambil daftar log aktivitas berdasarkan ID pengguna.
func (r *pgActivityLogRepository) ListByUserID(ctx context.Context, userID uuid.UUID) ([]*models.ActivityLog, error) {
	query := `
		SELECT id, user_id, company_id, store_id, action_type, description,
			target_entity, target_entity_id, ip_address, user_agent, log_time
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
	return logs, rows.Err()
}

// ListByCompanyID mengambil daftar log aktivitas berdasarkan ID perusahaan.
func (r *pgActivityLogRepository) ListByCompanyID(ctx context.Context, companyID uuid.UUID) ([]*models.ActivityLog, error) {
	query := `
		SELECT id, user_id, company_id, store_id, action_type, description,
			target_entity, target_entity_id, ip_address, user_agent, log_time
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
	return logs, rows.Err()
}

// ListByStoreID mengambil daftar log aktivitas berdasarkan ID toko.
func (r *pgActivityLogRepository) ListByStoreID(ctx context.Context, storeID uuid.UUID) ([]*models.ActivityLog, error) {
	query := `
		SELECT id, user_id, company_id, store_id, action_type, description,
			target_entity, target_entity_id, ip_address, user_agent, log_time
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
	return logs, rows.Err()
}
