package postgres

import (
	"context"
	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

// StockMovementRepository mendefinisikan interface untuk operasi data terkait StockMovement.
type StockMovementRepository interface {
	Create(ctx context.Context, movement *models.StockMovement) error
	ListByStoreProduct(ctx context.Context, storeProductID uuid.UUID) ([]*models.StockMovement, error)
}

// pgStockMovementRepository adalah implementasi dari StockMovementRepository untuk PostgreSQL.
type pgStockMovementRepository struct {
	db DBExecutor
}

// NewPgStockMovementRepository adalah constructor untuk membuat instance baru dari pgStockMovementRepository.
func NewPgStockMovementRepository(db DBExecutor) StockMovementRepository {
	return &pgStockMovementRepository{db: db}
}

// Create menyisipkan catatan pergerakan stok baru.
func (r *pgStockMovementRepository) Create(ctx context.Context, sm *models.StockMovement) error {
	query := `
		INSERT INTO stock_movements (id, store_product_id, store_id, movement_type, quantity_changed,
			movement_date, reference_id, reference_type, notes, created_by_user_id, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`
	_, err := r.db.ExecContext(ctx, query,
		sm.ID, sm.StoreProductID, sm.StoreID, sm.MovementType, sm.QuantityChanged, sm.MovementDate,
		sm.ReferenceID, sm.ReferenceType, sm.Notes, sm.CreatedByUserID, sm.CreatedAt,
	)
	return err
}

// ListByStoreProduct mengambil daftar pergerakan stok untuk produk tertentu di sebuah toko.
func (r *pgStockMovementRepository) ListByStoreProduct(ctx context.Context, storeProductID uuid.UUID) ([]*models.StockMovement, error) {
	query := `
		SELECT id, store_product_id, store_id, movement_type, quantity_changed, movement_date,
			reference_id, reference_type, notes, created_by_user_id, created_at
		FROM stock_movements
		WHERE store_product_id = $1
		ORDER BY movement_date DESC`
	rows, err := r.db.QueryContext(ctx, query, storeProductID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movements []*models.StockMovement
	for rows.Next() {
		sm := &models.StockMovement{}
		if err := rows.Scan(
			&sm.ID, &sm.StoreProductID, &sm.StoreID, &sm.MovementType, &sm.QuantityChanged, &sm.MovementDate,
			&sm.ReferenceID, &sm.ReferenceType, &sm.Notes, &sm.CreatedByUserID, &sm.CreatedAt,
		); err != nil {
			return nil, err
		}
		movements = append(movements, sm)
	}
	return movements, rows.Err()
}
