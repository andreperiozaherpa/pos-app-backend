package postgres

import (
	"context"
	"database/sql"
	"errors"

	"pos-app/backend/internal/core/repository"
	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

type stockTransferRepositoryPG struct {
	db *sql.DB
}

func NewStockTransferRepositoryPG(db *sql.DB) repository.StockTransferRepository {
	return &stockTransferRepositoryPG{db: db}
}

func (r *stockTransferRepositoryPG) Create(ctx context.Context, transfer *models.InternalStockTransfer) error {
	query := `
        INSERT INTO internal_stock_transfers (
			id, transfer_code, company_id, source_store_id, destination_store_id, transfer_date,
			status, notes, requested_by_user_id, approved_by_user_id, shipped_by_user_id, received_by_user_id,
			created_at, updated_at
		) VALUES (
			$1, $2, $3, $4, $5, $6,
			$7, $8, $9, $10, $11, $12,
			NOW(), NOW()
		)`
	_, err := r.db.ExecContext(ctx, query,
		transfer.ID, transfer.TransferCode, transfer.CompanyID, transfer.SourceStoreID, transfer.DestinationStoreID, transfer.TransferDate,
		transfer.Status, transfer.Notes,
		transfer.RequestedByUserID, transfer.ApprovedByUserID, transfer.ShippedByUserID, transfer.ReceivedByUserID,
	)
	return err
}

func (r *stockTransferRepositoryPG) GetByID(ctx context.Context, id uuid.UUID) (*models.InternalStockTransfer, error) {
	query := `
        SELECT id, transfer_code, company_id, source_store_id, destination_store_id, transfer_date,
			   status, notes, requested_by_user_id, approved_by_user_id, shipped_by_user_id, received_by_user_id,
			   created_at, updated_at
        FROM internal_stock_transfers WHERE id=$1`
	row := r.db.QueryRowContext(ctx, query, id)
	st := &models.InternalStockTransfer{}
	err := row.Scan(
		&st.ID, &st.TransferCode, &st.CompanyID, &st.SourceStoreID, &st.DestinationStoreID, &st.TransferDate,
		&st.Status, &st.Notes, &st.RequestedByUserID, &st.ApprovedByUserID, &st.ShippedByUserID, &st.ReceivedByUserID,
		&st.CreatedAt, &st.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}
	return st, nil
}

func (r *stockTransferRepositoryPG) Update(ctx context.Context, transfer *models.InternalStockTransfer) error {
	query := `
        UPDATE internal_stock_transfers SET
			transfer_code = $1,
			company_id = $2,
			source_store_id = $3,
			destination_store_id = $4,
			transfer_date = $5,
			status = $6,
			notes = $7,
			requested_by_user_id = $8,
			approved_by_user_id = $9,
			shipped_by_user_id = $10,
			received_by_user_id = $11,
			updated_at = NOW()
        WHERE id=$12`
	res, err := r.db.ExecContext(ctx, query,
		transfer.TransferCode, transfer.CompanyID, transfer.SourceStoreID, transfer.DestinationStoreID, transfer.TransferDate,
		transfer.Status, transfer.Notes,
		transfer.RequestedByUserID, transfer.ApprovedByUserID, transfer.ShippedByUserID, transfer.ReceivedByUserID,
		transfer.ID,
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

func (r *stockTransferRepositoryPG) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM internal_stock_transfers WHERE id=$1`
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

func (r *stockTransferRepositoryPG) ListByStore(ctx context.Context, storeID uuid.UUID, limit, offset int) ([]*models.InternalStockTransfer, error) {
	query := `
        SELECT id, transfer_code, company_id, source_store_id, destination_store_id, transfer_date,
			   status, notes, requested_by_user_id, approved_by_user_id, shipped_by_user_id, received_by_user_id,
			   created_at, updated_at
        FROM internal_stock_transfers
        WHERE source_store_id = $1 OR destination_store_id = $1
        ORDER BY transfer_date DESC
        LIMIT $2 OFFSET $3`
	rows, err := r.db.QueryContext(ctx, query, storeID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transfers []*models.InternalStockTransfer
	for rows.Next() {
		st := &models.InternalStockTransfer{}
		if err := rows.Scan(
			&st.ID, &st.TransferCode, &st.CompanyID, &st.SourceStoreID, &st.DestinationStoreID, &st.TransferDate,
			&st.Status, &st.Notes, &st.RequestedByUserID, &st.ApprovedByUserID, &st.ShippedByUserID, &st.ReceivedByUserID,
			&st.CreatedAt, &st.UpdatedAt,
		); err != nil {
			return nil, err
		}
		transfers = append(transfers, st)
	}
	return transfers, nil
}
