### Entitas Relasi Database (ERD)

Berikut adalah Entity Relationship Diagram (ERD) yang menggambarkan struktur database aplikasi ini:

// --- Core Tenant & Organization ---
Table companies {
id UUID [pk, default: `uuid_generate_v4()`]
name VARCHAR(255) [not null]
address TEXT
contact_info JSONB
tax_id_number VARCHAR(100)
default_tax_percentage DECIMAL(5,2) [note: "Fallback if granular tax_rates not used"]
created_at TIMESTAMPTZ [default: `now()`]
updated_at TIMESTAMPTZ [default: `now()`]
}

Table business_lines {
id UUID [pk, default: `uuid_generate_v4()`]
company_id UUID [not null, ref: > companies.id]
name VARCHAR(255) [not null]
description TEXT
created_at TIMESTAMPTZ [default: `now()`]
updated_at TIMESTAMPTZ [default: `now()`]
}

Table stores {
id UUID [pk, default: `uuid_generate_v4()`]
business_line_id UUID [not null, ref: > business_lines.id]
parent_store_id UUID [ref: - stores.id]
name VARCHAR(255) [not null]
store_code VARCHAR(50)
store_type VARCHAR(50) [not null, note: "CHECK (store_type IN ('PUSAT', 'CABANG', 'RANTING'))"]
address TEXT
phone_number VARCHAR(50)
created_at TIMESTAMPTZ [default: `now()`]
updated_at TIMESTAMPTZ [default: `now()`]
}

// --- User Management ---
Table users {
id UUID [pk, default: `uuid_generate_v4()`]
user_type VARCHAR(50) [not null, note: "CHECK (user_type IN ('EMPLOYEE', 'CUSTOMER'))"]
username VARCHAR(100) [unique]
password_hash VARCHAR(255)
full_name VARCHAR(255)
email VARCHAR(255) [unique]
phone_number VARCHAR(50) [unique]
is_active BOOLEAN [default: true]
created_at TIMESTAMPTZ [default: `now()`]
updated_at TIMESTAMPTZ [default: `now()`]
}

Table employees {
user_id UUID [pk, ref: > users.id]
company_id UUID [not null, ref: > companies.id]
store_id UUID [ref: - stores.id]
employee_id_number VARCHAR(100)
join_date DATE
position VARCHAR(100)
created_at TIMESTAMPTZ [default: `now()`]
updated_at TIMESTAMPTZ [default: `now()`]
}

Table roles {
id SERIAL [pk]
name VARCHAR(100) [unique, not null]
description TEXT
}

Table employee_roles {
employee_user_id UUID [pk, ref: > employees.user_id]
role_id INTEGER [pk, ref: > roles.id]
}

// --- Access Rights Management ---
Table permissions {
id SERIAL [pk]
name VARCHAR(255) [unique, not null, note: "e.g., 'product:create', 'transaction:view_all'"]
description TEXT
group_name VARCHAR(100) [note: "For UI grouping, e.g., 'Product Management'"]
}

Table role_permissions {
role_id INTEGER [pk, ref: > roles.id]
permission_id INTEGER [pk, ref: > permissions.id]
}

Table customers {
user_id UUID [pk, ref: > users.id]
company_id UUID [not null, ref: > companies.id]
membership_number VARCHAR(100)
join_date DATE
points INTEGER [default: 0]
tier VARCHAR(50)
created_at TIMESTAMPTZ [default: `now()`]
updated_at TIMESTAMPTZ [default: `now()`]
Indexes {
(company_id, membership_number) [unique]
}
}

// --- Shift Management ---
Table shifts {
id UUID [pk, default: `uuid_generate_v4()`]
employee_user_id UUID [not null, ref: > employees.user_id]
store_id UUID [not null, ref: > stores.id]
shift_date DATE [not null]
start_time TIME [not null]
end_time TIME [not null]
actual_check_in TIMESTAMPTZ
actual_check_out TIMESTAMPTZ
notes TEXT
created_by_user_id UUID [ref: > users.id]
created_at TIMESTAMPTZ [default: `now()`]
updated_at TIMESTAMPTZ [default: `now()`]
}

// --- Supplier Management ---
Table suppliers {
id UUID [pk, default: `uuid_generate_v4()`]
company_id UUID [not null, ref: > companies.id]
name VARCHAR(255) [not null]
contact_person VARCHAR(255)
email VARCHAR(255)
phone_number VARCHAR(50)
address TEXT
created_at TIMESTAMPTZ [default: `now()`]
updated_at TIMESTAMPTZ [default: `now()`]
}

// --- Tax Management ---
Table tax_rates {
id SERIAL [pk]
company_id UUID [not null, ref: > companies.id]
name VARCHAR(100) [not null, note: "e.g., PPN 11%, Service Charge 5%"]
rate_percentage DECIMAL(5,2) [not null]
description TEXT
is_active BOOLEAN [default: true]
created_at TIMESTAMPTZ [default: `now()`]
updated_at TIMESTAMPTZ [default: `now()`]
}

// --- Centralized Product Definition ---
Table master_products {
id UUID [pk, default: `uuid_generate_v4()`]
company_id UUID [not null, ref: > companies.id]
master_product_code VARCHAR(100) [not null]
name VARCHAR(255) [not null]
description TEXT
category VARCHAR(100)
unit_of_measure VARCHAR(50)
barcode VARCHAR(255)
default_tax_rate_id INTEGER [ref: - tax_rates.id]
image_url VARCHAR(255) // Added image_url column
created_at TIMESTAMPTZ [default: `now()`]
updated_at TIMESTAMPTZ [default: `now()`]
Indexes {
(company_id, master_product_code) [unique]
}
}

Table store_products {
id UUID [pk, default: `uuid_generate_v4()`]
master_product_id UUID [not null, ref: > master_products.id]
store_id UUID [not null, ref: > stores.id]
supplier_id UUID [ref: - suppliers.id]
store_specific_sku VARCHAR(100)
purchase_price DECIMAL(15,2) [not null]
selling_price DECIMAL(15,2) [not null]
wholesale_price DECIMAL(15,2)
stock INTEGER [not null, default: 0]
minimum_stock_level INTEGER
expiry_date DATE
created_at TIMESTAMPTZ [default: `now()`]
updated_at TIMESTAMPTZ [default: `now()`]
Indexes {
(store_id, master_product_id) [unique]
(store_id, store_specific_sku) [unique]
}
}

// --- Transaction Management ---
Table transactions {
id UUID [pk, default: `uuid_generate_v4()`]
transaction_code VARCHAR(255) [unique, not null]
store_id UUID [not null, ref: > stores.id]
cashier_employee_user_id UUID [not null, ref: > employees.user_id]
customer_user_id UUID [ref: - customers.user_id]
active_shift_id UUID [ref: - shifts.id]
transaction_date TIMESTAMPTZ [not null, default: `now()`]
subtotal_amount DECIMAL(15,2) [not null]
total_item_discount_amount DECIMAL(15,2) [default: 0]
subtotal_after_item_discounts DECIMAL(15,2) [not null]
transaction_level_discount_amount DECIMAL(15,2) [default: 0]
taxable_amount DECIMAL(15,2) [not null]
total_tax_amount DECIMAL(15,2) [default: 0]
final_total_amount DECIMAL(15,2) [not null]
received_amount DECIMAL(15,2) [not null]
change_amount DECIMAL(15,2) [not null]
payment_method VARCHAR(50)
notes TEXT
created_at TIMESTAMPTZ [default: `now()`]
updated_at TIMESTAMPTZ [default: `now()`]
}

Table transaction_items {
id UUID [pk, default: `uuid_generate_v4()`]
transaction_id UUID [not null, ref: > transactions.id]
store_product_id UUID [not null, ref: > store_products.id]
quantity INTEGER [not null]
price_per_unit_at_transaction DECIMAL(15,2) [not null]
item_subtotal_before_discount DECIMAL(15,2) [not null]
item_discount_amount DECIMAL(15,2) [default: 0]
item_subtotal_after_discount DECIMAL(15,2) [not null]
applied_tax_rate_id INTEGER [ref: - tax_rates.id]
applied_tax_rate_percentage DECIMAL(5,2)
tax_amount_for_item DECIMAL(15,2) [default: 0]
item_final_total DECIMAL(15,2) [not null]
created_at TIMESTAMPTZ [default: `now()`]
updated_at TIMESTAMPTZ [default: `now()`]
}

// --- Discount Management ---
Table discounts {
id UUID [pk, default: `uuid_generate_v4()`]
company_id UUID [not null, ref: > companies.id]
name VARCHAR(255) [not null]
description TEXT
discount_type VARCHAR(50) [not null, note: "CHECK (discount_type IN ('PERCENTAGE', 'FIXED_AMOUNT'))"]
discount_value DECIMAL(15,2) [not null]
applicable_to VARCHAR(50) [not null, note: "CHECK (applicable_to IN ('MASTER_PRODUCT', 'STORE_PRODUCT', 'CATEGORY', 'TOTAL_TRANSACTION', 'CUSTOMER_TIER'))"]
master_product_id_applicable UUID [ref: - master_products.id]
store_product_id_applicable UUID [ref: - store_products.id]
category_applicable VARCHAR(100)
customer_tier_applicable VARCHAR(50)
min_purchase_amount DECIMAL(15,2)
start_date DATE [not null]
end_date DATE [not null]
is_active BOOLEAN [default: true]
created_at TIMESTAMPTZ [default: `now()`]
updated_at TIMESTAMPTZ [default: `now()`]
}

Table applied_item_discounts {
transaction_item_id UUID [pk, ref: > transaction_items.id]
discount_id UUID [pk, ref: > discounts.id]
applied_discount_amount_on_item DECIMAL(15,2) [not null]
}

Table applied_transaction_discounts {
transaction_id UUID [pk, ref: > transactions.id]
discount_id UUID [pk, ref: > discounts.id]
applied_discount_amount_on_transaction DECIMAL(15,2) [not null]
}

// --- Purchasing & Stock Management ---
Table purchase_orders {
id UUID [pk, default: `uuid_generate_v4()`]
store_id UUID [not null, ref: > stores.id]
supplier_id UUID [not null, ref: > suppliers.id]
order_date DATE [not null]
expected_delivery_date DATE
status VARCHAR(50) [not null, note: "CHECK (status IN ('PENDING', 'ORDERED', 'PARTIALLY_RECEIVED', 'RECEIVED', 'CANCELLED'))"]
total_amount DECIMAL(15,2)
notes TEXT
created_by_user_id UUID [not null, ref: > users.id]
created_at TIMESTAMPTZ [default: `now()`]
updated_at TIMESTAMPTZ [default: `now()`]
}

Table purchase_order_items {
id UUID [pk, default: `uuid_generate_v4()`]
purchase_order_id UUID [not null, ref: > purchase_orders.id]
master_product_id UUID [not null, ref: > master_products.id]
quantity_ordered INTEGER [not null]
purchase_price_per_unit DECIMAL(15,2) [not null]
quantity_received INTEGER [default: 0]
subtotal DECIMAL(15,2) [not null]
created_at TIMESTAMPTZ [default: `now()`]
updated_at TIMESTAMPTZ [default: `now()`]
}

Table internal_stock_transfers {
id UUID [pk, default: `uuid_generate_v4()`]
transfer_code VARCHAR(100) [unique, not null]
company_id UUID [not null, ref: > companies.id]
source_store_id UUID [not null, ref: > stores.id]
destination_store_id UUID [not null, ref: > stores.id]
transfer_date DATE [not null]
status VARCHAR(50) [not null, note: "CHECK (status IN ('PENDING', 'APPROVED', 'SHIPPED', 'PARTIALLY_RECEIVED', 'RECEIVED', 'CANCELLED'))"]
notes TEXT
requested_by_user_id UUID [ref: - users.id]
approved_by_user_id UUID [ref: - users.id]
shipped_by_user_id UUID [ref: - users.id]
received_by_user_id UUID [ref: - users.id]
created_at TIMESTAMPTZ [default: `now()`]
updated_at TIMESTAMPTZ [default: `now()`]
}

Table internal_stock_transfer_items {
id UUID [pk, default: `uuid_generate_v4()`]
internal_stock_transfer_id UUID [not null, ref: > internal_stock_transfers.id]
source_store_product_id UUID [not null, ref: > store_products.id]
quantity_requested INTEGER [not null]
quantity_shipped INTEGER [default: 0]
quantity_received INTEGER [default: 0]
notes TEXT
created_at TIMESTAMPTZ [default: `now()`]
updated_at TIMESTAMPTZ [default: `now()`]
}

Table stock_movements {
id UUID [pk, default: `uuid_generate_v4()`]
store_product_id UUID [not null, ref: > store_products.id]
store_id UUID [not null, ref: > stores.id]
movement_type VARCHAR(50) [not null, note: "CHECK (movement_type IN ('SALE', 'PURCHASE_RECEIPT', 'RETURN_TO_SUPPLIER', 'CUSTOMER_RETURN', 'ADJUSTMENT_IN', 'ADJUSTMENT_OUT', 'TRANSFER_OUT', 'TRANSFER_IN'))"]
quantity_changed INTEGER [not null]
movement_date TIMESTAMPTZ [not null, default: `now()`]
reference_id UUID
reference_type VARCHAR(100)
notes TEXT
created_by_user_id UUID [ref: - users.id]
created_at TIMESTAMPTZ [default: `now()`]
}

// --- Auditing ---
Table activity_logs {
id BIGSERIAL [pk]
user_id UUID [ref: - users.id]
company_id UUID [ref: - companies.id]
store_id UUID [ref: - stores.id]
action_type VARCHAR(255) [not null]
description TEXT [not null]
target_entity VARCHAR(100)
target_entity_id VARCHAR(255)
ip_address VARCHAR(100)
user_agent TEXT
log_time TIMESTAMPTZ [not null, default: `now()`]
}

// --- Financials & Expenses ---
Table operational_expenses {
id UUID [pk, default: `uuid_generate_v4()`]
company_id UUID [not null, ref: > companies.id]
store_id UUID [ref: - stores.id] // Biaya bisa terkait toko spesifik atau perusahaan secara umum (store_id NULL)
expense_date DATE [not null]
category VARCHAR(100) [not null, note: "e.g., 'Gaji Karyawan', 'Sewa Toko', 'Listrik & Air', 'Pemasaran', 'Perlengkapan Kantor', 'Lain-lain'"]
description TEXT
amount DECIMAL(15,2) [not null]
created_by_user_id UUID [ref: > users.id] // Siapa yang mencatat biaya ini
created_at TIMESTAMPTZ [default: `now()`]
updated_at TIMESTAMPTZ [default: `now()`]
}
