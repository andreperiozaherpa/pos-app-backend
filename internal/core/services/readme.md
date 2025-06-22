# /Users/andre/Programing/aplikasi perusahaan/pos-backend/internal/core/services/readme.md

Direktori ini berisi implementasi dari **Lapisan Service**. Lapisan ini adalah jantung dari aplikasi yang berisi semua logika bisnis inti (use cases).

Setiap file service (misalnya, `auth_service.go`) bertanggung jawab atas satu modul fungsional. Service ini mengorkestrasi panggilan ke satu atau lebih repository untuk menyelesaikan sebuah tugas, menerapkan validasi, dan memastikan aturan bisnis terpenuhi.

Selain file service utama, direktori ini juga dapat berisi file pendukung yang digunakan bersama oleh service lain dalam paket ini, seperti `errors.go` yang mendefinisikan error-error spesifik untuk logika bisnis.

Contoh file:

- `auth_service.go`: Menangani logika autentikasi dan otorisasi.
- `user_service.go`: Menangani operasi pengguna umum.
- `employee_service.go`: Menangani logika bisnis terkait karyawan.
- `customer_service.go`: Menangani logika bisnis terkait pelanggan/member.
- `role_service.go`: Menangani logika bisnis terkait peran pengguna.
- `permission_service.go`: Menangani logika bisnis terkait izin/hak akses.
- `company_service.go`: Menangani logika bisnis terkait perusahaan.
- `business_line_service.go`: Menangani logika bisnis terkait lini bisnis.
- `store_service.go`: Menangani logika bisnis terkait toko.
- `master_product_service.go`: Menangani logika bisnis terkait master produk.
- `store_product_service.go`: Menangani logika bisnis terkait produk di toko.
- `supplier_service.go`: Menangani logika bisnis terkait supplier.
- `tax_rate_service.go`: Menangani logika bisnis terkait tarif pajak.
- `transaction_service.go`: Menangani logika bisnis terkait transaksi penjualan.
- `shift_service.go`: Menangani logika bisnis terkait shift karyawan.
- `purchase_order_service.go`: Menangani logika bisnis terkait pesanan pembelian.
- `internal_stock_transfer_service.go`: Menangani logika bisnis terkait transfer stok internal.
- `activity_log_service.go`: Menangani logika bisnis terkait pencatatan aktivitas.
- `discount_service.go`: Menangani logika bisnis terkait diskon.
- `operational_expense_service.go`: Menangani logika bisnis terkait pengeluaran operasional.

## File Pendukung

Direktori ini juga berisi file-file pendukung yang digunakan oleh satu atau lebih service untuk fungsionalitas umum dalam lapisan ini.

- `errors.go`: Mendefinisikan error-error kustom yang spesifik untuk logika bisnis di lapisan service.
  - **Status: ✅ SELESAI & TERUJI**

---

## Detail Fungsionalitas Service

Berikut adalah daftar service yang akan diimplementasikan beserta metode-metode utamanya:

- **`AuthService`**:
  - **Status: ✅ SELESAI & TERUJI**
  - `Login(ctx context.Context, username, password string) (*models.User, string, error)`
  - `Logout(ctx context.Context, token string) error`
  - `ValidateToken(ctx context.Context, token string) (*models.User, error)`
- **`UserService`**:
  - **Status: ✅ SELESAI & TERUJI**
  - `RegisterUser(ctx context.Context, user *models.User) error`
  - `GetUserProfile(ctx context.Context, userID uuid.UUID) (*models.User, error)`
  - `UpdateUserProfile(ctx context.Context, user *models.User) error`
- **`EmployeeService`**:
  - **Status: ✅ SELESAI & TERUJI**
  - `AddEmployee(ctx context.Context, employee *models.Employee, user *models.User, password string, roleIDs []int32) (*models.Employee, error)`
  - `UpdateEmployee(ctx context.Context, employee *models.Employee) error`
  - `DeactivateEmployee(ctx context.Context, employeeUserID uuid.UUID) error`
  - `AssignRoleToEmployee(ctx context.Context, employeeUserID uuid.UUID, roleID int32) error`
  - `RemoveRoleFromEmployee(ctx context.Context, employeeUserID uuid.UUID, roleID int32) error`
  - `ListEmployees(ctx context.Context, companyID uuid.UUID) ([]*models.Employee, error)`
    **`CustomerService`**:
  - **Status: ✅ SELESAI & TERUJI**
  - `AddCustomer(ctx context.Context, customer *models.Customer, user *models.User, password string) (*models.Customer, error)`
  - `UpdateCustomer(ctx context.Context, customer *models.Customer, user *models.User) error`
  - `GetCustomerByPhoneNumber(ctx context.Context, phoneNumber string) (*models.Customer, error)`
  - `UpdateCustomerPoints(ctx context.Context, customerUserID uuid.UUID, points int) error`
- **`RoleService`**:
  - **Status: ⬜ TO-DO**
  - `CreateRole(ctx context.Context, role *models.Role) error`
  - `UpdateRole(ctx context.Context, role *models.Role) error`
  - `DeleteRole(ctx context.Context, roleID int32) error`
  - `AssignPermissionToRole(ctx context.Context, roleID int32, permissionID int32) error`
  - `RemovePermissionFromRole(ctx context.Context, roleID int32, permissionID int32) error`
  - `ListRoles(ctx context.Context) ([]*models.Role, error)`
- **`PermissionService`**:
  - **Status: ⬜ TO-DO**
  - `CreatePermission(ctx context.Context, permission *models.Permission) error`
  - `ListPermissions(ctx context.Context) ([]*models.Permission, error)`
- **`CompanyService`**:
  - **Status: ⬜ TO-DO**
  - `CreateCompany(ctx context.Context, company *models.Company) error`
  - `UpdateCompany(ctx context.Context, company *models.Company) error`
  - `GetCompanyByID(ctx context.Context, companyID uuid.UUID) (*models.Company, error)`
- **`BusinessLineService`**:
  - **Status: ⬜ TO-DO**
  - `CreateBusinessLine(ctx context.Context, bl *models.BusinessLine) error`
  - `UpdateBusinessLine(ctx context.Context, bl *models.BusinessLine) error`
  - `ListBusinessLines(ctx context.Context, companyID uuid.UUID) ([]*models.BusinessLine, error)`
- **`StoreService`**:
  - **Status: ⬜ TO-DO**
  - `CreateStore(ctx context.Context, store *models.Store) error`
  - `UpdateStore(ctx context.Context, store *models.Store) error`
  - `ListStores(ctx context.Context, businessLineID uuid.UUID) ([]*models.Store, error)`
- **`MasterProductService`**:
  - **Status: ⬜ TO-DO**
  - `CreateMasterProduct(ctx context.Context, mp *models.MasterProduct) error`
  - `UpdateMasterProduct(ctx context.Context, mp *models.MasterProduct) error`
  - `ListMasterProducts(ctx context.Context, companyID uuid.UUID) ([]*models.MasterProduct, error)`
- **`StoreProductService`**:
  - **Status: ⬜ TO-DO**
  - `CreateStoreProduct(ctx context.Context, sp *models.StoreProduct) error`
  - `UpdateStoreProduct(ctx context.Context, sp *models.StoreProduct) error`
  - `ListStoreProducts(ctx context.Context, storeID uuid.UUID) ([]*models.StoreProduct, error)`
  - `AdjustStock(ctx context.Context, storeProductID uuid.UUID, quantity int, reason string, userID uuid.UUID) error`
- **`SupplierService`**:
  - **Status: ⬜ TO-DO**
  - `CreateSupplier(ctx context.Context, supplier *models.Supplier) error`
  - `UpdateSupplier(ctx context.Context, supplier *models.Supplier) error`
  - `ListSuppliers(ctx context.Context, companyID uuid.UUID) ([]*models.Supplier, error)`
- **`TaxRateService`**:
  - **Status: ⬜ TO-DO**
  - `CreateTaxRate(ctx context.Context, tr *models.TaxRate) error`
  - `UpdateTaxRate(ctx context.Context, tr *models.TaxRate) error`
  - `ListTaxRates(ctx context.Context, companyID uuid.UUID) ([]*models.TaxRate, error)`
- **`TransactionService`**:
  - **Status: ⬜ TO-DO**
  - `CreateTransaction(ctx context.Context, transaction *models.Transaction, items []models.TransactionItem) (*models.Transaction, error)`
  - `GetTransactionByID(ctx context.Context, transactionID uuid.UUID) (*models.Transaction, error)`
  - `ListTransactions(ctx context.Context, filter *TransactionFilter) ([]*models.Transaction, error)`
  - `RefundTransaction(ctx context.Context, transactionID uuid.UUID, refundItems []RefundItem) error`
- **`ShiftService`**:
  - **Status: ⬜ TO-DO**
  - `CreateShift(ctx context.Context, shift *models.Shift) error`
  - `CheckIn(ctx context.Context, shiftID uuid.UUID, employeeUserID uuid.UUID) error`
  - `CheckOut(ctx context.Context, shiftID uuid.UUID, employeeUserID uuid.UUID) error`
  - `ListShifts(ctx context.Context, employeeUserID uuid.UUID, dateFilter *ShiftDateFilter) ([]*models.Shift, error)`
- **`PurchaseOrderService`**:
  - **Status: ⬜ TO-DO**
  - `CreatePurchaseOrder(ctx context.Context, po *models.PurchaseOrder, items []models.PurchaseOrderItem) (*models.PurchaseOrder, error)`
  - `ListPurchaseOrders(ctx context.Context, storeID uuid.UUID, filter *POFilter) ([]*models.PurchaseOrder, error)`
  - `ReceivePurchaseOrder(ctx context.Context, poID uuid.UUID, receivedItems []ReceivedItem) error`
- **`InternalStockTransferService`**:
  - **Status: ⬜ TO-DO**
  - `CreateTransferRequest(ctx context.Context, transfer *models.InternalStockTransfer, items []models.InternalStockTransferItem) (*models.InternalStockTransfer, error)`
  - `ApproveTransfer(ctx context.Context, transferID uuid.UUID, approvedByUserID uuid.UUID) error`
  - `ShipTransfer(ctx context.Context, transferID uuid.UUID, shippedItems []ShippedItem, shippedByUserID uuid.UUID) error`
  - `ReceiveTransfer(ctx context.Context, transferID uuid.UUID, receivedItems []ReceivedItem, receivedByUserID uuid.UUID) error`
- **`ActivityLogService`**:
  - **Status: ⬜ TO-DO**
  - `RecordActivity(ctx context.Context, log *models.ActivityLog) error`
  - `ListActivityLogs(ctx context.Context, filter *ActivityLogFilter) ([]*models.ActivityLog, error)`
- **`DiscountService`**:
  - **Status: ⬜ TO-DO**
  - `CreateDiscount(ctx context.Context, discount *models.Discount) error`
  - `UpdateDiscount(ctx context.Context, discount *models.Discount) error`
  - `ApplyDiscountToTransaction(ctx context.Context, transactionID uuid.UUID, discountID uuid.UUID, amount float64) error`
  - `ApplyDiscountToItem(ctx context.Context, transactionItemID uuid.UUID, discountID uuid.UUID, amount float64) error`
- **`OperationalExpenseService`**:
  - **Status: ⬜ TO-DO**
  - `CreateExpense(ctx context.Context, expense *models.OperationalExpense) error`
  - `UpdateExpense(ctx context.Context, expense *models.OperationalExpense) error`
  - `ListExpenses(ctx context.Context, companyID uuid.UUID, filter *ExpenseFilter) ([]*models.OperationalExpense, error)`
