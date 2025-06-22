# /Users/andre/Programing/aplikasi perusahaan/pos-backend/tests/repository/readme.md

Direktori ini berisi **unit test untuk lapisan Repository** (`internal/data`). Test di sini berfokus pada validasi interaksi langsung dengan database.

Tujuan utama folder ini adalah untuk:

- **Memverifikasi Query SQL**: Memastikan bahwa semua query SQL (INSERT, SELECT, UPDATE, DELETE) yang ditulis dalam repository berfungsi dengan benar.
- **Menjamin Integritas Data**: Menguji bagaimana repository menangani operasi data, termasuk constraint dan relasi antar tabel.
- **Mengisolasi Pengujian Data**: Menguji lapisan data secara terpisah dari logika bisnis untuk memastikan fondasi aplikasi kokoh.

---

## Struktur Subdirektori

Direktori ini diatur ke dalam subdirektori berdasarkan jenis database yang diuji:

- `postgres/`: Berisi unit test yang spesifik untuk implementasi repository PostgreSQL.
  - **Status: ✅ SELESAI & TERUJI**

---

## Daftar Repository yang Diuji

Berikut adalah daftar repository yang memiliki unit test di direktori ini, beserta status pengujiannya:

- `ActivityLogRepository`:
  - **Status: ✅ SELESAI & TERUJI**
- `AppliedItemDiscountRepository`:
  - **Status: ✅ SELESAI & TERUJI**
- `AppliedTransactionDiscountRepository`:
  - **Status: ✅ SELESAI & TERUJI**
- `BusinessLineRepository`:
  - **Status: ✅ SELESAI & TERUJI**
- `CompanyRepository`:
  - **Status: ✅ SELESAI & TERUJI**
- `CustomerRepository`:
  - **Status: ✅ SELESAI & TERUJI**
- `DiscountRepository`:
  - **Status: ✅ SELESAI & TERUJI**
- `EmployeeRepository`:
  - **Status: ✅ SELESAI & TERUJI**
- `EmployeeRoleRepository`:
  - **Status: ✅ SELESAI & TERUJI**
- `InternalStockTransferRepository`:
  - **Status: ✅ SELESAI & TERUJI**
- `MasterProductRepository`:
  - **Status: ✅ SELESAI & TERUJI**
- `OperationalExpenseRepository`:
  - **Status: ✅ SELESAI & TERUJI**
- `PermissionRepository`:
  - **Status: ✅ SELESAI & TERUJI**
- `PurchaseOrderRepository`:
  - **Status: ✅ SELESAI & TERUJI**
- `RolePermissionRepository`:
  - **Status: ✅ SELESAI & TERUJI**
- `RoleRepository`:
  - **Status: ✅ SELESAI & TERUJI**
- `ShiftRepository`:
  - **Status: ✅ SELESAI & TERUJI**
- `StockMovementRepository`:
  - **Status: ✅ SELESAI & TERUJI**
- `StoreProductRepository`:
  - **Status: ✅ SELESAI & TERUJI**
- `StoreRepository`:
  - **Status: ✅ SELESAI & TERUJI**
- `SupplierRepository`:
  - **Status: ✅ SELESAI & TERUJI**
- `TaxRateRepository`:
  - **Status: ✅ SELESAI & TERUJI**
- `TransactionRepository`:
  - **Status: ✅ SELESAI & TERUJI**
- `UserRepository`:
  - **Status: ✅ SELESAI & TERUJI**

---

## Catatan

Test di dalam direktori ini biasanya memerlukan koneksi ke database tes yang berjalan untuk memvalidasi operasi secara nyata. Mekanisme setup dan teardown (seperti `TestMain` dan fungsi `cleanup()`) sangat penting untuk memastikan setiap test berjalan dalam keadaan yang bersih dan terisolasi.
