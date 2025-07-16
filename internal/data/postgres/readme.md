# Dokumentasi Folder `internal/data/postgres`

Folder `internal/data/postgres` merupakan tempat implementasi repository yang berinteraksi langsung dengan database PostgreSQL. Setiap repository bertanggung jawab untuk operasi CRUD (Create, Read, Update, Delete), query custom, dan pengelolaan data entitas bisnis aplikasi.

---

## 📋 Checklist Implementasi Repository

| No  | Repository                                                                | Status     | Keterangan                                    |
| --- | ------------------------------------------------------------------------- | ---------- | --------------------------------------------- |
| 1   | UserRepository (`user_repository_pg.go`) ok                               | ✅ Selesai | CRUD & query user, autentikasi, pencarian     |
| 2   | EmployeeRepository (`employee_repository_pg.go`) ok                       | ✅ Selesai | CRUD employee, absensi, filter                |
| 3   | CustomerRepository (`customer_repository_pg.go`) ok                       | ✅ Selesai | CRUD customer, loyalti, histori transaksi     |
| 4   | RoleRepository (`role_repository_pg.go`) ok                               | ✅ Selesai | CRUD role, penugasan role ke user             |
| 5   | PermissionRepository (`permission_repository_pg.go`) ok                   | ✅ Selesai | CRUD permission, penugasan permission ke role |
| 6   | SupplierRepository (`supplier_repository_pg.go`) ok                       | ✅ Selesai | CRUD supplier, status aktif/arsip             |
| 7   | MasterProductRepository (`master_product_repository_pg.go`) ok            | ✅ Selesai | CRUD produk pusat, stok, varian               |
| 8   | StoreProductRepository (`store_product_repository_pg.go`)                 | ✅ Selesai | CRUD produk toko, update stok/harga           |
| 9   | TransactionRepository (`transaction_repository_pg.go`)                    | ✅ Selesai | CRUD transaksi, audit trail                   |
| 10  | TransactionItemRepository (`transaction_item_repository_pg.go`)           | ✅ Selesai | CRUD item transaksi                           |
| 11  | PurchaseOrderRepository (`purchase_order_repository_pg.go`)               | ✅ Selesai | CRUD purchase order                           |
| 12  | PurchaseOrderItemRepository (`purchase_order_item_repository_pg.go`)      | ✅ Selesai | CRUD item purchase order                      |
| 13  | StockTransferRepository (`stock_transfer_repository_pg.go`)               | ✅ Selesai | CRUD transfer stok antar toko                 |
| 14  | StockTransferItemRepository (`stock_transfer_item_repository_pg.go`)      | ✅ Selesai | CRUD item transfer stok                       |
| 15  | ShiftRepository (`shift_repository_pg.go`)                                | ✅ Selesai | CRUD shift, absensi, cuti                     |
| 16  | ActivityLogRepository (`activity_log_repository_pg.go`)                   | ✅ Selesai | CRUD log aktivitas user                       |
| 17  | ExpenseRepository (`expense_repository_pg.go`)                            | ✅ Selesai | CRUD biaya operasional, laporan               |
| 18  | CompanyFinancialSummaryRepository (`company_financial_summary_pg.go`)     | ✅ Selesai | CRUD ringkasan keuangan perusahaan            |
| 19  | PurchaseOrderHistoryRepository (`purchase_order_history_pg.go`)           | ✅ Selesai | CRUD histori purchase order                   |
| 20  | MasterProductHistoryRepository (`master_product_history_pg.go`)           | ✅ Selesai | CRUD histori produk pusat                     |
| 21  | SalesReportRepository (`sales_report_pg.go`)                              | ✅ Selesai | CRUD laporan penjualan                        |
| 22  | StockReportRepository (`stock_report_pg.go`)                              | ✅ Selesai | CRUD laporan stok                             |
| 23  | ProfitLossReportRepository (`profit_loss_report_pg.go`)                   | ✅ Selesai | CRUD laporan laba rugi                        |
| 24  | EmployeePerformanceReportRepository (`employee_performance_report_pg.go`) | ✅ Selesai | CRUD laporan kinerja karyawan                 |
| 25  | CustomerActivityReportRepository (`customer_activity_report_pg.go`)       | ✅ Selesai | CRUD laporan aktivitas customer               |
| 26  | ShiftAttendanceRepository (`shift_attendance_pg.go`)                      | ✅ Selesai | CRUD data absensi shift                       |
| 27  | ShiftSwapRepository (`shift_swap_pg.go`)                                  | ✅ Selesai | CRUD permintaan tukar shift                   |
| 28  | StockMovementSummaryRepository (`stock_movement_summary_pg.go`)           | ✅ Selesai | CRUD ringkasan mutasi stok                    |
| 29  | StockTransferHistoryRepository (`stock_transfer_history_pg.go`)           | ✅ Selesai | CRUD histori transfer stok                    |
| 30  | StoreProductStockUpdateRepository (`store_product_stock_update_pg.go`)    | ✅ Selesai | CRUD update stok produk toko                  |
| 31  | TransactionAuditLogRepository (`transaction_audit_log_pg.go`)             | ✅ Selesai | CRUD histori audit transaksi                  |
| 32  | PaymentInfoRepository (`payment_info_pg.go`)                              | ✅ Selesai | CRUD info pembayaran transaksi                |
| 33  | TransactionSummaryRepository (`transaction_summary_pg.go`)                | ✅ Selesai | CRUD ringkasan transaksi                      |
| 34  | UserLoginHistoryRepository (`user_login_history_pg.go`)                   | ✅ Selesai | CRUD histori login user                       |
| 35  | Database Connection Management (`db.go`)                                  | ✅ Selesai | Pengelolaan koneksi & transaksi DB            |
| 36  | Unit & Integration Testing                                                | ⬜ Belum   | Pengujian repository                          |
| 37  | Error Handling & Logging                                                  | ⬜ Belum   | Penanganan error & logging                    |
| 38  | Optimasi Query                                                            | ⬜ Belum   | Optimasi query & index                        |
| 39  | Pagination, Filtering, Sorting                                            | ⬜ Belum   | Fitur pagination, filter, sort                |
| 40  | Dokumentasi & Contoh Penggunaan                                           | ⬜ Belum   | Dokumentasi & contoh kode                     |
| 41  | Multi-Tenant Support                                                      | ⬜ Belum   | Filter data by tenant                         |
| 42  | Migration Tools Integration                                               | ⬜ Belum   | Integrasi tools migrasi DB                    |

> **Keterangan:**  
> Checklist di atas memonitor progres implementasi repository PostgreSQL untuk seluruh entitas aplikasi.  
> Setiap repository wajib menyediakan method CRUD dan custom sesuai kebutuhan bisnis.

---

## 📑 Detail Implementasi Repository

## Best Practice & Prinsip Implementasi

- Implementasikan interface repository di folder `internal/core/repository`.
- Selalu gunakan `context.Context` di setiap method repository.
- Gunakan dependency injection untuk kemudahan testing.
- Hindari import domain package secara langsung di layer data, gunakan interface.
- Tangani error dan logging secara konsisten.
- Gunakan transaction management untuk operasi yang memerlukan atomicity.
- Amankan query dari SQL Injection (gunakan parameterized query).
- Dokumentasikan setiap repository dan method secara detail.

---

Dokumentasi ini bertujuan untuk mendukung pengembangan, pemeliharaan, dan pengujian repository PostgreSQL secara efektif di aplikasi backend Go.
