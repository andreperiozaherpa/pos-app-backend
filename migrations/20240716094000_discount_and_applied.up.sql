-- 20240716094000_discount_and_applied.sql

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
-- 1. discounts
CREATE TABLE IF NOT EXISTS discounts (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4 (),
    company_id UUID NOT NULL REFERENCES companies (id),
    name VARCHAR(255) NOT NULL,
    description TEXT,
    discount_type VARCHAR(50) NOT NULL CHECK (
        discount_type IN ('PERCENTAGE', 'FIXED_AMOUNT')
    ),
    discount_value DECIMAL(15, 2) NOT NULL,
    applicable_to VARCHAR(50) NOT NULL CHECK (
        applicable_to IN (
            'MASTER_PRODUCT',
            'STORE_PRODUCT',
            'CATEGORY',
            'TOTAL_TRANSACTION',
            'CUSTOMER_TIER'
        )
    ),
    master_product_id_applicable UUID REFERENCES master_products (id),
    store_product_id_applicable UUID REFERENCES store_products (id),
    category_applicable VARCHAR(100),
    customer_tier_applicable VARCHAR(50),
    min_purchase_amount DECIMAL(15, 2),
    start_date DATE NOT NULL,
    end_date DATE NOT NULL,
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now()
);

-- 2. applied_item_discounts
CREATE TABLE IF NOT EXISTS applied_item_discounts (
    transaction_item_id UUID NOT NULL REFERENCES transaction_items (id),
    discount_id UUID NOT NULL REFERENCES discounts (id),
    applied_discount_amount_on_item DECIMAL(15, 2) NOT NULL,
    PRIMARY KEY (
        transaction_item_id,
        discount_id
    )
);

-- 3. applied_transaction_discounts
CREATE TABLE IF NOT EXISTS applied_transaction_discounts (
    transaction_id UUID NOT NULL REFERENCES transactions (id),
    discount_id UUID NOT NULL REFERENCES discounts (id),
    applied_discount_amount_on_transaction DECIMAL(15, 2) NOT NULL,
    PRIMARY KEY (transaction_id, discount_id)
);

-- Optional: tambahkan index untuk mempercepat query jika diperlukan