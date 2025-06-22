# /Users/andre/Programing/aplikasi perusahaan/pos-backend/tests/mocks/readme.md

Direktori ini berisi implementasi **mock dari interface repository** (dan mungkin interface service atau utilitas lainnya di masa depan). Mock digunakan dalam unit test untuk mengisolasi unit kode yang sedang diuji dari dependensi eksternal.

Tujuan utama folder ini adalah untuk:

- **Isolasi Pengujian**: Memungkinkan pengujian logika bisnis service secara terpisah dari implementasi database atau sistem eksternal lainnya.
- **Kontrol Perilaku Dependensi**: Memungkinkan kita untuk menentukan perilaku spesifik dari dependensi (misalnya, mengembalikan error tertentu, mengembalikan data tertentu) untuk menguji berbagai skenario.
- **Mempercepat Pengujian**: Menghilangkan kebutuhan akan koneksi database yang sebenarnya, membuat unit test berjalan lebih cepat.

---

## Daftar Mock

Berikut adalah daftar mock yang ada atau akan dibuat di direktori ini, beserta statusnya:

- `UserRepositoryMock`: Mock untuk `internal/data/postgres.UserRepository`.
  - **Status: ✅ SELESAI & TERUJI**
- `ActivityLogRepositoryMock`: Mock untuk `internal/data/postgres.ActivityLogRepository`.
  - **Status: ⬜ TO-DO**
- `AppliedItemDiscountRepositoryMock`: Mock untuk `internal/data/postgres.AppliedItemDiscountRepository`.
  - **Status: ⬜ TO-DO**
- `AppliedTransactionDiscountRepositoryMock`: Mock untuk `internal/data/postgres.AppliedTransactionDiscountRepository`.
  - **Status: ⬜ TO-DO**
- `BusinessLineRepositoryMock`: Mock untuk `internal/data/postgres.BusinessLineRepository`.
  - **Status: ⬜ TO-DO**
- `CompanyRepositoryMock`: Mock untuk `internal/data/postgres.CompanyRepository`.
  - **Status: ⬜ TO-DO**
- `CustomerRepositoryMock`: Mock untuk `internal/data/postgres.CustomerRepository`.
  - **Status: ✅ SELESAI & TERUJI**
- `DiscountRepositoryMock`: Mock untuk `internal/data/postgres.DiscountRepository`.
  - **Status: ⬜ TO-DO**
- `EmployeeRepositoryMock`: Mock untuk `internal/data/postgres.EmployeeRepository`.
  - **Status: ✅ SELESAI & TERUJI**
- `EmployeeRoleRepositoryMock`: Mock untuk `internal/data/postgres.EmployeeRoleRepository`.
  - **Status: ✅ SELESAI & TERUJI**
- `InternalStockTransferRepositoryMock`: Mock untuk `internal/data/postgres.InternalStockTransferRepository`.
  - **Status: ⬜ TO-DO**
- `MasterProductRepositoryMock`: Mock untuk `internal/data/postgres.MasterProductRepository`.
  - **Status: ⬜ TO-DO**
- `OperationalExpenseRepositoryMock`: Mock untuk `internal/data/postgres.OperationalExpenseRepository`.
  - **Status: ⬜ TO-DO**
- `PermissionRepositoryMock`: Mock untuk `internal/data/postgres.PermissionRepository`.
  - **Status: ⬜ TO-DO**
- `PurchaseOrderRepositoryMock`: Mock untuk `internal/data/postgres.PurchaseOrderRepository`.
  - **Status: ⬜ TO-DO**
- `RolePermissionRepositoryMock`: Mock untuk `internal/data/postgres.RolePermissionRepository`.
  - **Status: ⬜ TO-DO**
- `RoleRepositoryMock`: Mock untuk `internal/data/postgres.RoleRepository`.
  - **Status: ✅ SELESAI & TERUJI**
- `ShiftRepositoryMock`: Mock untuk `internal/data/postgres.ShiftRepository`.
  - **Status: ⬜ TO-DO**
- `StockMovementRepositoryMock`: Mock untuk `internal/data/postgres.StockMovementRepository`.
  - **Status: ⬜ TO-DO**
- `StoreProductRepositoryMock`: Mock untuk `internal/data/postgres.StoreProductRepository`.
  - **Status: ⬜ TO-DO**
- `StoreRepositoryMock`: Mock untuk `internal/data/postgres.StoreRepository`.
  - **Status: ⬜ TO-DO**
- `SupplierRepositoryMock`: Mock untuk `internal/data/postgres.SupplierRepository`.
  - **Status: ⬜ TO-DO**
- `TaxRateRepositoryMock`: Mock untuk `internal/data/postgres.TaxRateRepository`.
  - **Status: ⬜ TO-DO**
- `TransactionRepositoryMock`: Mock untuk `internal/data/postgres.TransactionRepository`.
  - **Status: ⬜ TO-DO**

---

## Catatan

Mocking adalah teknik penting dalam unit testing untuk memastikan bahwa test hanya menguji unit kode yang dimaksud, tanpa dipengaruhi oleh perilaku dependensi eksternal. Mock ini akan digunakan secara luas di lapisan `tests/services` dan `tests/api`.
