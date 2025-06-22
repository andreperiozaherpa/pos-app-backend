# /internal/models

Direktori ini berisi definisi struct Go yang merepresentasikan entitas data dalam aplikasi. Setiap file `.go` di sini biasanya mewakili satu tabel dalam database.

Struct ini digunakan di seluruh aplikasi, dari lapisan data (repository) untuk memetakan hasil query database, hingga lapisan service untuk logika bisnis, dan lapisan API untuk transfer data (DTOs - Data Transfer Objects).

---

## Tinjauan Kebutuhan Model

Berikut adalah daftar model yang dibutuhkan berdasarkan ERD proyek, beserta status pengerjaannya.

### 1. Modul: core_tenant_and_organization

- `company`: Merepresentasikan tabel `companies`.
  - **Status: ✅ SELESAI**
- `business_line`: Merepresentasikan tabel `business_lines`.
  - **Status: ✅ SELESAI**
- `store`: Merepresentasikan tabel `stores`.
  - **Status: ✅ SELESAI**

### 2. Modul: user_management

- `user`: Merepresentasikan tabel `users`.
  - **Status: ✅ SELESAI**
- `employee`: Merepresentasikan tabel `employees`.
  - **Status: ✅ SELESAI**
- `customer`: Merepresentasikan tabel `customers`.
  - **Status: ✅ SELESAI**
- `role`: Merepresentasikan tabel `roles`.
  - **Status: ✅ SELESAI**
- `permission`: Merepresentasikan tabel `permissions`.
  - **Status: ✅ SELESAI**
- `employee_role`: Merepresentasikan tabel pivot `employee_roles`.
  - **Status: ✅ SELESAI**

### 3. Modul: product_and_supplier_management

- `supplier`: Merepresentasikan tabel `suppliers`.
  - **Status: ✅ SELESAI**
- `master_product`: Merepresentasikan tabel `master_products`.
  - **Status: ✅ SELESAI**
- `store_product`: Merepresentasikan tabel `store_products`. (File: `product.go`, Struct: `StoreProduct`)
  - **Status: ✅ SELESAI**
- `tax_rate`: Merepresentasikan tabel `tax_rates`.
  - **Status: ✅ SELESAI**

### 4. Modul: transaction_management

- `transaction`: Merepresentasikan tabel `transactions`.
  - **Status: ✅ SELESAI**
- `transaction_item`: Merepresentasikan tabel `transaction_items`. (Didefinisikan dalam `transaction.go`)
  - **Status: ✅ SELESAI**

### 5. Modul: discount_management

- `discount`: Merepresentasikan tabel `discounts`.
  - **Status: ✅ SELESAI**
- `applied_item_discount`: Merepresentasikan tabel `applied_item_discounts`. (File: `applied_discount.go`)
  - **Status: ✅ SELESAI**
- `applied_transaction_discount`: Merepresentasikan tabel `applied_transaction_discounts`. (File: `applied_discount.go`)
  - **Status: ✅ SELESAI**

### 6. Modul: shift_management

- `shift`: Merepresentasikan tabel `shifts`.
  - **Status: ✅ SELESAI**

### 7. Modul: purchasing_and_stock_management

- `purchase_order`: Merepresentasikan tabel `purchase_orders`. (File: `purchase_order.go`)
  - **Status: ✅ SELESAI**
- `purchase_order_item`: Merepresentasikan tabel `purchase_order_items`. (Didefinisikan dalam `purchase_order.go`)
  - **Status: ✅ SELESAI**
- `internal_stock_transfer`: Merepresentasikan tabel `internal_stock_transfers`. (File: `stock_transfer.go`)
  - **Status: ✅ SELESAI**
- `internal_stock_transfer_item`: Merepresentasikan tabel `internal_stock_transfer_items`. (Didefinisikan dalam `stock_transfer.go`)
  - **Status: ✅ SELESAI**
- `stock_movement`: Merepresentasikan tabel `stock_movements`. (File: `stock_movement.go`)
  - **Status: ✅ SELESAI**

### 8. Modul: auditing

- `activity_log`: Merepresentasikan tabel `activity_logs`.
  - **Status: ✅ SELESAI**

### 9. Modul: expense_management

- `operational_expense`: Merepresentasikan tabel `operational_expenses`. (File: `operational_expense.go`)
  - **Status: ✅ SELESAI**
