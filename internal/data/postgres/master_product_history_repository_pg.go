package postgres

import (
	"context"
	"database/sql"

	"pos-app/backend/internal/core/repository"
	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

type masterProductHistoryRepositoryPG struct {
	db *sql.DB
}

func NewMasterProductHistoryRepositoryPG(db *sql.DB) repository.MasterProductHistoryRepository {
	return &masterProductHistoryRepositoryPG{db: db}
}

// Create menyimpan histori perubahan untuk master product baru.
func (r *masterProductHistoryRepositoryPG) Create(ctx context.Context, history *models.MasterProductHistory) error {
	query := `
        INSERT INTO master_product_histories (id, master_product_id, changed_by, change_type, changed_at, notes)
        VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := r.db.ExecContext(ctx, query,
		history.ID, history.MasterProductID, history.ChangedBy,
		history.ChangeType, history.ChangedAt, history.Notes)
	return err
}

// ListByMasterProduct mengambil histori perubahan untuk master product tertentu.
func (r *masterProductHistoryRepositoryPG) ListByMasterProduct(ctx context.Context, masterProductID uuid.UUID) ([]*models.MasterProductHistory, error) {
	query := `
        SELECT id, master_product_id, changed_by, change_type, changed_at, notes
        FROM master_product_histories
        WHERE master_product_id=$1
        ORDER BY changed_at DESC`
	rows, err := r.db.QueryContext(ctx, query, masterProductID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var histories []*models.MasterProductHistory
	for rows.Next() {
		history := &models.MasterProductHistory{}
		if err := rows.Scan(
			&history.ID, &history.MasterProductID, &history.ChangedBy,
			&history.ChangeType, &history.ChangedAt, &history.Notes); err != nil {
			return nil, err
		}
		histories = append(histories, history)
	}
	return histories, nil
}
