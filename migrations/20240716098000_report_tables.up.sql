-- 20240716098000_report_tables.sql
-- Modul report: sales_reports, stock_reports, profit_loss_reports, employee_performance_reports, customer_activity_reports, transaction_summaries

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
-- ===========================
-- Table: sales_reports
-- ===========================
CREATE TABLE IF NOT EXISTS sales_reports (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4 (),
    store_id UUID NOT NULL REFERENCES stores (id),
    report_date DATE NOT NULL,
    total_sales DECIMAL(15, 2) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT now()
);

-- ===========================
-- Table: stock_reports
-- ===========================
CREATE TABLE IF NOT EXISTS stock_reports (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4 (),
    store_id UUID NOT NULL REFERENCES stores (id),
    report_date DATE NOT NULL,
    total_stock_value DECIMAL(18, 2) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT now()
);

-- ===========================
-- Table: profit_loss_reports
-- ===========================
CREATE TABLE IF NOT EXISTS profit_loss_reports (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4 (),
    company_id UUID NOT NULL REFERENCES companies (id),
    report_date DATE NOT NULL,
    total_revenue DECIMAL(18, 2) NOT NULL,
    total_expense DECIMAL(18, 2) NOT NULL,
    net_profit DECIMAL(18, 2) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT now()
);

-- ===========================
-- Table: employee_performance_reports
-- ===========================
CREATE TABLE IF NOT EXISTS employee_performance_reports (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4 (),
    employee_user_id UUID NOT NULL REFERENCES employees (user_id),
    report_date DATE NOT NULL,
    performance_score DECIMAL(5, 2) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT now()
);

-- ===========================
-- Table: customer_activity_reports
-- ===========================
CREATE TABLE IF NOT EXISTS customer_activity_reports (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4 (),
    customer_user_id UUID NOT NULL REFERENCES customers (user_id),
    company_id UUID NOT NULL REFERENCES companies (id),
    activity_date DATE NOT NULL,
    total_transactions INTEGER DEFAULT 0,
    total_amount DECIMAL(15, 2) DEFAULT 0,
    points_earned INTEGER DEFAULT 0,
    last_transaction_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ DEFAULT now()
);

-- ===========================
-- Table: transaction_summaries
-- ===========================
CREATE TABLE IF NOT EXISTS transaction_summaries (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4 (),
    store_id UUID NOT NULL REFERENCES stores (id),
    total_transactions INTEGER NOT NULL,
    total_revenue DECIMAL(15, 2) NOT NULL,
    period_start DATE NOT NULL,
    period_end DATE NOT NULL,
    created_at TIMESTAMPTZ DEFAULT now()
);