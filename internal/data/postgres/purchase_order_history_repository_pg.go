package postgres

import (
	"context"
	"database/sql"

	"pos-app/backend/internal/core/repository"
	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

type purchaseOrderHistoryRepositoryPG struct {
	db *sql.DB
}

func NewPurchaseOrderHistoryRepositoryPG(db *sql.DB) repository.PurchaseOrderHistoryRepository {
	return &purchaseOrderHistoryRepositoryPG{db: db}
}

func (r *purchaseOrderHistoryRepositoryPG) Create(ctx context.Context, history *models.PurchaseOrderHistory) error {
	query := `
        INSERT INTO purchase_order_histories (id, purchase_order_id, status, changed_by, changed_at, notes)
        VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := r.db.ExecContext(ctx, query,
		history.ID, history.PurchaseOrderID, history.Status,
		history.ChangedBy, history.ChangedAt, history.Notes)
	return err
}

func (r *purchaseOrderHistoryRepositoryPG) ListByPurchaseOrder(ctx context.Context, purchaseOrderID uuid.UUID) ([]*models.PurchaseOrderHistory, error) {
	query := `
        SELECT id, purchase_order_id, status, changed_by, changed_at, notes
        FROM purchase_order_histories
        WHERE purchase_order_id=$1
        ORDER BY changed_at DESC`
	rows, err := r.db.QueryContext(ctx, query, purchaseOrderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var histories []*models.PurchaseOrderHistory
	for rows.Next() {
		history := &models.PurchaseOrderHistory{}
		if err := rows.Scan(
			&history.ID, &history.PurchaseOrderID, &history.Status,
			&history.ChangedBy, &history.ChangedAt, &history.Notes); err != nil {
			return nil, err
		}
		histories = append(histories, history)
	}
	return histories, nil
}
