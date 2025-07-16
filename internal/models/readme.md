# /internal/models

Direktori ini berisi definisi struct Go yang merepresentasikan entitas data utama aplikasi.  
Setiap file `.go` di sini umumnya mewakili satu tabel pada database utama sesuai dengan ERD aplikasi.

Struct-struct ini digunakan di seluruh aplikasi, mulai dari lapisan data (repository untuk mapping hasil query), logika bisnis (service), hingga API (sebagai DTO: Data Transfer Objects).

---

## 📦 Daftar Model & Status

| Modul                                  | Model / Struct               | File                                         | Status     |
| -------------------------------------- | ---------------------------- | -------------------------------------------- | ---------- |
| **1. Tenant & Organisasi**             | `Company`                    | company.go                                   | ✅ Selesai |
|                                        | `BusinessLine`               | business_line.go                             | ✅ Selesai |
|                                        | `Store`                      | store.go                                     | ✅ Selesai |
| **2. User & Role Management**          | `User`                       | user.go                                      | ✅ Selesai |
|                                        | `Employee`                   | employee.go                                  | ✅ Selesai |
|                                        | `Customer`                   | customer.go                                  | ✅ Selesai |
|                                        | `Role`                       | role.go                                      | ✅ Selesai |
|                                        | `Permission`                 | permission.go                                | ✅ Selesai |
|                                        | `EmployeeRole` (Pivot)       | employee_role.go                             | ✅ Selesai |
| **3. Produk & Supplier**               | `Supplier`                   | supplier.go                                  | ✅ Selesai |
|                                        | `MasterProduct`              | master_product.go                            | ✅ Selesai |
|                                        | `StoreProduct`               | product.go                                   | ✅ Selesai |
|                                        | `TaxRate`                    | tax_rate.go                                  | ✅ Selesai |
| **4. Transaksi**                       | `Transaction`                | transaction.go                               | ✅ Selesai |
|                                        | `TransactionItem`            | transaction.go                               | ✅ Selesai |
| **5. Diskon & Benefit**                | `Discount`                   | discount.go                                  | ✅ Selesai |
|                                        | `AppliedItemDiscount`        | applied_discount.go                          | ✅ Selesai |
|                                        | `AppliedTransactionDiscount` | applied_discount.go                          | ✅ Selesai |
| **6. Shift Management**                | `Shift`                      | shift.go                                     | ✅ Selesai |
| **7. Purchasing & Stok**               | `PurchaseOrder`              | purchase_order.go                            | ✅ Selesai |
|                                        | `PurchaseOrderItem`          | purchase_order.go                            | ✅ Selesai |
|                                        | `InternalStockTransfer`      | stock_transfer.go                            | ✅ Selesai |
|                                        | `InternalStockTransferItem`  | stock_transfer.go                            | ✅ Selesai |
|                                        | `StockMovement`              | stock_movement.go                            | ✅ Selesai |
| **8. Audit & Activity Log**            | `ActivityLog`                | activity_log.go                              | ✅ Selesai |
| **9. Pengeluaran Operasional**         | `OperationalExpense`         | operational_expense.go                       | ✅ Selesai |
| **10. Advance/Supporting (Opsional)**  | `EmployeeAttendance`         | employee_attendance.go (sudah ada struct)    | ✅ Selesai |
|                                        | `EmployeeLeave`              | employee_leave.go                            | ✅ Selesai |
|                                        | `EmployeePerformanceSummary` | employee_performance.go (sudah ada struct)   | ✅ Selesai |
|                                        | `ContactHistory`             | contact_history.go (sudah ada struct)        | ✅ Selesai |
|                                        | `MasterProductVariant`       | master_product_variant.go (sudah ada struct) | ✅ Selesai |
|                                        | `TransactionAuditTrail`      | transaction_audit.go (sudah ada struct)      | ✅ Selesai |
|                                        | `TransactionRefund`          | transaction_refund.go (sudah ada struct)     | ✅ Selesai |
|                                        | `TransactionReceipt`         | transaction_receipt.go                       | ✅ Selesai |
|                                        | `ExpenseReport`              | expense_report.go                            | ✅ Selesai |
|                                        | `Notification`               | notification.go                              | ✅ Selesai |
|                                        | `AuthSession`                | auth_session.go                              | ✅ Selesai |
|                                        | `RBACAssignment`             | rbac_assignment.go                           | ✅ Selesai |
|                                        | `ImportHistory`              | import_history.go                            | ✅ Selesai |
|                                        | `ExportHistory`              | export_history.go                            | ✅ Selesai |
|                                        | `FileMetadata`               | file_metadata.go                             | ✅ Selesai |
| **11. Financial Summary (Baru)**       | `CompanyFinancialSummary`    | company_financial_summary.go                 | ✅ Selesai |
| **12. Advance Product History (Baru)** | `MasterProductHistory`       | master_product_history.go                    | ⬜ Belum   |
|                                        | `PurchaseOrderHistory`       | purchase_order_history.go                    | ✅ Selesai |
| **13. Laporan & Statistik (Baru)**     | `SalesReport`                | sales_report.go                              | ✅ Selesai |
|                                        | `StockReport`                | stock_report.go                              | ✅ Selesai |
|                                        | `ProfitLossReport`           | profit_loss_report.go                        | ✅ Selesai |
|                                        | `EmployeePerformanceReport`  | employee_performance_report.go               | ✅ Selesai |
|                                        | `CustomerActivityReport`     | customer_activity_report.go                  | ✅ Selesai |
|                                        | `ShiftAttendance`            | shift_attendance.go                          | ✅ Selesai |
|                                        | `ShiftSwap`                  | shift_swap.go                                | ✅ Selesai |
|                                        | `StockMovementSummary`       | stock_movement_summary.go                    | ✅ Selesai |
|                                        | `StockTransfer`              | stock_transfer.go                            | ✅ Selesai |
|                                        | `StockTransferHistory`       | stock_transfer_history.go                    | ✅ Selesai |
|                                        | `StoreProductStockUpdate`    | store_product_stock_update.go                | ✅ Selesai |
|                                        | `TransactionAuditLog`        | transaction_audit_log.go                     | ✅ Selesai |
|                                        | `PaymentInfo`                | payment_info.go                              | ✅ Selesai |
|                                        | `TransactionSummary`         | transaction_summary.go                       | ✅ Selesai |
|                                        | `UserLoginHistory`           | user_login_history.go                        | ✅ Selesai |

---

## 📚 Penjelasan Modul & Relasi

### 1. **Tenant & Organisasi**

- **Company, BusinessLine, Store:**  
  Struktur multi-perusahaan, multi-lini usaha, hingga toko/cabang/ranting (mendukung skema Pusat → Cabang → Ranting).

### 2. **User & Role Management**

- **User, Employee, Customer, Role, Permission, EmployeeRole:**  
  Mendukung multi-role, manajemen karyawan & member, serta otorisasi (RBAC, pivot table untuk multi-role).

### 3. **Produk & Supplier**

- **Supplier, MasterProduct, StoreProduct, TaxRate:**  
  Manajemen barang, supplier, stok per toko, harga beli/jual/grosir, dan granular tax.

### 4. **Transaksi**

- **Transaction, TransactionItem:**  
  Penjualan kasir (POS), dengan kode transaksi unik, tracking kasir, customer, shift, dan diskon.

### 5. **Diskon & Benefit**

- **Discount, AppliedItemDiscount, AppliedTransactionDiscount:**  
  Sistem diskon per item, transaksi, tier member, serta program promosi.

### 6. **Shift Management**

- **Shift:**  
  Penjadwalan dan absensi karyawan per toko/shift.

### 7. **Purchasing & Stok**

- **PurchaseOrder, PurchaseOrderItem, InternalStockTransfer, InternalStockTransferItem, StockMovement:**  
  Proses pembelian barang ke supplier, transfer antar toko, dan mutasi stok terintegrasi.

### 8. **Audit & Activity Log**

- **ActivityLog:**  
  Audit trail aktivitas user, monitoring keamanan & compliance.

### 9. **Pengeluaran Operasional**

- **OperationalExpense:**  
  Semua biaya operasional di luar transaksi penjualan (gaji, listrik, sewa, dll).

### 10. **Advance/Supporting (Opsional)**

- **EmployeeAttendance, EmployeeLeave, EmployeePerformanceSummary, ContactHistory, MasterProductVariant, TransactionAuditTrail, TransactionRefund, TransactionReceipt, ExpenseReport, Notification, AuthSession, RBACAssignment, ImportHistory, ExportHistory, FileMetadata:**  
  Struct-struct ini digunakan untuk kebutuhan advance sesuai breakdown service, mendukung fitur-fitur tambahan dan pelacakan lebih detail.

### 11. Financial Summary (Baru)

- **CompanyFinancialSummary:**  
  Struct untuk menyimpan ringkasan data finansial perusahaan seperti pendapatan, pengeluaran, dan laba dalam periode tertentu.

### 12. Advance Product History (Baru)

- **MasterProductHistory:**  
  Struct untuk menyimpan histori perubahan dan aktivitas pada master product, mendukung pelacakan versi dan audit produk.
- **PurchaseOrderHistory** : Struct untuk menyimpan histori perubahan pada Purchase Order, mendukung audit dan tracking perubahan.

### 13. **Laporan & Statistik**

- **SalesReport**  
  Merangkum data penjualan dalam periode tertentu.

- **StockReport**  
  Menampilkan laporan stok barang secara menyeluruh.

- **ProfitLossReport**  
  Laporan laba rugi perusahaan selama periode tertentu.

- **EmployeePerformanceReport**  
  Statistik dan evaluasi kinerja karyawan.

- **CustomerActivityReport**  
  Analisis aktivitas dan transaksi pelanggan.

- **ShiftAttendance**  
  Data kehadiran karyawan per shift.

- **ShiftSwap**  
  Catatan tukar shift antar karyawan.

- **StockMovementSummary**  
  Ringkasan mutasi stok produk per periode.

- **StockTransfer**  
  Informasi transfer stok antar cabang atau toko.

- **StockTransferHistory**  
  Riwayat perubahan transfer stok.

- **StoreProductStockUpdate**  
  Pembaruan stok produk per toko.

- **TransactionAuditLog**  
  Catatan audit perubahan transaksi.

- **PaymentInfo**  
  Informasi terkait pembayaran transaksi.

- **TransactionSummary**  
  Ringkasan keseluruhan transaksi.

- **UserLoginHistory**  
  Riwayat login pengguna aplikasi.

---

## 🗂️ Konvensi Penamaan & Struktur

- Nama file mengikuti entitas/tabel (snake_case), satu entitas utama per file.
- Jika terdapat tabel pivot/relasi (misal: `employee_roles`), digabung di file terkait.
- Setiap model menggunakan tag `db` untuk ORM/SQL dan `json` untuk kebutuhan API/DTO.
- Jika ada entitas baru/perubahan ERD, **wajib update README ini**.
- **Struct advance/supporting harus didaftarkan di sini agar terdokumentasi dengan baik.**

---

## 🧑‍💻 Contoh Struct Model (Standar)

```go
type ExampleModel struct {
	ID        uuid.UUID      `db:"id" json:"id"`
	Name      string         `db:"name" json:"name"`
	CreatedAt time.Time      `db:"created_at" json:"created_at"`
	UpdatedAt time.Time      `db:"updated_at" json:"updated_at"`
}
```

---

## 📎 Lihat juga:

- [README service/domain](../core/services/readme.md) — untuk logika bisnis & usecase.
- [README repository](../core/repository/readme.md) — untuk interface repository & mapping DB.

---

**Terakhir update:** 2025-06-24 (andre)

---

**Note:** Struct-struct di atas merupakan kebutuhan advance sesuai breakdown service.  
Jika ada entitas baru atau perubahan field, wajib update README ini dan file Go terkait.

---
