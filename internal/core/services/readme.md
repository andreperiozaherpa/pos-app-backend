# internal/core/services

Direktori ini berisi _interface_ service (use case layer) untuk setiap entitas utama aplikasi.  
Service layer bertanggung jawab mengatur logika bisnis, memanggil repository, dan menjalankan aturan domain.

Gunakan folder ini untuk:

- Mendefinisikan kontrak use case (logika bisnis) per domain
- Mempermudah unit testing dengan dependency injection
- Memisahkan logika bisnis dari infrastruktur

---

## 📋 Checklist Service

> **Status:**  
> ⬜ Belum = method service belum dibuat  
> ✅ Selesai = method service sudah dibuat

| No  | File Name                   | Interface Name         | Status     | Catatan                            |
| --- | --------------------------- | ---------------------- | ---------- | ---------------------------------- |
| 1   | user_service.go             | UserService            | ✅ Selesai |                                    |
| 2   | employee_service.go         | EmployeeService        | ✅ Selesai |                                    |
| 3   | customer_service.go         | CustomerService        | ✅ Selesai |                                    |
| 4   | role_service.go             | RoleService            | ✅ Selesai |                                    |
| 5   | permission_service.go       | PermissionService      | ✅ Selesai |                                    |
| 6   | supplier_service.go         | SupplierService        | ⬜ Belum   |                                    |
| 7   | master_product_service.go   | MasterProductService   | ⬜ Belum   |                                    |
| 8   | store_product_service.go    | StoreProductService    | ⬜ Belum   |                                    |
| 9   | tax_rate_service.go         | TaxRateService         | ⬜ Belum   |                                    |
| 10  | transaction_service.go      | TransactionService     | ⬜ Belum   |                                    |
| 11  | transaction_item_service.go | TransactionItemService | ⬜ Belum   | Opsional                           |
| 12  | discount_service.go         | DiscountService        | ⬜ Belum   |                                    |
| 13  | shift_service.go            | ShiftService           | ⬜ Belum   |                                    |
| 14  | purchase_order_service.go   | PurchaseOrderService   | ⬜ Belum   |                                    |
| 15  | stock_transfer_service.go   | StockTransferService   | ⬜ Belum   |                                    |
| 16  | stock_movement_service.go   | StockMovementService   | ⬜ Belum   |                                    |
| 17  | activity_log_service.go     | ActivityLogService     | ⬜ Belum   |                                    |
| 18  | expense_service.go          | ExpenseService         | ⬜ Belum   |                                    |
| 19  | company_service.go          | CompanyService         | ⬜ Belum   |                                    |
| 20  | store_service.go            | StoreService           | ⬜ Belum   |                                    |
| 21  | business_line_service.go    | BusinessLineService    | ⬜ Belum   |                                    |
| 22  | rbac_service.go             | RBACService            | ⬜ Belum   | Opsional gabungan role/permission  |
| 23  | notification_service.go     | NotificationService    | ⬜ Belum   | Opsional notifikasi (email/SMS/WA) |
| 24  | auth_service.go             | AuthService            | ⬜ Belum   | Opsional autentikasi terpisah      |
| 25  | report_service.go           | ReportService          | ⬜ Belum   | Opsional agregat laporan           |
| 26  | import_export_service.go    | ImportExportService    | ⬜ Belum   | Opsional bulk import/export        |
| 27  | file_storage_service.go     | FileStorageService     | ⬜ Belum   | Opsional upload file/gambar        |

---

## 📑 Rancangan Method Utama & Custom per Service

### **UserService**

- ✅ `RegisterUser` _(wajib)_: Membuat user baru (employee/customer)
- ✅ `GetUserByID` _(wajib)_: Mengambil user berdasarkan ID
- ✅ `GetUserByUsername` _(wajib)_: Mengambil user berdasarkan username
- ✅ `GetUserByEmail` _(wajib)_: Mengambil user berdasarkan email
- ✅ `GetUserByPhoneNumber` _(wajib)_: Mengambil user berdasarkan nomor telepon
- ✅ `UpdateUser` _(wajib)_: Memperbarui data user
- ✅ `DeleteUser` _(wajib)_: Menghapus user berdasarkan ID
- ✅ `ListUsersByType` _(wajib)_: Mengambil list user berdasarkan tipe (EMPLOYEE/CUSTOMER)
- ✅ `AuthenticateUser` _(custom wajib)_: Validasi login (username/email/phone + password)
- ✅ `ChangePassword` _(custom wajib)_: Ganti password user
- ✅ `ActivateUser` _(opsional)_: Aktifkan user non-aktif
- ✅ `DeactivateUser` _(opsional)_: Nonaktifkan user (soft block)
- ✅ `SearchUsers` _(opsional)_: Pencarian user fleksibel
- ⬜ `ResetPassword` _(opsional)_: Reset password via email/OTP
- ⬜ `LockUserAccount` _(opsional)_: Mengunci akun setelah beberapa kali gagal login
- ⬜ `UnlockUserAccount` _(opsional)_: Membuka kunci akun
- ⬜ `UpdateUserProfilePicture` _(opsional)_: Update foto profil user
- ⬜ `ListUserRoles` _(opsional)_: Mendapatkan role-role user
- ⬜ `GetUserLoginHistory` _(opsional)_: Mendapatkan histori login user
- ⬜ `SendUserVerificationEmail` _(opsional)_: Kirim email verifikasi user
- ⬜ `VerifyUserEmailToken` _(opsional)_: Verifikasi token email user
- ⬜ `SendUserOTP` _(opsional)_: Kirim OTP ke user
- ⬜ `ValidateUserOTP` _(opsional)_: Validasi OTP user
- ⬜ `ListUserActivityLogs` _(opsional)_: Daftar aktivitas user
- ⬜ `SetUserTwoFactorAuth` _(opsional)_: Mengatur two-factor authentication

### **EmployeeService**

- ✅ `RegisterEmployee` _(wajib)_: Membuat employee baru
- ✅ `GetEmployeeByID` _(wajib)_: Mengambil employee by ID
- ✅ `UpdateEmployee` _(wajib)_: Update data employee
- ✅ `DeleteEmployee` _(wajib)_: Delete employee
- ✅ `ListEmployeesByCompanyID` _(wajib)_: Daftar karyawan per company
- ✅ `ListEmployeesByStoreID` _(wajib)_: Daftar karyawan per store
- ⬜ `AssignEmployeeToStore` _(opsional)_: Menetapkan employee ke store tertentu
- ⬜ `UpdateEmployeeStatus` _(opsional)_: Update status kerja (aktif/nonaktif/cuti)
- ⬜ `GetEmployeeAttendance` _(opsional)_: Mendapatkan absensi employee
- ⬜ `ListEmployeeRoles` _(opsional)_: Mendapatkan role-role employee
- ⬜ `ListEmployeeAttendanceByDateRange` _(opsional)_: Daftar absensi employee dalam rentang tanggal
- ⬜ `GetEmployeeLeaveHistory` _(opsional)_: Riwayat cuti employee
- ⬜ `AssignEmployeeRole` _(opsional)_: Menetapkan role ke employee
- ⬜ `RemoveEmployeeFromStore` _(opsional)_: Menghapus penugasan employee dari store
- ⬜ `GetEmployeePerformanceSummary` _(opsional)_: Ringkasan kinerja employee
- ⬜ `ExportEmployeeData` _(opsional)_: Export data employee ke excel/CSV

### **CustomerService**

- ✅ `RegisterCustomer` _(wajib)_: Membuat customer baru
- ✅ `GetCustomerByID` _(wajib)_: Mengambil customer by ID
- ✅ `UpdateCustomer` _(wajib)_: Update data customer
- ✅ `DeleteCustomer` _(wajib)_: Delete customer
- ✅ `ListCustomersByCompanyID` _(wajib)_: Daftar customer per company
- ⬜ `SearchCustomers` _(opsional)_: Pencarian customer fleksibel
- ⬜ `DeactivateCustomer` _(opsional)_: Nonaktifkan customer (soft delete)
- ⬜ `ListCustomerTransactions` _(opsional)_: Mendapatkan transaksi customer
- ⬜ `GetCustomerLoyaltyPoints` _(opsional)_: Mendapatkan poin loyalitas customer
- ⬜ `UpdateCustomerLoyaltyPoints` _(opsional)_: Update poin loyalitas customer
- ⬜ `ExportCustomers` _(opsional)_: Export data customer ke excel/CSV
- ⬜ `BulkImportCustomers` _(opsional)_: Import massal customer
- ⬜ `GetCustomerContactHistory` _(opsional)_: Riwayat kontak customer (call/email/wa)

### **RoleService**

- ✅ `CreateRole` _(wajib)_: Membuat role baru
- ✅ `GetRoleByID` _(wajib)_: Mengambil role berdasarkan ID
- ✅ `UpdateRole` _(wajib)_: Memperbarui role
- ✅ `DeleteRole` _(wajib)_: Menghapus role
- ✅ `ListRoles` _(wajib)_: List semua role
- ✅ `AssignRoleToEmployee` _(opsional)_: Menetapkan role ke employee
- ⬜ `ListRolePermissions` _(opsional)_: Mendapatkan permission-role
- ⬜ `RemoveRoleFromEmployee` _(opsional)_: Menghapus role dari employee
- ⬜ `AssignRoleToMultipleEmployees` _(opsional)_: Menetapkan role ke banyak employee sekaligus
- ⬜ `ListUsersByRole` _(opsional)_: Mendapatkan daftar user untuk role tertentu
- ⬜ `CloneRole` _(opsional)_: Mengkloning role beserta permission
- ⬜ `ExportRoles` _(opsional)_: Export daftar role

### **PermissionService**

- ✅ `CreatePermission` _(wajib)_: Membuat permission baru
- ✅ `GetPermissionByID` _(wajib)_: Mengambil permission berdasarkan ID
- ✅ `UpdatePermission` _(wajib)_: Memperbarui permission
- ✅ `DeletePermission` _(wajib)_: Menghapus permission
- ✅ `ListPermissions` _(wajib)_: List semua permission
- ⬜ `AssignPermissionToRole` _(opsional)_: Menetapkan permission ke role
- ⬜ `RemovePermissionFromRole` _(opsional)_: Menghapus permission dari role
- ⬜ `ListRolesByPermission` _(opsional)_: Mendapatkan role-role yang memiliki permission tertentu
- ⬜ `BulkAssignPermissionsToRole` _(opsional)_: Assign banyak permission ke role sekaligus
- ⬜ `ExportPermissions` _(opsional)_: Export daftar permission

### **SupplierService**

- ⬜ `CreateSupplier` _(wajib)_: Membuat supplier baru
- ⬜ `GetSupplierByID` _(wajib)_: Mengambil supplier berdasarkan ID
- ⬜ `UpdateSupplier` _(wajib)_: Memperbarui supplier
- ⬜ `DeleteSupplier` _(wajib)_: Menghapus supplier
- ⬜ `ListSuppliersByCompanyID` _(wajib)_: List supplier pada company
- ⬜ `SearchSuppliers` _(opsional)_: Search by name/category/city
- ⬜ `BulkImportSuppliers` _(opsional)_: Import excel/CSV supplier secara massal
- ⬜ `FindSupplierByPhoneOrEmail` _(opsional)_: Cari supplier berdasarkan no hp/email
- ⬜ `DeactivateSupplier` _(opsional)_: Nonaktifkan supplier (soft delete)
- ⬜ `ApproveSupplier` _(opsional)_: Menyetujui supplier baru
- ⬜ `GetSupplierTransactions` _(opsional)_: Mendapatkan transaksi supplier
- ⬜ `ExportSuppliers` _(opsional)_: Export data supplier ke excel/CSV
- ⬜ `ListSupplierProducts` _(opsional)_: Daftar produk yang dimiliki supplier
- ⬜ `BulkUpdateSupplierStatus` _(opsional)_: Update status supplier secara massal
- ⬜ `GetSupplierContactHistory` _(opsional)_: Riwayat komunikasi dengan supplier

### **MasterProductService**

- ⬜ `CreateMasterProduct` _(wajib)_: Membuat produk pusat
- ⬜ `GetMasterProductByID` _(wajib)_: Mengambil produk pusat berdasarkan ID
- ⬜ `UpdateMasterProduct` _(wajib)_: Memperbarui produk pusat
- ⬜ `DeleteMasterProduct` _(wajib)_: Menghapus produk pusat
- ⬜ `ListMasterProductsByCompanyID` _(wajib)_: List produk pusat pada company
- ⬜ `GetMasterProductHistory` _(opsional)_: Melihat histori perubahan/master product
- ⬜ `DeactivateMasterProduct` _(opsional)_: Menonaktifkan produk pusat (soft delete/nonaktif)
- ⬜ `SearchMasterProducts` _(opsional)_: Pencarian produk pusat fleksibel
- ⬜ `BulkImportMasterProducts` _(opsional)_: Import produk pusat secara massal
- ⬜ `GetMasterProductStockLevels` _(opsional)_: Mendapatkan level stok produk pusat
- ⬜ `ArchiveMasterProduct` _(opsional)_: Mengarsipkan produk pusat
- ⬜ `RestoreMasterProduct` _(opsional)_: Mengembalikan produk pusat dari arsip
- ⬜ `ListMasterProductVariants` _(opsional)_: Daftar varian produk pusat
- ⬜ `ExportMasterProducts` _(opsional)_: Export produk pusat ke excel/CSV
- ⬜ `SyncMasterProductWithStoreProducts` _(opsional)_: Sinkronisasi produk pusat dengan produk toko

### **StoreProductService**

- ⬜ `CreateStoreProduct` _(wajib)_: Membuat produk toko
- ⬜ `GetStoreProductByID` _(wajib)_: Mengambil produk toko berdasarkan ID
- ⬜ `UpdateStoreProduct` _(wajib)_: Memperbarui produk toko
- ⬜ `DeleteStoreProduct` _(wajib)_: Menghapus produk toko
- ⬜ `ListStoreProductsByStoreID` _(wajib)_: List produk pada store
- ⬜ `UpdateStoreProductStock` _(opsional)_: Update stok produk toko
- ⬜ `SearchStoreProducts` _(opsional)_: Pencarian produk toko fleksibel
- ⬜ `BulkUpdateStoreProductStock` _(opsional)_: Update stok banyak produk toko sekaligus
- ⬜ `ArchiveStoreProduct` _(opsional)_: Mengarsipkan produk toko
- ⬜ `RestoreStoreProduct` _(opsional)_: Mengembalikan produk toko dari arsip
- ⬜ `ListStoreProductMovements` _(opsional)_: Daftar mutasi stok produk toko
- ⬜ `ExportStoreProducts` _(opsional)_: Export produk toko ke excel/CSV

### **TaxRateService**

- ⬜ `CreateTaxRate` _(wajib)_: Membuat tax rate baru
- ⬜ `GetTaxRateByID` _(wajib)_: Mengambil tax rate berdasarkan ID
- ⬜ `UpdateTaxRate` _(wajib)_: Memperbarui tax rate
- ⬜ `DeleteTaxRate` _(wajib)_: Menghapus tax rate
- ⬜ `ListTaxRatesByCompanyID` _(wajib)_: List tax rate pada company
- ⬜ `CalculateTaxForTransaction` _(opsional)_: Menghitung pajak untuk transaksi tertentu
- ⬜ `ArchiveTaxRate` _(opsional)_: Mengarsipkan tax rate
- ⬜ `RestoreTaxRate` _(opsional)_: Mengembalikan tax rate dari arsip
- ⬜ `ExportTaxRates` _(opsional)_: Export tax rate ke excel/CSV
- ⬜ `ListTaxRatesByDateRange` _(opsional)_: Daftar tax rate berdasarkan rentang tanggal

### **TransactionService**

- ⬜ `CreateTransaction` _(wajib)_: Membuat transaksi baru
- ⬜ `GetTransactionByID` _(wajib)_: Mengambil transaksi berdasarkan ID
- ⬜ `GetTransactionByCode` _(wajib)_: Mengambil transaksi berdasarkan kode unik
- ⬜ `ListTransactionsByStoreID` _(wajib)_: List transaksi pada store
- ⬜ `RefundTransaction` _(opsional)_: Melakukan refund transaksi (pembatalan/pengembalian)
- ⬜ `ExportTransactions` _(opsional)_: Export transaksi ke format excel/CSV
- ⬜ `RecalculateTransactionTotals` _(opsional)_: Hitung ulang total transaksi (misal ada update item/discount)
- ⬜ `GetTransactionAuditTrail` _(opsional)_: Mendapatkan histori perubahan transaksi/audit log
- ⬜ `ListTransactionsByCustomerID` _(opsional)_: List transaksi per customer
- ⬜ `VoidTransaction` _(opsional)_: Membatalkan transaksi
- ⬜ `ApplyDiscountToTransaction` _(opsional)_: Menerapkan diskon pada transaksi
- ⬜ `PrintReceipt` _(opsional)_: Cetak struk transaksi
- ⬜ `ProcessPayment` _(opsional)_: Memproses pembayaran transaksi
- ⬜ `ValidateTransactionStock` _(opsional)_: Validasi stok sebelum transaksi diproses
- ⬜ `ListTransactionsByDateRange` _(opsional)_: Daftar transaksi berdasarkan rentang tanggal
- ⬜ `ExportTransactionReceipts` _(opsional)_: Export struk transaksi
- ⬜ `GetTransactionPaymentStatus` _(opsional)_: Mendapatkan status pembayaran transaksi
- ⬜ `ListTransactionRefunds` _(opsional)_: Daftar refund transaksi
- ⬜ `GetTransactionSummaryByDay` _(opsional)_: Ringkasan transaksi per hari
- ⬜ `NotifyTransactionStatusChange` _(opsional)_: Kirim notifikasi perubahan status transaksi

### **TransactionItemService** _(opsional)_

- ⬜ `CreateTransactionItem` _(wajib)_: Membuat item transaksi
- ⬜ `GetTransactionItemByID` _(wajib)_: Mengambil item transaksi berdasarkan ID
- ⬜ `ListTransactionItemsByTransactionID` _(wajib)_: List item per transaksi
- ⬜ `UpdateTransactionItem` _(wajib)_: Memperbarui item transaksi
- ⬜ `DeleteTransactionItem` _(wajib)_: Menghapus item transaksi
- ⬜ `CalculateItemDiscount` _(opsional)_: Menghitung diskon item transaksi
- ⬜ `ListTransactionItemsByProductID` _(opsional)_: Daftar item transaksi berdasarkan produk
- ⬜ `ExportTransactionItems` _(opsional)_: Export item transaksi ke excel/CSV
- ⬜ `BulkUpdateTransactionItems` _(opsional)_: Update banyak item transaksi sekaligus

### **DiscountService**

- ⬜ `CreateDiscount` _(wajib)_: Membuat diskon baru
- ⬜ `GetDiscountByID` _(wajib)_: Mengambil diskon berdasarkan ID
- ⬜ `UpdateDiscount` _(wajib)_: Memperbarui diskon
- ⬜ `DeleteDiscount` _(wajib)_: Menghapus diskon
- ⬜ `ListDiscountsByCompanyID` _(wajib)_: List diskon pada company
- ⬜ `AssignDiscountToProduct` _(opsional)_: Menetapkan diskon ke produk tertentu
- ⬜ `BulkUpdateDiscounts` _(opsional)_: Update diskon secara massal (misal: periode/produk)
- ⬜ `FindActiveDiscounts` _(opsional)_: Cari diskon aktif
- ⬜ `CheckDiscountEligibility` _(opsional)_: Cek kelayakan diskon untuk transaksi/customer
- ⬜ `RemoveDiscountFromProduct` _(opsional)_: Menghapus diskon dari produk
- ⬜ `ListProductsByDiscountID` _(opsional)_: Daftar produk yang mendapat diskon tertentu
- ⬜ `ExportDiscounts` _(opsional)_: Export daftar diskon
- ⬜ `ArchiveDiscount` _(opsional)_: Mengarsipkan diskon
- ⬜ `RestoreDiscount` _(opsional)_: Mengembalikan diskon dari arsip

### **ShiftService**

- ⬜ `CreateShift` _(wajib)_: Membuat shift baru
- ⬜ `GetShiftByID` _(wajib)_: Mengambil shift berdasarkan ID
- ⬜ `ListShiftsByEmployeeID` _(wajib)_: List shift per employee
- ⬜ `ListShiftsByStoreAndDateRange` _(wajib)_: List shift pada store dan rentang tanggal
- ⬜ `UpdateShift` _(wajib)_: Update shift
- ⬜ `DeleteShift` _(wajib)_: Hapus shift
- ⬜ `GetShiftAttendance` _(opsional)_: Melihat absensi pada shift (siapa hadir, siapa absen)
- ⬜ `ApproveShiftSwap` _(opsional)_: Menyetujui/tolak permintaan tukar shift
- ⬜ `ExportShifts` _(opsional)_: Export data shift ke excel/CSV
- ⬜ `RecordCheckIn` _(opsional)_: Mencatat waktu check-in employee
- ⬜ `RecordCheckOut` _(opsional)_: Mencatat waktu check-out employee
- ⬜ `RequestShiftSwap` _(opsional)_: Mengajukan permintaan tukar shift
- ⬜ `CancelShift` _(opsional)_: Membatalkan shift
- ⬜ `ListShiftsByDateRange` _(opsional)_: Daftar shift berdasarkan tanggal
- ⬜ `ExportShiftAttendance` _(opsional)_: Export absensi shift ke excel/CSV
- ⬜ `BulkUpdateShifts` _(opsional)_: Update data shift secara massal
- ⬜ `ListShiftSwaps` _(opsional)_: Daftar permintaan tukar shift

### **PurchaseOrderService**

- ⬜ `CreatePurchaseOrder` _(wajib)_: Membuat PO baru
- ⬜ `GetPurchaseOrderByID` _(wajib)_: Mengambil PO berdasarkan ID
- ⬜ `UpdatePurchaseOrder` _(wajib)_: Memperbarui PO
- ⬜ `DeletePurchaseOrder` _(wajib)_: Menghapus PO
- ⬜ `ListPurchaseOrdersByStoreID` _(wajib)_: List PO pada store
- ⬜ `ApprovePurchaseOrder` _(opsional)_: Menyetujui PO
- ⬜ `ReceivePurchaseOrder` _(opsional)_: Menerima barang PO
- ⬜ `CancelPurchaseOrder` _(opsional)_: Membatalkan PO
- ⬜ `ListPurchaseOrdersBySupplierID` _(opsional)_: List PO berdasarkan supplier
- ⬜ `GeneratePurchaseOrderReport` _(opsional)_: Laporan PO
- ⬜ `ExportPurchaseOrders` _(opsional)_: Export PO ke excel/CSV
- ⬜ `ListPurchaseOrdersByDateRange` _(opsional)_: Daftar PO berdasarkan tanggal
- ⬜ `GetPurchaseOrderHistory` _(opsional)_: Riwayat perubahan PO
- ⬜ `NotifyPurchaseOrderStatusChange` _(opsional)_: Kirim notifikasi perubahan status PO

### **StockTransferService**

- ⬜ `CreateStockTransfer` _(wajib)_: Membuat transfer stok
- ⬜ `GetStockTransferByID` _(wajib)_: Mengambil transfer stok berdasarkan ID
- ⬜ `UpdateStockTransfer` _(wajib)_: Update transfer stok
- ⬜ `DeleteStockTransfer` _(wajib)_: Hapus transfer stok
- ⬜ `ListStockTransfersByCompanyID` _(wajib)_: List transfer stok company
- ⬜ `ApproveStockTransfer` _(opsional)_: Menyetujui transfer stok
- ⬜ `CancelStockTransfer` _(opsional)_: Membatalkan transfer stok
- ⬜ `ExportStockTransfers` _(opsional)_: Export transfer stok ke excel/CSV
- ⬜ `ListStockTransfersByDateRange` _(opsional)_: Daftar transfer stok berdasarkan tanggal
- ⬜ `GetStockTransferHistory` _(opsional)_: Riwayat perubahan transfer stok

### **StockMovementService**

- ⬜ `CreateStockMovement` _(wajib)_: Membuat mutasi stok
- ⬜ `ListStockMovementsByStoreProductID` _(wajib)_: List mutasi per produk toko
- ⬜ `GetStockMovementByID` _(opsional)_: Mendapatkan detail mutasi stok
- ⬜ `ListStockMovementsByDateRange` _(opsional)_: List mutasi berdasarkan rentang tanggal
- ⬜ `ExportStockMovements` _(opsional)_: Export mutasi stok ke excel/CSV
- ⬜ `GetStockMovementSummary` _(opsional)_: Ringkasan mutasi stok per produk/periode

### **ActivityLogService**

- ⬜ `CreateActivityLog` _(wajib)_: Membuat log aktivitas
- ⬜ `ListActivityLogsByUserID` _(wajib)_: List log aktivitas per user
- ⬜ `ListActivityLogsByCompanyID` _(wajib)_: List log aktivitas per company
- ⬜ `ListActivityLogsByStoreID` _(wajib)_: List log aktivitas per store
- ⬜ `SearchActivityLogs` _(opsional)_: Cari log aktivitas berdasarkan keyword, tanggal, atau aksi
- ⬜ `ExportActivityLogs` _(opsional)_: Export log aktivitas ke excel/CSV
- ⬜ `DeleteOldActivityLogs` _(opsional)_: Menghapus log aktivitas lama (retensi)
- ⬜ `GetActivityLogDetail` _(opsional)_: Mendapatkan detail satu log aktivitas
- ⬜ `ListActivityLogsByDateRange` _(opsional)_: Daftar log aktivitas per rentang tanggal

### **ExpenseService**

- ⬜ `CreateExpense` _(wajib)_: Membuat expense baru
- ⬜ `GetExpenseByID` _(wajib)_: Mengambil expense berdasarkan ID
- ⬜ `UpdateExpense` _(wajib)_: Memperbarui expense
- ⬜ `DeleteExpense` _(wajib)_: Hapus expense
- ⬜ `ListExpensesByCompanyID` _(wajib)_: List expense per company
- ⬜ `ListExpensesByStoreID` _(wajib)_: List expense per store
- ⬜ `ApproveExpense` _(opsional)_: Menyetujui expense
- ⬜ `GenerateExpenseReport` _(opsional)_: Laporan pengeluaran
- ⬜ `ExportExpenses` _(opsional)_: Export daftar expense ke excel/CSV
- ⬜ `ListExpensesByDateRange` _(opsional)_: Daftar expense berdasarkan tanggal
- ⬜ `ApproveMultipleExpenses` _(opsional)_: Menyetujui banyak expense sekaligus

### **CompanyService**

- ⬜ `CreateCompany` _(wajib)_: Membuat company baru
- ⬜ `GetCompanyByID` _(wajib)_: Mengambil company berdasarkan ID
- ⬜ `UpdateCompany` _(wajib)_: Memperbarui company
- ⬜ `DeleteCompany` _(wajib)_: Menghapus company
- ⬜ `ListAllCompanies` _(wajib)_: List seluruh company
- ⬜ `SearchCompanies` _(opsional)_: Pencarian company
- ⬜ `GetCompanyFinancialSummary` _(opsional)_: Ringkasan keuangan company
- ⬜ `ExportCompanies` _(opsional)_: Export daftar company ke excel/CSV
- ⬜ `ArchiveCompany` _(opsional)_: Mengarsipkan company
- ⬜ `RestoreCompany` _(opsional)_: Mengembalikan company dari arsip
- ⬜ `ListCompanyStores` _(opsional)_: Daftar store pada company

### **StoreService**

- ⬜ `CreateStore` _(wajib)_: Membuat store baru
- ⬜ `GetStoreByID` _(wajib)_: Mengambil store berdasarkan ID
- ⬜ `UpdateStore` _(wajib)_: Memperbarui store
- ⬜ `DeleteStore` _(wajib)_: Menghapus store
- ⬜ `ListStoresByBusinessLineID` _(wajib)_: List store per business line
- ⬜ `ListStoresByCompanyID` _(opsional)_: List store per company
- ⬜ `ActivateStore` _(opsional)_: Mengaktifkan store
- ⬜ `DeactivateStore` _(opsional)_: Menonaktifkan store
- ⬜ `ExportStores` _(opsional)_: Export daftar store ke excel/CSV
- ⬜ `ArchiveStore` _(opsional)_: Mengarsipkan store
- ⬜ `RestoreStore` _(opsional)_: Mengembalikan store dari arsip
- ⬜ `ListStoreEmployees` _(opsional)_: Daftar karyawan pada store

### **BusinessLineService**

- ⬜ `CreateBusinessLine` _(wajib)_: Membuat business line
- ⬜ `GetBusinessLineByID` _(wajib)_: Mengambil business line berdasarkan ID
- ⬜ `UpdateBusinessLine` _(wajib)_: Memperbarui business line
- ⬜ `DeleteBusinessLine` _(wajib)_: Menghapus business line
- ⬜ `ListBusinessLinesByCompanyID` _(wajib)_: List business line per company
- ⬜ `SearchBusinessLines` _(opsional)_: Pencarian business line
- ⬜ `ExportBusinessLines` _(opsional)_: Export daftar business line ke excel/CSV
- ⬜ `ArchiveBusinessLine` _(opsional)_: Mengarsipkan business line
- ⬜ `RestoreBusinessLine` _(opsional)_: Mengembalikan business line dari arsip

### **RBACService** _(opsional)_

- ⬜ `AssignRoleToUser` _(opsional)_: Memberi role ke user
- ⬜ `AssignPermissionToRole` _(opsional)_: Memberi permission ke role
- ⬜ `CheckUserPermission` _(opsional)_: Cek permission user
- ⬜ `RevokeRoleFromUser` _(opsional)_: Mencabut role dari user
- ⬜ `RevokePermissionFromRole` _(opsional)_: Mencabut permission dari role
- ⬜ `ListAllRBACAssignments` _(opsional)_: Daftar semua assignment role-permission-user
- ⬜ `ExportRBACConfig` _(opsional)_: Export konfigurasi RBAC

### **NotificationService** _(opsional)_

- ⬜ `SendNotification` _(opsional)_: Kirim notifikasi
- ⬜ `ScheduleNotification` _(opsional)_: Jadwalkan notifikasi
- ⬜ `CancelScheduledNotification` _(opsional)_: Membatalkan notifikasi terjadwal
- ⬜ `GetNotificationStatus` _(opsional)_: Mendapatkan status notifikasi
- ⬜ `ListNotificationsByUserID` _(opsional)_: Daftar notifikasi per user
- ⬜ `ExportNotifications` _(opsional)_: Export daftar notifikasi

### **AuthService** _(opsional)_

- ⬜ `Login` _(opsional)_: Login user
- ⬜ `Logout` _(opsional)_: Logout user
- ⬜ `RefreshToken` _(opsional)_: Refresh JWT token
- ⬜ `ValidateToken` _(opsional)_: Validasi JWT token
- ⬜ `ChangePassword` _(opsional)_: Ganti password (jika belum ada di UserService)
- ⬜ `SendPasswordResetLink` _(opsional)_: Kirim link reset password
- ⬜ `ValidatePasswordResetToken` _(opsional)_: Validasi token reset password
- ⬜ `GetAuthSessionInfo` _(opsional)_: Mendapatkan info sesi autentikasi user

### **ReportService** _(opsional)_

- ⬜ `GenerateSalesReport` _(opsional)_: Laporan penjualan
- ⬜ `GenerateStockReport` _(opsional)_: Laporan stok
- ⬜ `GenerateProfitLossReport` _(opsional)_: Laporan laba rugi
- ⬜ `GenerateEmployeePerformanceReport` _(opsional)_: Laporan kinerja karyawan
- ⬜ `GenerateCustomerActivityReport` _(opsional)_: Laporan aktivitas customer
- ⬜ `ExportReportToPDF` _(opsional)_: Export laporan ke PDF
- ⬜ `ScheduleReportGeneration` _(opsional)_: Menjadwalkan pembuatan laporan otomatis
- ⬜ `GenerateCustomReport` _(opsional)_: Membuat laporan custom sesuai filter

### **ImportExportService** _(opsional)_

- ⬜ `ImportData` _(opsional)_: Import data
- ⬜ `ExportData` _(opsional)_: Export data
- ⬜ `ValidateImportData` _(opsional)_: Validasi data sebelum import
- ⬜ `GenerateImportTemplate` _(opsional)_: Membuat template import
- ⬜ `ListImportHistory` _(opsional)_: Daftar histori import
- ⬜ `ListExportHistory` _(opsional)_: Daftar histori export
- ⬜ `CancelImportExportTask` _(opsional)_: Membatalkan proses import/export berjalan

### **FileStorageService** _(opsional)_

- ⬜ `UploadFile` _(opsional)_: Upload file/gambar
- ⬜ `DownloadFile` _(opsional)_: Download file/gambar
- ⬜ `DeleteFile` _(opsional)_: Menghapus file/gambar
- ⬜ `ListFiles` _(opsional)_: List file/gambar di storage
- ⬜ `GetFileMetadata` _(opsional)_: Mendapatkan metadata file
- ⬜ `ShareFileLink` _(opsional)_: Membagikan link file
- ⬜ `ArchiveFile` _(opsional)_: Mengarsipkan file
- ⬜ `RestoreFile` _(opsional)_: Mengembalikan file dari arsip
