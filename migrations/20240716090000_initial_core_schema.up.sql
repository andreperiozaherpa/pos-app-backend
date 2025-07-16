-- 20240716090000_initial_core_schema.sql
-- Migrasi initial core schema (tabel inti: perusahaan, organisasi, user, role, permission, customer)

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
-- 1. companies
CREATE TABLE IF NOT EXISTS companies (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4 (),
    name VARCHAR(255) NOT NULL,
    address TEXT,
    contact_info JSONB,
    tax_id_number VARCHAR(100),
    default_tax_percentage DECIMAL(5, 2),
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now()
);

-- 2. business_lines
CREATE TABLE IF NOT EXISTS business_lines (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4 (),
    company_id UUID NOT NULL REFERENCES companies (id),
    name VARCHAR(255) NOT NULL,
    description TEXT,
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now()
);

-- 3. stores
CREATE TABLE IF NOT EXISTS stores (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4 (),
    business_line_id UUID NOT NULL REFERENCES business_lines (id),
    parent_store_id UUID REFERENCES stores (id),
    name VARCHAR(255) NOT NULL,
    store_code VARCHAR(50),
    store_type VARCHAR(50) NOT NULL,
    address TEXT,
    phone_number VARCHAR(50),
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now(),
    CONSTRAINT store_type_check CHECK (
        store_type IN ('PUSAT', 'CABANG', 'RANTING')
    )
);

-- 4. users
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4 (),
    user_type VARCHAR(50) NOT NULL,
    username VARCHAR(100) UNIQUE,
    password_hash VARCHAR(255),
    full_name VARCHAR(255),
    email VARCHAR(255) UNIQUE,
    phone_number VARCHAR(50) UNIQUE,
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now(),
    CONSTRAINT user_type_check CHECK (
        user_type IN ('EMPLOYEE', 'CUSTOMER')
    )
);

-- 5. employees
CREATE TABLE IF NOT EXISTS employees (
    user_id UUID PRIMARY KEY REFERENCES users (id),
    company_id UUID NOT NULL REFERENCES companies (id),
    store_id UUID REFERENCES stores (id),
    employee_id_number VARCHAR(100),
    join_date DATE,
    position VARCHAR(100),
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now()
);

-- 6. roles
CREATE TABLE IF NOT EXISTS roles (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) UNIQUE NOT NULL,
    description TEXT
);

-- 7. employee_roles (pivot)
CREATE TABLE IF NOT EXISTS employee_roles (
    employee_user_id UUID REFERENCES employees (user_id) ON DELETE CASCADE,
    role_id INTEGER REFERENCES roles (id) ON DELETE CASCADE,
    PRIMARY KEY (employee_user_id, role_id)
);

-- 8. permissions
CREATE TABLE IF NOT EXISTS permissions (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) UNIQUE NOT NULL,
    description TEXT,
    group_name VARCHAR(100)
);

-- 9. role_permissions (pivot)
CREATE TABLE IF NOT EXISTS role_permissions (
    role_id INTEGER REFERENCES roles (id) ON DELETE CASCADE,
    permission_id INTEGER REFERENCES permissions (id) ON DELETE CASCADE,
    PRIMARY KEY (role_id, permission_id)
);

-- 10. customers
CREATE TABLE IF NOT EXISTS customers (
    user_id UUID PRIMARY KEY REFERENCES users (id),
    company_id UUID NOT NULL REFERENCES companies (id),
    membership_number VARCHAR(100),
    join_date DATE,
    points INTEGER DEFAULT 0,
    tier VARCHAR(50),
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now(),
    CONSTRAINT unique_company_membership_number UNIQUE (company_id, membership_number)
);