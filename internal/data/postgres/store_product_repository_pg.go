package postgres

import (
	"context"
	"database/sql"
	"errors"
	"pos-app/backend/internal/core/repository"
	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

type storeProductRepositoryPG struct {
	db *sql.DB
}

func NewStoreProductRepositoryPG(db *sql.DB) repository.StoreProductRepository {
	return &storeProductRepositoryPG{db: db}
}

func (r *storeProductRepositoryPG) Create(ctx context.Context, sp *models.StoreProduct) error {
	query := `
		INSERT INTO store_products (
			id, store_id, master_product_id, supplier_id, store_specific_sku,
			purchase_price, selling_price, wholesale_price, stock,
			minimum_stock_level, expiry_date, created_at, updated_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, NOW(), NOW())`
	_, err := r.db.ExecContext(ctx, query,
		sp.ID, sp.StoreID, sp.MasterProductID, sp.SupplierID, sp.StoreSpecificSKU,
		sp.PurchasePrice, sp.SellingPrice, sp.WholesalePrice, sp.Stock,
		sp.MinimumStockLevel, sp.ExpiryDate,
	)
	return err
}

func (r *storeProductRepositoryPG) GetByID(ctx context.Context, id uuid.UUID) (*models.StoreProduct, error) {
	query := `
		SELECT id, store_id, master_product_id, supplier_id, store_specific_sku,
		       purchase_price, selling_price, wholesale_price, stock,
		       minimum_stock_level, expiry_date, created_at, updated_at
		FROM store_products
		WHERE id = $1`
	row := r.db.QueryRowContext(ctx, query, id)
	sp := &models.StoreProduct{}
	err := row.Scan(
		&sp.ID, &sp.StoreID, &sp.MasterProductID, &sp.SupplierID, &sp.StoreSpecificSKU,
		&sp.PurchasePrice, &sp.SellingPrice, &sp.WholesalePrice, &sp.Stock,
		&sp.MinimumStockLevel, &sp.ExpiryDate, &sp.CreatedAt, &sp.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}
	return sp, nil
}

func (r *storeProductRepositoryPG) Update(ctx context.Context, sp *models.StoreProduct) error {
	query := `
		UPDATE store_products SET
			supplier_id = $1,
			store_specific_sku = $2,
			purchase_price = $3,
			selling_price = $4,
			wholesale_price = $5,
			stock = $6,
			minimum_stock_level = $7,
			expiry_date = $8,
			updated_at = NOW()
		WHERE id = $9`
	result, err := r.db.ExecContext(ctx, query,
		sp.SupplierID, sp.StoreSpecificSKU, sp.PurchasePrice, sp.SellingPrice, sp.WholesalePrice,
		sp.Stock, sp.MinimumStockLevel, sp.ExpiryDate, sp.ID,
	)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return repository.ErrNotFound
	}
	return nil
}

func (r *storeProductRepositoryPG) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM store_products WHERE id = $1`
	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return repository.ErrNotFound
	}
	return nil
}

// UpdateStock mengubah stok produk toko
func (r *storeProductRepositoryPG) UpdateStock(ctx context.Context, storeProductID uuid.UUID, quantity int) error {
	query := `UPDATE store_products SET stock = $1, updated_at = NOW() WHERE id = $2`
	result, err := r.db.ExecContext(ctx, query, quantity, storeProductID)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return repository.ErrNotFound
	}
	return nil
}

// UpdatePrice mengubah harga jual produk toko
func (r *storeProductRepositoryPG) UpdatePrice(ctx context.Context, storeProductID uuid.UUID, newSellingPrice float64) error {
	query := `UPDATE store_products SET selling_price = $1, updated_at = NOW() WHERE id = $2`
	result, err := r.db.ExecContext(ctx, query, newSellingPrice, storeProductID)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return repository.ErrNotFound
	}
	return nil
}

// ListByStore mengambil semua produk berdasarkan store_id.
func (r *storeProductRepositoryPG) ListByStore(ctx context.Context, storeID uuid.UUID) ([]*models.StoreProduct, error) {
	query := `
		SELECT id, store_id, master_product_id, supplier_id, store_specific_sku,
		       purchase_price, selling_price, wholesale_price, stock,
		       minimum_stock_level, expiry_date, created_at, updated_at
		FROM store_products
		WHERE store_id = $1
		ORDER BY store_specific_sku`
	rows, err := r.db.QueryContext(ctx, query, storeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*models.StoreProduct
	for rows.Next() {
		sp := &models.StoreProduct{}
		if err := rows.Scan(
			&sp.ID, &sp.StoreID, &sp.MasterProductID, &sp.SupplierID, &sp.StoreSpecificSKU,
			&sp.PurchasePrice, &sp.SellingPrice, &sp.WholesalePrice, &sp.Stock,
			&sp.MinimumStockLevel, &sp.ExpiryDate, &sp.CreatedAt, &sp.UpdatedAt,
		); err != nil {
			return nil, err
		}
		products = append(products, sp)
	}
	return products, nil
}

// GetByStoreAndMasterProduct mengambil data produk toko berdasarkan store_id dan master_product_id.
func (r *storeProductRepositoryPG) GetByStoreAndMasterProduct(ctx context.Context, storeID, masterProductID uuid.UUID) (*models.StoreProduct, error) {
	query := `
		SELECT id, store_id, master_product_id, supplier_id, store_specific_sku,
		       purchase_price, selling_price, wholesale_price, stock,
		       minimum_stock_level, expiry_date, created_at, updated_at
		FROM store_products
		WHERE store_id = $1 AND master_product_id = $2`
	row := r.db.QueryRowContext(ctx, query, storeID, masterProductID)
	sp := &models.StoreProduct{}
	err := row.Scan(
		&sp.ID, &sp.StoreID, &sp.MasterProductID, &sp.SupplierID, &sp.StoreSpecificSKU,
		&sp.PurchasePrice, &sp.SellingPrice, &sp.WholesalePrice, &sp.Stock,
		&sp.MinimumStockLevel, &sp.ExpiryDate, &sp.CreatedAt, &sp.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}
	return sp, nil
}
