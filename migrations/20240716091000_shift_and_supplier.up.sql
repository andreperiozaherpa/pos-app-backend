-- 20240716091000_shift_and_supplier.sql
-- Migrasi tabel: shifts, shift_attendances, shift_swaps, suppliers

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
-- 1. suppliers
CREATE TABLE IF NOT EXISTS suppliers (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4 (),
    company_id UUID NOT NULL REFERENCES companies (id),
    name VARCHAR(255) NOT NULL,
    contact_person VARCHAR(255),
    email VARCHAR(255),
    phone_number VARCHAR(50),
    address TEXT,
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now()
);

-- 2. shifts
CREATE TABLE IF NOT EXISTS shifts (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4 (),
    employee_user_id UUID NOT NULL REFERENCES employees (user_id),
    store_id UUID NOT NULL REFERENCES stores (id),
    shift_date DATE NOT NULL,
    start_time TIME NOT NULL,
    end_time TIME NOT NULL,
    actual_check_in TIMESTAMPTZ,
    actual_check_out TIMESTAMPTZ,
    notes TEXT,
    created_by_user_id UUID REFERENCES users (id),
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now()
);

-- 3. shift_attendances
CREATE TABLE IF NOT EXISTS shift_attendances (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4 (),
    shift_id UUID NOT NULL REFERENCES shifts (id),
    employee_id UUID NOT NULL REFERENCES employees (user_id),
    check_in TIMESTAMPTZ,
    check_out TIMESTAMPTZ,
    created_at TIMESTAMPTZ DEFAULT now()
);

-- 4. shift_swaps
CREATE TABLE IF NOT EXISTS shift_swaps (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4 (),
    requested_by_employee_user_id UUID NOT NULL REFERENCES employees (user_id),
    requested_to_employee_user_id UUID NOT NULL REFERENCES employees (user_id),
    shift_id UUID NOT NULL REFERENCES shifts (id),
    requested_shift_date DATE NOT NULL,
    reason TEXT,
    status VARCHAR(50) NOT NULL,
    approved_by_user_id UUID REFERENCES users (id),
    approved_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now(),
    CONSTRAINT shift_swap_status_check CHECK (
        status IN (
            'PENDING',
            'APPROVED',
            'REJECTED',
            'CANCELLED'
        )
    )
);