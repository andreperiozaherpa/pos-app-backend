-- 20240716093000_transactions.sql
-- Modul transactions dan transaction_items

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
-- 1. transactions
CREATE TABLE IF NOT EXISTS transactions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4 (),
    transaction_code VARCHAR(255) NOT NULL UNIQUE,
    store_id UUID NOT NULL REFERENCES stores (id),
    cashier_employee_user_id UUID NOT NULL REFERENCES employees (user_id),
    customer_user_id UUID REFERENCES customers (user_id),
    active_shift_id UUID REFERENCES shifts (id),
    transaction_date TIMESTAMPTZ NOT NULL DEFAULT now(),
    subtotal_amount DECIMAL(15, 2) NOT NULL,
    total_item_discount_amount DECIMAL(15, 2) DEFAULT 0,
    subtotal_after_item_discounts DECIMAL(15, 2) NOT NULL,
    transaction_level_discount_amount DECIMAL(15, 2) DEFAULT 0,
    taxable_amount DECIMAL(15, 2) NOT NULL,
    total_tax_amount DECIMAL(15, 2) DEFAULT 0,
    final_total_amount DECIMAL(15, 2) NOT NULL,
    received_amount DECIMAL(15, 2) NOT NULL,
    change_amount DECIMAL(15, 2) NOT NULL,
    payment_method VARCHAR(50),
    notes TEXT,
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now()
);

-- 2. transaction_items
CREATE TABLE IF NOT EXISTS transaction_items (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4 (),
    transaction_id UUID NOT NULL REFERENCES transactions (id),
    store_product_id UUID NOT NULL REFERENCES store_products (id),
    quantity INTEGER NOT NULL,
    price_per_unit_at_transaction DECIMAL(15, 2) NOT NULL,
    item_subtotal_before_discount DECIMAL(15, 2) NOT NULL,
    item_discount_amount DECIMAL(15, 2) DEFAULT 0,
    item_subtotal_after_discount DECIMAL(15, 2) NOT NULL,
    applied_tax_rate_id INTEGER REFERENCES tax_rates (id),
    applied_tax_rate_percentage DECIMAL(5, 2),
    tax_amount_for_item DECIMAL(15, 2) DEFAULT 0,
    item_final_total DECIMAL(15, 2) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now()
);

-- 3. payment_info
CREATE TABLE IF NOT EXISTS payment_info (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4 (),
    transaction_id UUID NOT NULL REFERENCES transactions (id),
    payment_method VARCHAR(50) NOT NULL,
    amount DECIMAL(15, 2) NOT NULL,
    payment_date TIMESTAMPTZ DEFAULT now()
);