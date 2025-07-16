-- 20240716096000_audit_and_operational.sql
-- Modul Audit & Keuangan: activity_logs, operational_expenses

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
-- ===========================
-- Table: activity_logs
-- ===========================
CREATE TABLE IF NOT EXISTS activity_logs (
    id BIGSERIAL PRIMARY KEY,
    user_id UUID REFERENCES users (id),
    company_id UUID REFERENCES companies (id),
    store_id UUID REFERENCES stores (id),
    action_type VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    target_entity VARCHAR(100),
    target_entity_id VARCHAR(255),
    ip_address VARCHAR(100),
    user_agent TEXT,
    log_time TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- ===========================
-- Table: operational_expenses
-- ===========================
CREATE TABLE IF NOT EXISTS operational_expenses (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4 (),
    company_id UUID NOT NULL REFERENCES companies (id),
    store_id UUID REFERENCES stores (id),
    expense_date DATE NOT NULL,
    category VARCHAR(100) NOT NULL,
    description TEXT,
    amount DECIMAL(15, 2) NOT NULL,
    created_by_user_id UUID REFERENCES users (id),
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now()
);