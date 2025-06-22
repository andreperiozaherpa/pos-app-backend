package postgres

import (
	"context"
	"database/sql"
	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

// StoreProductRepository mendefinisikan interface untuk operasi data terkait Product (store_product).
type StoreProductRepository interface {
	Create(ctx context.Context, product *models.StoreProduct) error
	GetByID(ctx context.Context, id uuid.UUID) (*models.StoreProduct, error) // Mengembalikan StoreProduct
	GetByStoreAndMasterProduct(ctx context.Context, storeID, masterProductID uuid.UUID) (*models.StoreProduct, error)
	ListByStore(ctx context.Context, storeID uuid.UUID) ([]*models.StoreProduct, error)
	Update(ctx context.Context, product *models.StoreProduct) error
	Delete(ctx context.Context, id uuid.UUID) error
}

// pgStoreProductRepository adalah implementasi dari StoreProductRepository untuk PostgreSQL.
type pgStoreProductRepository struct {
	db DBExecutor
}

// NewPgStoreProductRepository adalah constructor untuk membuat instance baru dari pgStoreProductRepository.
func NewPgStoreProductRepository(db DBExecutor) StoreProductRepository {
	return &pgStoreProductRepository{db: db}
}

// Implementasi metode-metode dari interface StoreProductRepository:

func (r *pgStoreProductRepository) Create(ctx context.Context, p *models.StoreProduct) error {
	query := `
		INSERT INTO store_products (id, master_product_id, store_id, supplier_id, store_specific_sku, 
			purchase_price, selling_price, wholesale_price, stock, minimum_stock_level, expiry_date, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)` // 13 placeholders
	_, err := r.db.ExecContext(ctx, query,
		p.ID, p.MasterProductID, p.StoreID, p.SupplierID, p.StoreSpecificSKU,
		p.PurchasePrice, p.SellingPrice, p.WholesalePrice, p.Stock, p.MinimumStockLevel,
		p.ExpiryDate, p.CreatedAt, p.UpdatedAt, // 13 parameters
	)
	return err
}

func (r *pgStoreProductRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.StoreProduct, error) {
	p := &models.StoreProduct{}
	query := `
		SELECT id, master_product_id, store_id, supplier_id, store_specific_sku, 
			purchase_price, selling_price, wholesale_price, stock, minimum_stock_level, expiry_date, created_at, updated_at
		FROM store_products
		WHERE id = $1`
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&p.ID, &p.MasterProductID, &p.StoreID, &p.SupplierID, &p.StoreSpecificSKU,
		&p.PurchasePrice, &p.SellingPrice, &p.WholesalePrice, &p.Stock, &p.MinimumStockLevel,
		&p.ExpiryDate, &p.CreatedAt, &p.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, err
	}
	return p, nil
}

func (r *pgStoreProductRepository) GetByStoreAndMasterProduct(ctx context.Context, storeID, masterProductID uuid.UUID) (*models.StoreProduct, error) {
	p := &models.StoreProduct{}
	query := `
		SELECT id, master_product_id, store_id, supplier_id, store_specific_sku, 
			purchase_price, selling_price, wholesale_price, stock, minimum_stock_level, expiry_date, created_at, updated_at
		FROM store_products
		WHERE store_id = $1 AND master_product_id = $2`
	err := r.db.QueryRowContext(ctx, query, storeID, masterProductID).Scan(
		&p.ID, &p.MasterProductID, &p.StoreID, &p.SupplierID, &p.StoreSpecificSKU,
		&p.PurchasePrice, &p.SellingPrice, &p.WholesalePrice, &p.Stock, &p.MinimumStockLevel,
		&p.ExpiryDate, &p.CreatedAt, &p.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, err
	}
	return p, nil
}

func (r *pgStoreProductRepository) ListByStore(ctx context.Context, storeID uuid.UUID) ([]*models.StoreProduct, error) {
	query := `
		SELECT id, master_product_id, store_id, supplier_id, store_specific_sku, 
			purchase_price, selling_price, wholesale_price, stock, minimum_stock_level, expiry_date, created_at, updated_at
		FROM store_products
		WHERE store_id = $1
		ORDER BY store_specific_sku ASC`
	rows, err := r.db.QueryContext(ctx, query, storeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*models.StoreProduct
	for rows.Next() {
		p := &models.StoreProduct{}
		if err := rows.Scan(
			&p.ID, &p.MasterProductID, &p.StoreID, &p.SupplierID, &p.StoreSpecificSKU,
			&p.PurchasePrice, &p.SellingPrice, &p.WholesalePrice, &p.Stock, &p.MinimumStockLevel,
			&p.ExpiryDate, &p.CreatedAt, &p.UpdatedAt,
		); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, rows.Err()
}

func (r *pgStoreProductRepository) Update(ctx context.Context, p *models.StoreProduct) error {
	query := `
		UPDATE store_products
		SET supplier_id = $1, store_specific_sku = $2, purchase_price = $3, selling_price = $4, 
			wholesale_price = $5, stock = $6, minimum_stock_level = $7, expiry_date = $8, updated_at = $9
		WHERE id = $10`
	_, err := r.db.ExecContext(ctx, query,
		p.SupplierID, p.StoreSpecificSKU,
		p.PurchasePrice, p.SellingPrice, p.WholesalePrice, p.Stock, p.MinimumStockLevel,
		p.ExpiryDate, p.UpdatedAt, p.ID,
	)
	return err
}

func (r *pgStoreProductRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM store_products WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
