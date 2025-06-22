# /Users/andre/Programing/aplikasi perusahaan/pos-backend/tests/services/readme.md

Direktori ini berisi **unit test untuk lapisan Service** (`internal/core/services`). Test di sini berfokus pada validasi logika bisnis inti aplikasi, orkestrasi panggilan ke satu atau lebih repository, dan penanganan kasus penggunaan (use cases).

Tujuan utama folder ini adalah untuk:

- **Menguji Logika Bisnis**: Memastikan bahwa aturan dan proses bisnis diimplementasikan dengan benar.
- **Memverifikasi Orkestrasi**: Memastikan bahwa service berinteraksi dengan repository (atau service lain) secara benar untuk menyelesaikan tugas yang kompleks.
- **Mengisolasi Pengujian**: Menggunakan mock untuk dependensi (seperti repository) agar service dapat diuji secara terisolasi dari infrastruktur.
- **Mencegah Regresi**: Memastikan bahwa perubahan pada service tidak merusak fungsionalitas yang sudah ada.

---

## Daftar Service yang Akan Diuji

Berikut adalah daftar service yang akan memiliki unit test di direktori ini, beserta status pengujiannya:

- `AuthService`:
  - **Status: ✅ SELESAI & TERUJI**
- `UserService`:
  - **Status: ✅ SELESAI & TERUJI**
- `EmployeeService`:
  - **Status: ⬜ TO-DO**
- `CustomerService`:
  - **Status: ⬜ TO-DO**
- `RoleService`:
  - **Status: ⬜ TO-DO**
- `PermissionService`:
  - **Status: ⬜ TO-DO**
- `CompanyService`:
  - **Status: ⬜ TO-DO**
- `BusinessLineService`:
  - **Status: ⬜ TO-DO**
- `StoreService`:
  - **Status: ⬜ TO-DO**
- `MasterProductService`:
  - **Status: ⬜ TO-DO**
- `StoreProductService`:
  - **Status: ⬜ TO-DO**
- `SupplierService`:
  - **Status: ⬜ TO-DO**
- `TaxRateService`:
  - **Status: ⬜ TO-DO**
- `TransactionService`:
  - **Status: ⬜ TO-DO**
- `ShiftService`:
  - **Status: ⬜ TO-DO**
- `PurchaseOrderService`:
  - **Status: ⬜ TO-DO**
- `InternalStockTransferService`:
  - **Status: ⬜ TO-DO**
- `ActivityLogService`:
  - **Status: ⬜ TO-DO**
- `DiscountService`:
  - **Status: ⬜ TO-DO**
- `OperationalExpenseService`:
  - **Status: ⬜ TO-DO**

---

## Catatan

Unit test service sangat penting untuk memastikan kebenaran logika bisnis tanpa perlu bergantung pada database atau komponen eksternal lainnya. Mocking digunakan secara ekstensif di lapisan ini untuk mengisolasi kode yang sedang diuji.
