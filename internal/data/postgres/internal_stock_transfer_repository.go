package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

// InternalStockTransferRepository mendefinisikan interface untuk operasi data terkait InternalStockTransfer.
type InternalStockTransferRepository interface {
	Create(ctx context.Context, ist *models.InternalStockTransfer) error
	GetByID(ctx context.Context, id uuid.UUID) (*models.InternalStockTransfer, error)
	ListByCompanyID(ctx context.Context, companyID uuid.UUID) ([]*models.InternalStockTransfer, error)
	Update(ctx context.Context, ist *models.InternalStockTransfer) error
	Delete(ctx context.Context, id uuid.UUID) error
}

// pgInternalStockTransferRepository adalah implementasi dari InternalStockTransferRepository untuk PostgreSQL.
type pgInternalStockTransferRepository struct {
	db DBExecutor
}

// NewPgInternalStockTransferRepository adalah constructor untuk membuat instance baru dari pgInternalStockTransferRepository.
func NewPgInternalStockTransferRepository(db DBExecutor) InternalStockTransferRepository {
	return &pgInternalStockTransferRepository{db: db}
}

// Implementasi metode-metode dari interface InternalStockTransferRepository:

func (r *pgInternalStockTransferRepository) Create(ctx context.Context, ist *models.InternalStockTransfer) error {
	tx, isTx := r.db.(*sql.Tx)
	var err error
	if !isTx {
		db, ok := r.db.(*sql.DB)
		if !ok {
			return fmt.Errorf("unexpected DBExecutor type; expected *sql.DB or *sql.Tx")
		}
		tx, err = db.BeginTx(ctx, nil)
	}

	if err != nil {
		return fmt.Errorf("gagal memulai transaksi: %w", err)
	}
	defer tx.Rollback() // Rollback jika ada error atau panic

	// 1. Sisipkan data ke tabel 'internal_stock_transfers' (header)
	istQuery := `
		INSERT INTO internal_stock_transfers (id, transfer_code, company_id, source_store_id, destination_store_id,
			transfer_date, status, notes, requested_by_user_id, approved_by_user_id, shipped_by_user_id,
			received_by_user_id, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)`
	_, err = tx.ExecContext(ctx, istQuery,
		ist.ID, ist.TransferCode, ist.CompanyID, ist.SourceStoreID, ist.DestinationStoreID,
		ist.TransferDate, ist.Status, ist.Notes, ist.RequestedByUserID, ist.ApprovedByUserID,
		ist.ShippedByUserID, ist.ReceivedByUserID, ist.CreatedAt, ist.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("gagal menyisipkan header internal stock transfer: %w", err)
	}

	// 2. Sisipkan setiap item ke tabel 'internal_stock_transfer_items'
	itemQuery := `
		INSERT INTO internal_stock_transfer_items (id, internal_stock_transfer_id, source_store_product_id,
			quantity_requested, quantity_shipped, quantity_received, notes, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	for _, item := range ist.Items {
		_, err = tx.ExecContext(ctx, itemQuery,
			item.ID, ist.ID, item.SourceStoreProductID,
			item.QuantityRequested, item.QuantityShipped, item.QuantityReceived, item.Notes, item.CreatedAt, item.UpdatedAt,
		)
		if err != nil {
			return fmt.Errorf("gagal menyisipkan item internal stock transfer (produk ID %s): %w", item.SourceStoreProductID, err)
		}
	}

	// 3. Commit transaksi
	if !isTx {
		return tx.Commit()
	}
	return nil
}

func (r *pgInternalStockTransferRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.InternalStockTransfer, error) {
	ist := &models.InternalStockTransfer{}
	// Query untuk header internal stock transfer
	istQuery := `
		SELECT id, transfer_code, company_id, source_store_id, destination_store_id,
			transfer_date, status, notes, requested_by_user_id, approved_by_user_id, shipped_by_user_id,
			received_by_user_id, created_at, updated_at
		FROM internal_stock_transfers
		WHERE id = $1`
	err := r.db.QueryRowContext(ctx, istQuery, id).Scan(
		&ist.ID, &ist.TransferCode, &ist.CompanyID, &ist.SourceStoreID, &ist.DestinationStoreID,
		&ist.TransferDate, &ist.Status, &ist.Notes, &ist.RequestedByUserID, &ist.ApprovedByUserID,
		&ist.ShippedByUserID, &ist.ReceivedByUserID, &ist.CreatedAt, &ist.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, err
	}

	// Query untuk item-item internal stock transfer
	itemQuery := `
		SELECT id, internal_stock_transfer_id, source_store_product_id,
			quantity_requested, quantity_shipped, quantity_received, notes, created_at, updated_at
		FROM internal_stock_transfer_items
		WHERE internal_stock_transfer_id = $1`
	rows, err := r.db.QueryContext(ctx, itemQuery, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []models.InternalStockTransferItem
	for rows.Next() {
		item := models.InternalStockTransferItem{}
		if err := rows.Scan(
			&item.ID, &item.InternalStockTransferID, &item.SourceStoreProductID,
			&item.QuantityRequested, &item.QuantityShipped, &item.QuantityReceived, &item.Notes, &item.CreatedAt, &item.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	ist.Items = items

	return ist, rows.Err()
}

func (r *pgInternalStockTransferRepository) ListByCompanyID(ctx context.Context, companyID uuid.UUID) ([]*models.InternalStockTransfer, error) {
	query := `
		SELECT id, transfer_code, company_id, source_store_id, destination_store_id,
			transfer_date, status, created_at
		FROM internal_stock_transfers
		WHERE company_id = $1
		ORDER BY transfer_date DESC`
	rows, err := r.db.QueryContext(ctx, query, companyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transfers []*models.InternalStockTransfer
	for rows.Next() {
		ist := &models.InternalStockTransfer{}
		if err := rows.Scan(
			&ist.ID, &ist.TransferCode, &ist.CompanyID, &ist.SourceStoreID, &ist.DestinationStoreID,
			&ist.TransferDate, &ist.Status, &ist.CreatedAt,
		); err != nil {
			return nil, err
		}
		transfers = append(transfers, ist)
	}
	return transfers, rows.Err()
}

func (r *pgInternalStockTransferRepository) Update(ctx context.Context, ist *models.InternalStockTransfer) error {
	query := `
		UPDATE internal_stock_transfers
		SET transfer_code = $1, company_id = $2, source_store_id = $3, destination_store_id = $4,
			transfer_date = $5, status = $6, notes = $7, requested_by_user_id = $8, approved_by_user_id = $9,
			shipped_by_user_id = $10, received_by_user_id = $11, updated_at = $12
		WHERE id = $13`
	_, err := r.db.ExecContext(ctx, query,
		ist.TransferCode, ist.CompanyID, ist.SourceStoreID, ist.DestinationStoreID,
		ist.TransferDate, ist.Status, ist.Notes, ist.RequestedByUserID, ist.ApprovedByUserID,
		ist.ShippedByUserID, ist.ReceivedByUserID, ist.UpdatedAt, ist.ID,
	)
	return err
}

func (r *pgInternalStockTransferRepository) Delete(ctx context.Context, id uuid.UUID) error {
	tx, isTx := r.db.(*sql.Tx)
	var err error
	if !isTx {
		db, ok := r.db.(*sql.DB)
		if !ok {
			return fmt.Errorf("unexpected DBExecutor type; expected *sql.DB or *sql.Tx")
		}
		tx, err = db.BeginTx(ctx, nil)
	}

	if err != nil {
		return fmt.Errorf("gagal memulai transaksi delete: %w", err)
	}
	defer tx.Rollback() // Rollback jika ada error

	// 1. Hapus item-item terkait dari internal_stock_transfer_items
	deleteItemsQuery := `DELETE FROM internal_stock_transfer_items WHERE internal_stock_transfer_id = $1`
	_, err = tx.ExecContext(ctx, deleteItemsQuery, id)
	if err != nil {
		return fmt.Errorf("gagal menghapus item internal stock transfer: %w", err)
	}

	// 2. Hapus internal_stock_transfer itu sendiri
	deleteISTQuery := `DELETE FROM internal_stock_transfers WHERE id = $1`
	_, err = tx.ExecContext(ctx, deleteISTQuery, id)
	if err != nil {
		return fmt.Errorf("gagal menghapus internal stock transfer: %w", err)
	}

	if !isTx {
		return tx.Commit()
	}
	return nil

}
