package postgres

import (
	"context"
	"database/sql"

	"pos-app/backend/internal/core/repository"
	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

type transactionAuditLogRepositoryPG struct {
	db *sql.DB
}

func NewTransactionAuditLogRepositoryPG(db *sql.DB) repository.TransactionAuditLogRepository {
	return &transactionAuditLogRepositoryPG{db: db}
}

func (r *transactionAuditLogRepositoryPG) Create(ctx context.Context, log *models.TransactionAuditLog) error {
	query := `
		INSERT INTO transaction_audit_logs
		(id, transaction_id, action_type, performed_by_user_id, performed_at, note)
		VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := r.db.ExecContext(ctx, query,
		log.ID, log.TransactionID, log.ActionType, log.PerformedByUserID, log.PerformedAt, log.Note,
	)
	return err
}

func (r *transactionAuditLogRepositoryPG) ListByTransactionID(ctx context.Context, transactionID uuid.UUID) ([]*models.TransactionAuditLog, error) {
	query := `
		SELECT id, transaction_id, action_type, performed_by_user_id, performed_at, note
		FROM transaction_audit_logs
		WHERE transaction_id = $1
		ORDER BY performed_at ASC`
	rows, err := r.db.QueryContext(ctx, query, transactionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var logs []*models.TransactionAuditLog
	for rows.Next() {
		var l models.TransactionAuditLog
		if err := rows.Scan(&l.ID, &l.TransactionID, &l.ActionType, &l.PerformedByUserID, &l.PerformedAt, &l.Note); err != nil {
			return nil, err
		}
		logs = append(logs, &l)
	}
	return logs, nil
}
