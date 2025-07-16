-- 20240716095000_stock_and_purchasing.sql

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
-- 1. purchase_orders
CREATE TABLE IF NOT EXISTS purchase_orders (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4 (),
    store_id UUID NOT NULL REFERENCES stores (id),
    supplier_id UUID NOT NULL REFERENCES suppliers (id),
    order_date DATE NOT NULL,
    expected_delivery_date DATE,
    status VARCHAR(50) NOT NULL CHECK (
        status IN (
            'PENDING',
            'ORDERED',
            'PARTIALLY_RECEIVED',
            'RECEIVED',
            'CANCELLED'
        )
    ),
    total_amount DECIMAL(15, 2),
    notes TEXT,
    created_by_user_id UUID NOT NULL REFERENCES users (id),
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now()
);

-- 2. purchase_order_items
CREATE TABLE IF NOT EXISTS purchase_order_items (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4 (),
    purchase_order_id UUID NOT NULL REFERENCES purchase_orders (id),
    master_product_id UUID NOT NULL REFERENCES master_products (id),
    quantity_ordered INTEGER NOT NULL,
    purchase_price_per_unit DECIMAL(15, 2) NOT NULL,
    quantity_received INTEGER DEFAULT 0,
    subtotal DECIMAL(15, 2) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now()
);

-- 3. internal_stock_transfers
CREATE TABLE IF NOT EXISTS internal_stock_transfers (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4 (),
    transfer_code VARCHAR(100) UNIQUE NOT NULL,
    company_id UUID NOT NULL REFERENCES companies (id),
    source_store_id UUID NOT NULL REFERENCES stores (id),
    destination_store_id UUID NOT NULL REFERENCES stores (id),
    transfer_date DATE NOT NULL,
    status VARCHAR(50) NOT NULL CHECK (
        status IN (
            'PENDING',
            'APPROVED',
            'SHIPPED',
            'PARTIALLY_RECEIVED',
            'RECEIVED',
            'CANCELLED'
        )
    ),
    notes TEXT,
    requested_by_user_id UUID REFERENCES users (id),
    approved_by_user_id UUID REFERENCES users (id),
    shipped_by_user_id UUID REFERENCES users (id),
    received_by_user_id UUID REFERENCES users (id),
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now()
);

-- 4. internal_stock_transfer_items
CREATE TABLE IF NOT EXISTS internal_stock_transfer_items (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4 (),
    internal_stock_transfer_id UUID NOT NULL REFERENCES internal_stock_transfers (id),
    source_store_product_id UUID NOT NULL REFERENCES store_products (id),
    quantity_requested INTEGER NOT NULL,
    quantity_shipped INTEGER DEFAULT 0,
    quantity_received INTEGER DEFAULT 0,
    notes TEXT,
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now()
);

-- 5. stock_movements
CREATE TABLE IF NOT EXISTS stock_movements (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4 (),
    store_product_id UUID NOT NULL REFERENCES store_products (id),
    store_id UUID NOT NULL REFERENCES stores (id),
    movement_type VARCHAR(50) NOT NULL CHECK (
        movement_type IN (
            'SALE',
            'PURCHASE_RECEIPT',
            'RETURN_TO_SUPPLIER',
            'CUSTOMER_RETURN',
            'ADJUSTMENT_IN',
            'ADJUSTMENT_OUT',
            'TRANSFER_OUT',
            'TRANSFER_IN'
        )
    ),
    quantity_changed INTEGER NOT NULL,
    movement_date TIMESTAMPTZ DEFAULT now() NOT NULL,
    reference_id UUID,
    reference_type VARCHAR(100),
    notes TEXT,
    created_by_user_id UUID REFERENCES users (id),
    created_at TIMESTAMPTZ DEFAULT now()
);

-- 6. store_product_stock_updates
CREATE TABLE IF NOT EXISTS store_product_stock_updates (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4 (),
    store_product_id UUID NOT NULL REFERENCES store_products (id),
    adjustment_type VARCHAR(50) NOT NULL,
    quantity_before INTEGER NOT NULL,
    quantity_after INTEGER NOT NULL,
    adjusted_by_user_id UUID NOT NULL REFERENCES users (id),
    reason TEXT,
    adjustment_date TIMESTAMPTZ NOT NULL DEFAULT now(),
    created_at TIMESTAMPTZ DEFAULT now()
);

-- 7. stock_movement_summaries
CREATE TABLE IF NOT EXISTS stock_movement_summaries (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4 (),
    store_product_id UUID NOT NULL REFERENCES store_products (id),
    period_start DATE NOT NULL,
    period_end DATE NOT NULL,
    total_in INTEGER NOT NULL DEFAULT 0,
    total_out INTEGER NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ DEFAULT now()
);

-- 8. stock_reports
CREATE TABLE IF NOT EXISTS stock_reports (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4 (),
    store_id UUID NOT NULL REFERENCES stores (id),
    report_date DATE NOT NULL,
    total_stock_value DECIMAL(18, 2) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT now()
);

-- 9. stock_transfer_histories
CREATE TABLE IF NOT EXISTS stock_transfer_histories (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4 (),
    stock_transfer_id UUID NOT NULL REFERENCES internal_stock_transfers (id),
    action VARCHAR(100) NOT NULL,
    action_date TIMESTAMPTZ NOT NULL DEFAULT now(),
    created_at TIMESTAMPTZ DEFAULT now()
);