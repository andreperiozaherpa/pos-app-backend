# Dokumentasi Folder Migrasi Database

---

## 1. Penjelasan Folder & Tujuan

Folder ini berisi file-file migrasi database yang digunakan untuk mengelola perubahan skema dan data secara terstruktur dan terkontrol. Tujuannya adalah:

- Menjamin konsistensi perubahan skema di seluruh lingkungan (development, staging, production).
- Menyediakan riwayat perubahan yang terdokumentasi dengan baik.
- Memudahkan kolaborasi antar pengembang tanpa konflik.
- Memfasilitasi rollback jika terjadi masalah pada migrasi.

---

## 2. Panduan File Migrasi

- File migrasi berupa skrip SQL (`.sql`) atau format yang sesuai dengan alat migrasi yang digunakan (misal: goose, golang-migrate).
- Penamaan file mengikuti format timestamp `YYYYMMDDHHMMSS_deskripsi.sql` untuk memastikan urutan eksekusi yang tepat.
- Setiap migrasi harus bersifat **idempotent** (dapat dijalankan ulang tanpa efek samping) dan **atomic** (berhasil sepenuhnya atau gagal sepenuhnya).

---

## 3. Checklist Progres Migrasi

| File Migrasi                             | Deskripsi Singkat                                                                                                                                                                                                                      | Status     |
| ---------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------- |
| 20240716090000_initial_core_schema.sql   | Semua tabel inti: companies, business_lines, stores, users, employees, roles, role_permissions, employee_roles, permissions, customers                                                                                                 | ✅ SELESAI |
| 20240716091000_shift_and_supplier.sql    | Modul shift (shifts), shift_attendances, shift_swaps, suppliers                                                                                                                                                                        | ✅ SELESAI |
| 20240716092000_tax_and_product.sql       | Modul tax_rates, master_products, store_products                                                                                                                                                                                       | ✅ SELESAI |
| 20240716093000_transactions.sql          | Modul transaksi: transactions, transaction_items, payment_info                                                                                                                                                                         | ✅ SELESAI |
| 20240716094000_discount_and_applied.sql  | Modul diskon: discounts, applied_item_discounts, applied_transaction_discounts                                                                                                                                                         | ✅ SELESAI |
| 20240716095000_stock_and_purchasing.sql  | Modul pembelian, stok: purchase_orders, purchase_order_items, internal_stock_transfers, internal_stock_transfer_items, stock_movements, store_product_stock_updates, stock_movement_summaries, stock_reports, stock_transfer_histories | ✅ SELESAI |
| 20240716096000_audit_and_operational.sql | Audit & keuangan: activity_logs, operational_expenses                                                                                                                                                                                  | ✅ SELESAI |
| 20240716097000_history_tables.sql        | Tabel histori: master_product_histories, purchase_order_histories, stock_transfer_histories, user_login_histories                                                                                                                      | ✅ SELESAI |
| 20240716098000_report_tables.sql         | Modul report: sales_reports, stock_reports, profit_loss_reports, employee_performance_reports, customer_activity_reports, transaction_summaries                                                                                        | ✅ SELESAI |

---

## 4. Catatan & Konvensi Penamaan

- **Idempotent**: Pastikan setiap migrasi dapat dijalankan berulang kali tanpa menghasilkan error atau duplikasi data.
- **Atomic**: Setiap migrasi harus berhasil sepenuhnya atau gagal tanpa mengubah kondisi database.
- **Urutan Timestamp**: Penamaan file migrasi menggunakan timestamp memastikan urutan eksekusi yang benar.
- **Update Migrasi**: Jika perlu melakukan perubahan pada migrasi yang sudah ada, buat file migrasi baru dengan timestamp lebih baru untuk menghindari konflik dan menjaga histori perubahan.
- **Rollback**: Disarankan membuat migrasi rollback untuk setiap perubahan besar agar mudah kembali ke kondisi sebelumnya jika terjadi masalah.

---

Dokumentasi ini harus selalu diperbarui seiring dengan bertambahnya migrasi baru agar memudahkan tracking dan kolaborasi tim.
