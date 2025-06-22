package postgres

import (
	"context"
	"database/sql"
	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

// StoreRepository mendefinisikan interface untuk operasi data terkait Store.
type StoreRepository interface {
	Create(ctx context.Context, store *models.Store) error
	GetByID(ctx context.Context, id uuid.UUID) (*models.Store, error)
	Update(ctx context.Context, store *models.Store) error
	Delete(ctx context.Context, id uuid.UUID) error
	// Metode ini akan berguna untuk menampilkan semua toko dalam satu lini bisnis.
	ListByBusinessLine(ctx context.Context, businessLineID uuid.UUID) ([]*models.Store, error)
}

// pgStoreRepository adalah implementasi dari StoreRepository untuk PostgreSQL.
type pgStoreRepository struct {
	db *sql.DB
}

// NewPgStoreRepository adalah constructor untuk membuat instance baru dari pgStoreRepository.
func NewPgStoreRepository(db *sql.DB) StoreRepository {
	return &pgStoreRepository{db: db}
}

// Implementasi metode-metode dari interface StoreRepository:

func (r *pgStoreRepository) Create(ctx context.Context, store *models.Store) error {
	query := `
		INSERT INTO stores (id, business_line_id, parent_store_id, name, store_code, store_type, address, phone_number, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`

	_, err := r.db.ExecContext(ctx, query,
		store.ID,
		store.BusinessLineID,
		store.ParentStoreID,
		store.Name,
		store.StoreCode,
		store.StoreType,
		store.Address,
		store.PhoneNumber,
		store.CreatedAt,
		store.UpdatedAt,
	)
	return err
}

func (r *pgStoreRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.Store, error) {
	store := &models.Store{}
	query := `
		SELECT id, business_line_id, parent_store_id, name, store_code, store_type, address, phone_number, created_at, updated_at
		FROM stores
		WHERE id = $1`

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&store.ID,
		&store.BusinessLineID,
		&store.ParentStoreID,
		&store.Name,
		&store.StoreCode,
		&store.StoreType,
		&store.Address,
		&store.PhoneNumber,
		&store.CreatedAt,
		&store.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, err
	}
	return store, nil
}

func (r *pgStoreRepository) Update(ctx context.Context, store *models.Store) error {
	query := `
		UPDATE stores
		SET business_line_id = $1, parent_store_id = $2, name = $3, store_code = $4, store_type = $5, address = $6, phone_number = $7, updated_at = $8
		WHERE id = $9`

	_, err := r.db.ExecContext(ctx, query,
		store.BusinessLineID,
		store.ParentStoreID,
		store.Name,
		store.StoreCode,
		store.StoreType,
		store.Address,
		store.PhoneNumber,
		store.UpdatedAt,
		store.ID,
	)
	return err
}

func (r *pgStoreRepository) Delete(ctx context.Context, id uuid.UUID) error {
	// Implementasi saat ini adalah hard delete.
	query := `DELETE FROM stores WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

func (r *pgStoreRepository) ListByBusinessLine(ctx context.Context, businessLineID uuid.UUID) ([]*models.Store, error) {
	query := `
		SELECT id, business_line_id, parent_store_id, name, store_code, store_type, address, phone_number, created_at, updated_at
		FROM stores
		WHERE business_line_id = $1
		ORDER BY name ASC`

	rows, err := r.db.QueryContext(ctx, query, businessLineID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stores []*models.Store
	for rows.Next() {
		store := &models.Store{}
		if err := rows.Scan(
			&store.ID,
			&store.BusinessLineID,
			&store.ParentStoreID,
			&store.Name,
			&store.StoreCode,
			&store.StoreType,
			&store.Address,
			&store.PhoneNumber,
			&store.CreatedAt,
			&store.UpdatedAt,
		); err != nil {
			return nil, err
		}
		stores = append(stores, store)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return stores, nil
}
