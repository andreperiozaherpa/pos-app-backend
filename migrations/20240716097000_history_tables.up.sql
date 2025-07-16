-- 20240716097000_history_tables.sql
-- Modul Histori: master_product_histories, purchase_order_histories, stock_transfer_histories, user_login_histories

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
-- ===========================
-- Table: master_product_histories
-- ===========================
CREATE TABLE IF NOT EXISTS master_product_histories (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4 (),
    master_product_id UUID NOT NULL REFERENCES master_products (id),
    changed_by UUID NOT NULL REFERENCES users (id),
    change_type VARCHAR(50) NOT NULL, -- e.g. 'CREATE', 'UPDATE', 'DELETE'
    changed_at TIMESTAMPTZ DEFAULT now(),
    notes TEXT
);

-- ===========================
-- Table: purchase_order_histories
-- ===========================
CREATE TABLE IF NOT EXISTS purchase_order_histories (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4 (),
    purchase_order_id UUID NOT NULL REFERENCES purchase_orders (id),
    changed_by_user_id UUID NOT NULL REFERENCES users (id),
    previous_status VARCHAR(50),
    new_status VARCHAR(50),
    change_time TIMESTAMPTZ DEFAULT now(),
    notes TEXT
);

-- ===========================
-- Table: stock_transfer_histories
-- ===========================
CREATE TABLE IF NOT EXISTS stock_transfer_histories (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4 (),
    stock_transfer_id UUID NOT NULL REFERENCES internal_stock_transfers (id),
    action VARCHAR(100) NOT NULL, -- e.g., CREATED, APPROVED, SHIPPED, RECEIVED, CANCELLED
    action_date TIMESTAMPTZ NOT NULL DEFAULT now(),
    created_at TIMESTAMPTZ DEFAULT now()
);

-- ===========================
-- Table: user_login_histories
-- ===========================
CREATE TABLE IF NOT EXISTS user_login_histories (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4 (),
    user_id UUID NOT NULL REFERENCES users (id),
    login_time TIMESTAMPTZ NOT NULL DEFAULT now(),
    ip_address VARCHAR(100),
    device_info TEXT
);