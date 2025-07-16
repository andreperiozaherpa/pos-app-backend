# Struktur Folder Proyek Aplikasi Perusahaan

```
├── cmd
│   ├── api
│   │   ├── main.go
│   │   └── readme.md
│   └── readme.md
├── dokumentasi.md
├── erd.md
├── erd_pos
├── go.mod
├── go.sum
├── internal
│   ├── api
│   │   ├── handlers
│   │   │   └── readme.md
│   │   └── readme.md
│   ├── config
│   │   ├── config.go
│   │   └── readme.md
│   ├── core
│   │   ├── dto
│   │   ├── mock
│   │   │   └── db_mock.go
│   │   ├── readme.md
│   │   ├── readme_test.md
│   │   ├── repository
│   │   │   ├── activity_log_repository.go
│   │   │   ├── applied_item_discount_repository.go
│   │   │   ├── applied_transaction_discount_repository.go
│   │   │   ├── business_line_repository.go
│   │   │   ├── company_financial_summary_repository.go
│   │   │   ├── company_repository.go
│   │   │   ├── customer_activity_report_repository.go
│   │   │   ├── customer_repository.go
│   │   │   ├── discount_repository.go
│   │   │   ├── employee_performance_report_repository.go
│   │   │   ├── employee_repository.go
│   │   │   ├── employee_role_repository.go
│   │   │   ├── error.go
│   │   │   ├── expense_repository.go
│   │   │   ├── internal_stock_transfer_item_repository.go
│   │   │   ├── internal_stock_transfer_repository.go
│   │   │   ├── master_product_history_repository.go
│   │   │   ├── master_product_repository.go
│   │   │   ├── operational_expense_repository.go
│   │   │   ├── payment_info_repository.go
│   │   │   ├── permission_repository.go
│   │   │   ├── profit_loss_report_repository.go
│   │   │   ├── purchase_order_history_repository.go
│   │   │   ├── purchase_order_item_repository.go
│   │   │   ├── purchase_order_repository.go
│   │   │   ├── readme.md
│   │   │   ├── role_permission_repository.go
│   │   │   ├── role_repository.go
│   │   │   ├── sales_report_repository.go
│   │   │   ├── shift_attendance_repository.go
│   │   │   ├── shift_repository.go
│   │   │   ├── shift_swap_repository.go
│   │   │   ├── stock_movement_repository.go
│   │   │   ├── stock_movement_summary_repository.go
│   │   │   ├── stock_report_repository.go
│   │   │   ├── stock_transfer_history_repository.go
│   │   │   ├── stock_transfer_item_repository.go
│   │   │   ├── stock_transfer_repository.go
│   │   │   ├── store_product_repository.go
│   │   │   ├── store_product_stock_update_repository.go
│   │   │   ├── store_repository.go
│   │   │   ├── supplier_repository.go
│   │   │   ├── tax_rate_repository.go
│   │   │   ├── transaction_audit_log_repository.go
│   │   │   ├── transaction_item_repository.go
│   │   │   ├── transaction_repository.go
│   │   │   ├── transaction_summary_repository.go
│   │   │   ├── user_login_history_repository.go
│   │   │   └── user_repository.go
│   │   ├── services
│   │   │   ├── activity_log_service.go
│   │   │   ├── auth_service.go
│   │   │   ├── business_line_service.go
│   │   │   ├── company_service.go
│   │   │   ├── customer_service.go
│   │   │   ├── discount_service.go
│   │   │   ├── employee_service.go
│   │   │   ├── expense_service.go
│   │   │   ├── file_storage_service.go
│   │   │   ├── import_export_service.go
│   │   │   ├── master_product_service.go
│   │   │   ├── notification_service.go
│   │   │   ├── permission_service.go
│   │   │   ├── purchase_order_service.go
│   │   │   ├── rbac_service.go
│   │   │   ├── readme.md
│   │   │   ├── report_service.go
│   │   │   ├── role_service.go
│   │   │   ├── shift_service.go
│   │   │   ├── stock_movement_service.go
│   │   │   ├── stock_transfer_service.go
│   │   │   ├── store_product_service.go
│   │   │   ├── store_service.go
│   │   │   ├── supplier_service.go
│   │   │   ├── tax_rate_service.go
│   │   │   ├── transaction_item_service.go
│   │   │   ├── transaction_service.go
│   │   │   └── user_service.go
│   │   └── test
│   │       └── user_repository_test.go
│   ├── data
│   │   ├── postgres
│   │   │   ├── activity_log_repository_pg.go
│   │   │   ├── company_financial_summary_repository_pg.go
│   │   │   ├── customer_activity_report_repository_pg.go
│   │   │   ├── customer_repository_pg.go
│   │   │   ├── db.go
│   │   │   ├── employee_performance_report_repository_pg.go
│   │   │   ├── employee_repository_pg.go
│   │   │   ├── errors.go
│   │   │   ├── expense_repository_pg.go
│   │   │   ├── logger.go
│   │   │   ├── master_product_history_repository_pg.go
│   │   │   ├── master_product_repository_pg.go
│   │   │   ├── pagination.go
│   │   │   ├── payment_info_repository_pg.go
│   │   │   ├── permission_repository_pg.go
│   │   │   ├── profit_loss_report_repository_pg.go
│   │   │   ├── purchase_order_history_repository_pg.go
│   │   │   ├── purchase_order_item_repository_pg.go
│   │   │   ├── purchase_order_repository_pg.go
│   │   │   ├── readme.md
│   │   │   ├── role_repository_pg.go
│   │   │   ├── sales_report_repository_pg.go
│   │   │   ├── shift_attendance_repository_pg.go
│   │   │   ├── shift_repository_pg.go
│   │   │   ├── shift_swap_repository_pg.go
│   │   │   ├── stock_movement_summary_repository_pg.go
│   │   │   ├── stock_report_repository_pg.go
│   │   │   ├── stock_transfer_history_repository_pg.go
│   │   │   ├── stock_transfer_item_repository_pg.go
│   │   │   ├── stock_transfer_repository_pg.go
│   │   │   ├── store_product_repository_pg.go
│   │   │   ├── store_product_stock_update_repository_pg.go
│   │   │   ├── supplier_repository_pg.go
│   │   │   ├── transaction_audit_log_repository_pg.go
│   │   │   ├── transaction_item_repository_pg.go
│   │   │   ├── transaction_repository_pg.go
│   │   │   ├── transaction_summary_repository_pg.go
│   │   │   ├── user_login_history_repository_pg.go
│   │   │   ├── user_repository_pg.go
│   │   │   └── uuid.go
│   │   └── readme.md
│   ├── database
│   │   ├── postgres.go
│   │   └── readme.md
│   ├── models
│   │   ├── CompanyFinancialSummary.go
│   │   ├── activity_log.go
│   │   ├── applied_discount.go
│   │   ├── auth_session.go
│   │   ├── business_line.go
│   │   ├── company.go
│   │   ├── contact_history.go
│   │   ├── customer.go
│   │   ├── customer_activity_report.go
│   │   ├── discount.go
│   │   ├── employee.go
│   │   ├── employee_attendance.go
│   │   ├── employee_leave.go
│   │   ├── employee_performance.go
│   │   ├── employee_performance_report.go
│   │   ├── employee_role.go
│   │   ├── expense_report.go
│   │   ├── export_history.go
│   │   ├── file_metadata.go
│   │   ├── import_history.go
│   │   ├── import_result.go
│   │   ├── import_validation_result.go
│   │   ├── master_product.go
│   │   ├── master_product_history.go
│   │   ├── master_product_variant.go
│   │   ├── notification.go
│   │   ├── operational_expense.go
│   │   ├── payment_info.go
│   │   ├── permission.go
│   │   ├── product.go
│   │   ├── profit_loss_report.go
│   │   ├── purchase_order.go
│   │   ├── purchase_order_history.go
│   │   ├── rbac_assignment.go
│   │   ├── readme.md
│   │   ├── role.go
│   │   ├── role_permission.go
│   │   ├── sales_report.go
│   │   ├── shift.go
│   │   ├── shift_attendance.go
│   │   ├── shift_swap.go
│   │   ├── stock_movement.go
│   │   ├── stock_movement_summary.go
│   │   ├── stock_report.go
│   │   ├── stock_transfer.go
│   │   ├── stock_transfer_history.go
│   │   ├── stock_transfer_item.go
│   │   ├── store.go
│   │   ├── store_product_stock_update.go
│   │   ├── supplier.go
│   │   ├── tax_rate.go
│   │   ├── transaction.go
│   │   ├── transaction_audit.go
│   │   ├── transaction_audit_log.go
│   │   ├── transaction_receipt.go
│   │   ├── transaction_refund.go
│   │   ├── transaction_summary.go
│   │   ├── user.go
│   │   └── user_login_history.go
│   └── utils
│       ├── jwt.go
│       ├── password.go
│       └── readme.md
├── migrations
│   └── readme.md
├── Docs
├── readme.md
└── struktur.md


```

---
