# /Users/andre/Programing/aplikasi perusahaan/pos-backend/internal/data/postgres/readme.md

Direktori ini berisi implementasi repository spesifik untuk PostgreSQL.
Setiap file repository dalam direktori ini mengimplementasikan interface seperti `UserRepository`, `ProductRepository`, dll. menggunakan `database/sql` dan driver `github.com/lib/pq`.

---

## Pola Desain

### Unit of Work dengan `DBExecutor`

Untuk mendukung operasi transaksional yang melibatkan beberapa repository (misalnya, dalam sebuah service), lapisan ini menggunakan `DBExecutor` interface.

- **`db_executor.go`**: File ini mendefinisikan `DBExecutor` interface yang diabstraksi dari metode-metode umum `*sql.DB` dan `*sql.Tx`. Hal ini memungkinkan setiap repository untuk beroperasi baik dengan koneksi database langsung maupun di dalam sebuah transaksi yang sedang berjalan, tanpa mengubah kode repository itu sendiri.
  - **Status: ✅ SELESAI & TERUJI**

---

## Tinjauan Lengkap Kebutuhan Repository

Berikut adalah daftar repository yang dibutuhkan berdasarkan ERD proyek, beserta status pengerjaannya.

### 1. Modul: core_tenant_and_organization

- `CompanyRepository`: Mengelola data perusahaan (tenant).
  - **Status: ✅ SELESAI & TERUJI**
- `BusinessLineRepository`: Mengelola lini bisnis yang dimiliki setiap perusahaan.
  - **Status: ✅ SELESAI & TERUJI**
- `StoreRepository`: Mengelola data toko, termasuk hierarkinya (pusat, cabang, ranting).
  - **Status: ✅ SELESAI & TERUJI**

### 2. Modul: user_management

- `UserRepository`: Mengelola data pengguna dasar (login, info kontak).
  - **Status: ✅ SELESAI & TERUJI**
- `EmployeeRepository`: Mengelola data spesifik karyawan dan menghubungkan `User` ke `Company` dan `Store`.
  - **Status: ✅ SELESAI & TERUJI**
- `CustomerRepository`: Mengelola data spesifik pelanggan/member.
  - **Status: ✅ SELESAI & TERUJI**
- `RoleRepository`: Mengelola daftar peran yang tersedia dalam sistem.
  - **Status: ✅ SELESAI & TERUJI**
- `EmployeeRoleRepository`: Mengelola hubungan antara karyawan dan peran.
  - **Status: ✅ SELESAI & TERUJI**
- `PermissionRepository`: Mengelola daftar izin/hak akses.
  - **Status: ✅ SELESAI & TERUJI**
- `RolePermissionRepository`: Mengelola hubungan antara peran dan izin.
- **Status: ✅ SELESAI & TERUJI**

### 3. Modul: product_and_supplier_management

- `SupplierRepository`: Mengelola data pemasok untuk setiap perusahaan.
  - **Status: ✅ SELESAI & TERUJI**
- `StoreProductRepository`: Mengelola data produk di setiap toko, termasuk harga, stok, barcode, dll.
  - **Status: ✅ SELESAI & TERUJI**
- `MasterProductRepository`: Mengelola definisi produk master yang tidak spesifik toko.
- **Status: ✅ SELESAI & TERUJI**
- `TaxRateRepository`: Mengelola tarif pajak yang berlaku.
- **Status: ✅ SELESAI & TERUJI**

### 4. Modul: transaction_management

- `TransactionRepository`: Mengelola data transaksi dan item-itemnya secara atomik.
  - **Status: ✅ SELESAI & TERUJI**

### 5. Modul: shift_management

- `ShiftRepository`: Mengelola jadwal shift, waktu check-in, dan check-out karyawan.
  - **Status: ✅ SELESAI & TERUJI**

### 6. Modul: purchasing_and_stock_management

- `PurchaseOrderRepository`: Mengelola pesanan pembelian ke supplier, termasuk item-itemnya.
  - **Status: ✅ SELESAI & TERUJI**
- `InternalStockTransferRepository`: Mengelola transfer barang antar toko dalam satu perusahaan.
  - **Status: ✅ SELESAI & TERUJI**
- `StockMovementRepository`: **Sangat Penting.** Mencatat setiap perubahan stok (penjualan, pembelian, transfer, penyesuaian) untuk tujuan audit dan pelacakan.
  - **Status: ✅ SELESAI & TERUJI**

### 7. Modul: discount_management

- `DiscountRepository`: Mengelola aturan diskon (misalnya, diskon 10% untuk kategori tertentu).
  - **Status: ✅ SELESAI & TERUJI**
- `AppliedTransactionDiscountRepository`: Mengelola tabel pivot untuk mencatat diskon apa yang diterapkan pada sebuah transaksi.
  - **Status: ✅ SELESAI & TERUJI**
- `AppliedItemDiscountRepository`: Mengelola tabel pivot untuk mencatat diskon apa yang diterapkan pada item transaksi.
  - **Status: ✅ SELESAI & TERUJI**

### 8. Modul: auditing

- `ActivityLogRepository`: Menyimpan log dari setiap tindakan penting yang dilakukan oleh pengguna.
  - **Status: ✅ SELESAI & TERUJI**

### 9. Modul: expense_management

- `OperationalExpenseRepository`: Mengelola data pengeluaran operasional.
  - **Status: ✅ SELESAI & TERUJI**
