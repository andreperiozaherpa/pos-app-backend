-- 20240716092000_tax_and_product.sql
-- Modul tax_rates, master_products, store_products

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
-- 1. tax_rates
CREATE TABLE IF NOT EXISTS tax_rates (
    id SERIAL PRIMARY KEY,
    company_id UUID NOT NULL REFERENCES companies (id),
    name VARCHAR(100) NOT NULL,
    rate_percentage DECIMAL(5, 2) NOT NULL,
    description TEXT,
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now()
);

-- 2. master_products
CREATE TABLE IF NOT EXISTS master_products (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4 (),
    company_id UUID NOT NULL REFERENCES companies (id),
    master_product_code VARCHAR(100) NOT NULL,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    category VARCHAR(100),
    unit_of_measure VARCHAR(50),
    barcode VARCHAR(255),
    default_tax_rate_id INTEGER REFERENCES tax_rates (id),
    image_url VARCHAR(255),
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now(),
    CONSTRAINT uq_master_product_code UNIQUE (
        company_id,
        master_product_code
    )
);

-- 3. store_products
CREATE TABLE IF NOT EXISTS store_products (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4 (),
    master_product_id UUID NOT NULL REFERENCES master_products (id),
    store_id UUID NOT NULL REFERENCES stores (id),
    supplier_id UUID REFERENCES suppliers (id),
    store_specific_sku VARCHAR(100),
    purchase_price DECIMAL(15, 2) NOT NULL,
    selling_price DECIMAL(15, 2) NOT NULL,
    wholesale_price DECIMAL(15, 2),
    stock INTEGER NOT NULL DEFAULT 0,
    minimum_stock_level INTEGER,
    expiry_date DATE,
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now(),
    CONSTRAINT uq_store_product UNIQUE (store_id, master_product_id),
    CONSTRAINT uq_store_sku UNIQUE (store_id, store_specific_sku)
);