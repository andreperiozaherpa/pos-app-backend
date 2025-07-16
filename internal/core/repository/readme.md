# internal/core/repository

Direktori ini berisi _interface_ repository untuk setiap entitas utama aplikasi (mengikuti pola Clean Architecture).  
Setiap file `.go` hanya mendefinisikan interface tanpa implementasi detail database (pure domain).

Gunakan folder ini untuk:

- Definisi kontrak repository per domain/entity
- Kemudahan mocking pada unit test
- Fleksibilitas ganti DB engine tanpa ubah core service

---

## 📋 Checklist Implementasi Repository

| No  | File Name                                  | Interface Name                       | Status     | Catatan       |
| --- | ------------------------------------------ | ------------------------------------ | ---------- | ------------- |
| 1   | activity_log_repository.go                 | ActivityLogRepository                | ⬛ Selesai |               |
| 2   | applied_item_discount_repository.go        | AppliedItemDiscountRepository        | ⬛ Selesai | Pivot         |
| 3   | applied_transaction_discount_repository.go | AppliedTransactionDiscountRepository | ⬛ Selesai | Pivot         |
| 4   | business_line_repository.go                | BusinessLineRepository               | ⬛ Selesai | Org/Tenant    |
| 5   | company_repository.go                      | CompanyRepository                    | ⬛ Selesai | Org/Tenant    |
| 6   | customer_repository.go                     | CustomerRepository                   | ⬛ Selesai |               |
| 7   | discount_repository.go                     | DiscountRepository                   | ⬛ Selesai |               |
| 8   | employee_repository.go                     | EmployeeRepository                   | ⬛ Selesai |               |
| 9   | employee_role_repository.go                | EmployeeRoleRepository               | ⬛ Selesai | Pivot         |
| 10  | internal_stock_transfer_repository.go      | InternalStockTransferRepository      | ⬛ Selesai |               |
| 11  | internal_stock_transfer_item_repository.go | InternalStockTransferItemRepository  | ⬛ Selesai | Tambahan baru |
| 12  | master_product_repository.go               | MasterProductRepository              | ⬛ Selesai |               |
| 13  | operational_expense_repository.go          | OperationalExpenseRepository         | ⬛ Selesai |               |
| 14  | permission_repository.go                   | PermissionRepository                 | ⬛ Selesai |               |
| 15  | purchase_order_repository.go               | PurchaseOrderRepository              | ⬛ Selesai |               |
| 16  | purchase_order_item_repository.go          | PurchaseOrderItemRepository          | ⬛ Selesai | Tambahan baru |
| 17  | role_permission_repository.go              | RolePermissionRepository             | ⬛ Selesai | Pivot         |
| 18  | role_repository.go                         | RoleRepository                       | ⬛ Selesai |               |
| 19  | shift_repository.go                        | ShiftRepository                      | ⬛ Selesai |               |
| 20  | stock_movement_repository.go               | StockMovementRepository              | ⬛ Selesai |               |
| 21  | store_product_repository.go                | StoreProductRepository               | ⬛ Selesai |               |
| 22  | store_repository.go                        | StoreRepository                      | ⬛ Selesai | Org/Tenant    |
| 23  | supplier_repository.go                     | SupplierRepository                   | ⬛ Selesai |               |
| 24  | tax_rate_repository.go                     | TaxRateRepository                    | ⬛ Selesai |               |
| 25  | transaction_repository.go                  | TransactionRepository                | ⬛ Selesai |               |
| 26  | transaction_item_repository.go             | TransactionItemRepository            | ⬛ Selesai | Tambahan baru |
| 27  | user_repository.go                         | UserRepository                       | ⬛ Selesai |               |

---

> **Status:**  
> ⬜ Belum = interface belum dibuat  
> ⬛ Selesai = interface sudah dibuat  
> ✅ Full = interface + implementasi infra sudah ada
>
> **Update file ini setiap progress!**

---

**Terakhir update:** 2025-06-24 (AI/andre)
