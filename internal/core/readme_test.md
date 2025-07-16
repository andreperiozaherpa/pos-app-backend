# Dokumentasi Unit Test Repository

Direktori ini berisi unit test untuk repository di layer `internal/core/repository`.  
Unit test bertujuan memastikan setiap method pada repository berjalan sesuai ekspektasi dan menangani error dengan benar.

---

## 📋 Checklist Unit Test Repository

| No  | Repository                          | Status   | Keterangan                                     |
| --- | ----------------------------------- | -------- | ---------------------------------------------- |
| 1   | UserRepository                      | ⬜ Belum | Test CRUD, autentikasi, pencarian user         |
| 2   | EmployeeRepository                  | ⬜ Belum | Test CRUD employee, absensi, filter            |
| 3   | CustomerRepository                  | ⬜ Belum | Test CRUD customer, loyalti, histori transaksi |
| 4   | RoleRepository                      | ⬜ Belum | Test CRUD role, penugasan role ke user         |
| 5   | PermissionRepository                | ⬜ Belum | Test CRUD permission, assignment ke role       |
| 6   | SupplierRepository                  | ⬜ Belum | Test CRUD supplier, status aktif, arsip        |
| 7   | MasterProductRepository             | ⬜ Belum | Test CRUD produk pusat, stok, varian           |
| 8   | StoreProductRepository              | ⬜ Belum | Test CRUD produk toko, update stok/harga       |
| 9   | TransactionRepository               | ⬜ Belum | Test CRUD transaksi, audit trail               |
| 10  | TransactionItemRepository           | ⬜ Belum | Test CRUD item transaksi                       |
| 11  | PurchaseOrderRepository             | ⬜ Belum | Test CRUD purchase order                       |
| 12  | PurchaseOrderItemRepository         | ⬜ Belum | Test CRUD item purchase order                  |
| 13  | StockTransferRepository             | ⬜ Belum | Test CRUD transfer stok antar toko             |
| 14  | StockTransferItemRepository         | ⬜ Belum | Test CRUD item transfer stok                   |
| 15  | ShiftRepository                     | ⬜ Belum | Test CRUD shift, absensi, cuti                 |
| 16  | ActivityLogRepository               | ⬜ Belum | Test CRUD log aktivitas user                   |
| 17  | ExpenseRepository                   | ⬜ Belum | Test CRUD biaya operasional, laporan           |
| 18  | CompanyFinancialSummaryRepository   | ⬜ Belum | Test CRUD ringkasan keuangan perusahaan        |
| 19  | PurchaseOrderHistoryRepository      | ⬜ Belum | Test CRUD histori purchase order               |
| 20  | MasterProductHistoryRepository      | ⬜ Belum | Test CRUD histori produk pusat                 |
| 21  | SalesReportRepository               | ⬜ Belum | Test CRUD laporan penjualan                    |
| 22  | StockReportRepository               | ⬜ Belum | Test CRUD laporan stok                         |
| 23  | ProfitLossReportRepository          | ⬜ Belum | Test CRUD laporan laba rugi                    |
| 24  | EmployeePerformanceReportRepository | ⬜ Belum | Test CRUD laporan kinerja karyawan             |
| 25  | CustomerActivityReportRepository    | ⬜ Belum | Test CRUD laporan aktivitas customer           |
| 26  | ShiftAttendanceRepository           | ⬜ Belum | Test CRUD data absensi shift                   |
| 27  | ShiftSwapRepository                 | ⬜ Belum | Test CRUD permintaan tukar shift               |
| 28  | StockMovementSummaryRepository      | ⬜ Belum | Test CRUD ringkasan mutasi stok                |
| 29  | StockTransferHistoryRepository      | ⬜ Belum | Test CRUD histori transfer stok                |
| 30  | StoreProductStockUpdateRepository   | ⬜ Belum | Test CRUD update stok produk toko              |
| 31  | TransactionAuditLogRepository       | ⬜ Belum | Test CRUD histori audit transaksi              |
| 32  | PaymentInfoRepository               | ⬜ Belum | Test CRUD info pembayaran transaksi            |
| 33  | TransactionSummaryRepository        | ⬜ Belum | Test CRUD ringkasan transaksi                  |
| 34  | UserLoginHistoryRepository          | ⬜ Belum | Test CRUD histori login user                   |

---

## Catatan

- Setiap unit test harus mencakup skenario CRUD utama, validasi error, dan kondisi khusus.
- Gunakan mock database untuk isolasi pengujian.
- Pastikan semua dependensi di-mock agar test berjalan deterministik.
- Dokumentasikan hasil pengujian dan perbaiki kode bila ditemukan bug.

---

Dokumentasi ini akan diupdate secara berkala seiring dengan progres pengembangan dan pengujian repository.
