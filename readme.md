# pos-app/backend

# pos-app/backend

Ini adalah direktori root untuk aplikasi Backend POS.

### Visi Aplikasi

Aplikasi ini adalah sistem Point of Sale (POS) multi-tenant yang dirancang untuk mendukung berbagai jenis perusahaan dengan struktur bisnis yang kompleks.

### Use Cases (Kasus Penggunaan)

Berikut adalah beberapa kasus penggunaan utama yang akan diimplementasikan dalam aplikasi, dikelompokkan berdasarkan modul fungsional.

#### 1. Auth & User Management

- **Login**: Pengguna (karyawan atau pelanggan) memasukkan kredensial dan mendapatkan akses ke sistem.
- **Logout**: Pengguna mengakhiri sesi dan keluar dari sistem.
- **Manage Employee**:
  - HRD/Manajer menambahkan karyawan baru ke sistem.
  - HRD/Manajer mengedit informasi karyawan (nama, posisi, dll.).
  - HRD/Manajer menonaktifkan/mengaktifkan akun karyawan.
- **Manage Roles**:
  - Admin membuat peran baru (misalnya, Kasir, Manajer Toko).
  - Admin mengedit deskripsi peran.
  - Admin menetapkan izin ke peran.
- **Assign Roles to Employee**:
  - HRD/Manajer menetapkan peran ke karyawan.
  - HRD/Manajer mencabut peran dari karyawan.
- **Manage Customer**:
  - Kasir/Manajer menambahkan pelanggan baru (member).
  - Kasir/Manajer mencari pelanggan berdasarkan nomor telepon/email.
  - Kasir/Manajer mengedit informasi pelanggan (alamat, tier).

#### 2. Organization & Store Management

- **Manage Company**:
  - Admin menambahkan perusahaan baru ke sistem (proses pendaftaran tenant).
  - Admin mengedit informasi perusahaan (nama, alamat, info kontak).
- **Manage Business Line**:
  - Admin/Manajer menambahkan lini bisnis baru ke perusahaan.
  - Admin/Manajer mengedit informasi lini bisnis (nama, deskripsi).
- **Manage Store**:
  - Admin/Manajer menambahkan toko baru ke lini bisnis.
  - Admin/Manajer mengedit informasi toko (nama, alamat, tipe toko).
  - Admin/Manajer mengatur hierarki toko (cabang, ranting).

#### 3. Product & Inventory Management

- **Manage Master Product**:
  - Manajer membuat master product baru (nama, kategori, deskripsi).
  - Manajer mengedit informasi master product.
- **Manage Store Product**:
  - Manajer menambahkan master product ke toko tertentu (create store product).
  - Manajer mengatur harga jual, harga beli, dan stok untuk store product.
  - Manajer mengedit informasi store product.
- **Manage Supplier**:
  - Manajer menambahkan supplier baru ke perusahaan.
  - Manajer mengedit informasi supplier (nama, kontak).
- **Manage Tax Rate**:
  - Admin menambahkan tarif pajak baru (misalnya, PPN 11%).
  - Admin mengedit/menonaktifkan tarif pajak.
- **Record Stock Movement**: Sistem secara otomatis mencatat setiap perubahan stok.

#### 4. Transaction Management

- **Create Transaction**:
  - Kasir memulai transaksi baru.
  - Kasir menambahkan item ke transaksi (memindai barcode atau mencari produk).
  - Kasir memberikan diskon (jika ada).
  - Kasir memilih metode pembayaran.
  - Sistem menghitung total, pajak, dan kembalian.
  - Kasir menyelesaikan transaksi dan mencetak struk.
- **View Transaction Detail**:
  - Manajer mencari transaksi berdasarkan ID atau tanggal.
  - Manajer melihat detail transaksi (item, total, kasir, pelanggan).
- **View Transaction List**:
  - Manajer melihat daftar transaksi dengan filter (tanggal, kasir, toko).
- **Refund Transaction**:
  - Manajer melakukan refund transaksi (jika diperlukan).

#### 5. Purchasing & Stock Transfer

- **Create Purchase Order**:
  - Pengelola membuat purchase order baru ke supplier.
  - Pengelola menambahkan item ke purchase order.
  - Pengelola mengirim purchase order ke supplier.
- **View Purchase Order List**:
  - Pengelola melihat daftar purchase order dengan filter (tanggal, supplier, status).
- **Receive Purchase Order**:
  - Pengelola menerima barang dari purchase order.
  - Sistem memperbarui stok produk.
- **Create Internal Stock Transfer**:
  - Pengelola membuat permintaan transfer stok antar toko.
- **Approve/Reject Stock Transfer Request**:
  - Manajer menyetujui atau menolak permintaan transfer stok.
- **Ship/Receive Stock Transfer**:
  - Karyawan mengirim barang dari toko asal.
  - Karyawan menerima barang di toko tujuan.
  - Sistem memperbarui stok di kedua toko.
- **Adjust Stock**:
  - Pengelola melakukan penyesuaian stok (karena kerusakan, kehilangan, dll.).
  - Sistem mencatat alasan penyesuaian stok.

#### 6. Shift Management

- **Manage Shift**:
  - Manajer membuat jadwal shift untuk karyawan.
  - Manajer mengedit jadwal shift.
- **Record Check-in/Check-out**:
  - Karyawan melakukan check-in saat mulai bekerja.
  - Karyawan melakukan check-out saat selesai bekerja.
  - Sistem mencatat waktu check-in/check-out dan menghitung jam kerja.

#### 7. Auditing

- **Record Activity**:
  - Sistem secara otomatis mencatat setiap aktivitas penting yang dilakukan oleh pengguna (login, edit produk, buat transaksi, dll.).
- **View Activity Log**:
  - Admin melihat log aktivitas untuk tujuan audit.

### Struktur Proyek

Struktur folder proyek ini mengikuti praktik terbaik Go untuk aplikasi backend:

```
/pos-backend
├── go.mod
├── go.sum                                  # Checksum dependensi proyek Go, dikelola oleh Go Modules
├── cmd/                                    # Berisi aplikasi yang dapat dieksekusi (executable applications)
│   └── api/                                # Direktori untuk aplikasi API utama
│       └── main.go                         # Titik masuk utama (entry point) aplikasi API, menginisialisasi server
├── internal/                               # Kode aplikasi privat (internal application code), tidak untuk diimpor oleh proyek eksternal
│   ├── api/                                # Lapisan API / Handler HTTP (presentation layer/controllers)
│   │   └── handlers/                       # Implementasi handler HTTP untuk setiap modul fungsional
│   │       ├── auth_handler.go             # Handler untuk autentikasi dan otorisasi pengguna (login, register)
│   │       ├── product_handler.go          # Handler untuk manajemen produk (CRUD produk)
│   │       ├── transaction_handler.go      # Handler untuk proses transaksi penjualan
│   │       └── ...                         # (handler untuk modul lain yang akan ditambahkan)
│   ├── core/                               # Logika bisnis inti / Use cases / Services (domain layer)
│   │   └── services/                       # Implementasi service yang berisi logika bisnis utama
│   │       ├── auth_service.go             # Service untuk autentikasi dan otorisasi
│   │       ├── product_service.go          # Service untuk manajemen produk
│   │       ├── transaction_service.go      # Service untuk transaksi penjualan
│   │       ├── employee_service.go         # (Kerangka awal) Service untuk manajemen karyawan
│   │       └── ...                         # (service untuk modul lain yang akan ditambahkan)
│   ├── data/                               # Lapisan akses data / Repositories / Storage (infrastructure layer)
│   │   └── postgres/                       # Implementasi repository spesifik untuk PostgreSQL
│   │       ├── user_repository.go          # Repository untuk entitas `user`
│   │       ├── company_repository.go       # Repository untuk entitas `company`
│   │       ├── store_repository.go         # Repository untuk entitas `store`
│   │       ├── employee_repository.go      # Repository untuk entitas `employee`
│   │       ├── business_line_repository.go # Repository untuk entitas `business_line`
│   │       ├── customer_repository.go      # Repository untuk entitas `customer`
│   │       ├── role_repository.go          # Repository untuk entitas `role`
│   │       ├── employee_role_repository.go # Repository untuk entitas `employee_role`
│   │       ├── supplier_repository.go      # Repository untuk entitas `supplier`
│   │       ├── store_product_repository.go # Repository untuk entitas `store_product`
│   │       ├── master_product_repository.go# Repository untuk entitas `master_product`
│   │       └── ...                         # (repository untuk modul lain)
│   ├── models/                             # Struktur data / Entitas (struct Go yang merepresentasikan tabel DB)
│   │   ├── activity_log.go                 # Model untuk log aktivitas pengguna
│   │   ├── applied_discount.go             # Model untuk diskon yang diterapkan pada transaksi/item
│   │   ├── business_line.go                # Model untuk lini bisnis perusahaan
│   │   ├── company.go                      # Model untuk data perusahaan (tenant)
│   │   ├── customer.go                     # Model untuk data pelanggan/member
│   │   ├── discount.go                     # Model untuk aturan diskon
│   │   ├── employee.go                     # Model untuk data karyawan
│   │   ├── employee_role.go                # Model untuk hubungan karyawan dan peran
│   │   ├── master_product.go               # Model untuk definisi produk master
│   │   ├── operational_expense.go          # Model untuk pengeluaran operasional
│   │   ├── permission.go                   # Model untuk izin/hak akses
│   │   ├── product.go                      # Model untuk produk (termasuk `store_product`)
│   │   ├── purchase_order.go               # Model untuk pesanan pembelian
│   │   ├── role.go                         # Model untuk peran pengguna
│   │   ├── shift.go                        # Model untuk jadwal shift karyawan
│   │   ├── stock_movement.go               # Model untuk catatan pergerakan stok
│   │   ├── stock_transfer.go               # Model untuk transfer stok internal (termasuk item transfer)
│   │   ├── store.go                        # Model untuk data toko
│   │   ├── supplier.go                     # Model untuk data pemasok
│   │   ├── tax_rate.go                     # Model untuk tarif pajak
│   │   ├── transaction.go                  # Model untuk transaksi penjualan (termasuk item transaksi)
│   │   └── user.go                         # Model untuk data pengguna dasar
│   ├── config/                             # Logika pemuatan dan struktur konfigurasi aplikasi
│   │   └── config.go                       # Definisi struktur konfigurasi aplikasi
│   ├── database/                           # Setup koneksi database
│   │   └── postgres.go                     # Fungsi untuk inisialisasi koneksi PostgreSQL
│   └── utils/                              # Fungsi-fungsi utilitas internal yang tidak spesifik modul
│       └── ...                             # (misalnya, helper untuk hashing password, validasi data, dll.)
├── pkg/                                    # Opsional: Kode library yang aman untuk digunakan oleh proyek eksternal
│   └── ...                                 # (misalnya, package untuk JWT, hashing, dll. jika tidak internal)
├── migrations/                             # File-file migrasi database (SQL scripts atau tool-specific files)
│   └── ...                                 # (misalnya, skrip SQL untuk membuat tabel, alter tabel)
├── tests/                                  # Direktori untuk semua test aplikasi
│   └── repository/                         # Test untuk lapisan repository
│       └── postgres/                       # Unit tests spesifik untuk repository PostgreSQL
│           ├── user_repository_test.go     # Test untuk UserRepository
│           ├── company_repository_test.go  # Test untuk CompanyRepository
│           ├── store_repository_test.go    # Test untuk StoreRepository
│           ├── activity_log_repository_test.go # Test untuk ActivityLogRepository
│           ├── applied_item_discount_repository_test.go # Test untuk AppliedItemDiscountRepository
│           ├── applied_transaction_discount_repository_test.go # Test untuk AppliedTransactionDiscountRepository
│           ├── business_line_repository_test.go # Test untuk BusinessLineRepository
│           ├── customer_repository_test.go # Test untuk CustomerRepository
│           ├── discount_repository_test.go   # Test untuk DiscountRepository
│           ├── employee_repository_test.go # Test untuk EmployeeRepository
│           ├── employee_role_repository_test.go # Test untuk EmployeeRoleRepository
│           ├── internal_stock_transfer_repository_test.go # Test untuk InternalStockTransferRepository
│           ├── master_product_repository_test.go # Test untuk MasterProductRepository
│           ├── operational_expense_repository_test.go # Test untuk OperationalExpenseRepository
│           ├── permission_repository_test.go  # Test untuk PermissionRepository
│           ├── purchase_order_repository_test.go # Test untuk PurchaseOrderRepository
│           ├── role_permission_repository_test.go # Test untuk RolePermissionRepository
│           ├── role_repository_test.go     # Test untuk RoleRepository
│           ├── shift_repository_test.go    # Test untuk ShiftRepository
│           ├── stock_movement_repository_test.go # Test untuk StockMovementRepository
│           ├── store_product_repository_test.go # Test untuk StoreProductRepository
│           ├── supplier_repository_test.go # Test untuk SupplierRepository
│           ├── tax_rate_repository_test.go    # Test untuk TaxRateRepository
│           ├── transaction_repository_test.go # Test untuk TransactionRepository
│           └── user_repository_test.go     # Test untuk UserRepository
│   └── services/                           # Unit tests spesifik untuk lapisan Service
│       └── auth_service_test.go            # Test untuk AuthService
└── Dockerfile                              # Konfigurasi untuk membangun container Docker
```

### Arsitektur Aplikasi

Aplikasi ini dibangun dengan pendekatan **Clean Architecture / Layered Architecture** untuk memastikan pemisahan tanggung jawab yang jelas, kemudahan pengujian, dan skalabilitas.

1.  **Lapisan Model (`internal/models`)**:

    - Berisi definisi struct Go yang merepresentasikan entitas data inti aplikasi. Ini adalah representasi langsung dari tabel database.
    - **Status: ✅ LENGKAP** (Semua model yang dibutuhkan telah dibuat dan terdokumentasi).

2.  **Lapisan Repository (`internal/data/postgres`)**:

    - Bertanggung jawab untuk interaksi langsung dengan database (PostgreSQL).
    - Mengimplementasikan interface repository yang didefinisikan di lapisan `data`.
    - Fokus pada operasi CRUD (Create, Read, Update, Delete) dan query data.
    - **Status: ✅ LENGKAP** (Semua repository yang didefinisikan dalam ERD telah selesai dan teruji).

3.  **Lapisan Service (`internal/core/services`)**:

    - Berisi logika bisnis inti aplikasi.
    - Mengorkestrasi panggilan ke satu atau lebih repository untuk menyelesaikan kasus penggunaan (use case) yang kompleks.
    - Melakukan validasi data, hashing password, manajemen transaksi database, dan penerapan aturan bisnis.
    - **Status: ⬜ SEDANG BERLANGSUNG** (Ini adalah fokus pengembangan saat ini. Implementasi penuh akan mencakup logika bisnis, validasi, dan orkestrasi repository).

4.  **Lapisan API / Handler (`internal/api/handlers`)**:
    - Bertanggung jawab untuk menerima permintaan HTTP, memanggil service yang sesuai, dan mengembalikan respons HTTP.
    - Tidak mengandung logika bisnis secara langsung, hanya bertindak sebagai antarmuka.
    - **Status: ⬜ BELUM DIMULAI** (Akan diimplementasikan setelah lapisan service cukup matang).

### Aspek Arsitektur Penting Lainnya

Selain lapisan inti di atas, beberapa aspek arsitektur penting lainnya yang akan dipertimbangkan dan diimplementasikan:

- **Dependency Injection (DI)**:
  - Akan digunakan untuk mengelola dependensi antar lapisan (misalnya, menyuntikkan repository ke service, dan service ke handler). Ini meningkatkan modularitas, kemudahan pengujian, dan pemeliharaan kode.
  - **Status: ⬜ TO-DO**
- **Penanganan Error Global**:
  - Strategi penanganan error yang konsisten akan diterapkan di seluruh lapisan aplikasi untuk memberikan respons error yang informatif dan aman kepada klien, tanpa mengekspos detail internal.
  - **Status: ⬜ TO-DO**
- **Logging dan Monitoring**:
  - Sistem logging yang komprehensif akan diintegrasikan untuk melacak aktivitas aplikasi, membantu debugging, audit, dan pemantauan performa.
  - **Status: ⬜ TO-DO**
- **Dokumentasi API**:
  - API akan didokumentasikan menggunakan standar seperti OpenAPI (Swagger) untuk memfasilitasi integrasi dengan klien frontend dan pihak ketiga.
  - **Status: ⬜ TO-DO**

### Kemajuan Proyek

Kita telah berhasil membangun fondasi yang kuat untuk aplikasi ini. Berikut adalah ringkasan kemajuan kita:

- **Setup Lingkungan Profesional**: Lingkungan pengembangan dan testing yang terpisah telah disiapkan, lengkap dengan koneksi database dan mekanisme pembersihan otomatis.
- **Pola Repository yang Solid**: Kita telah menetapkan dan mengikuti pola kerja yang andal untuk membangun repository dan unit test-nya.
- **Lapisan Model**: Semua model Go yang merepresentasikan entitas database telah dibuat dan terdokumentasi dengan baik.
  - **Status: ✅ LENGKAP**
- **Lapisan Repository**: Implementasi repository untuk interaksi dengan PostgreSQL telah selesai dan teruji.
  - **Status: ✅ LENGKAP**

#### Status Repository Saat Ini (`internal/data/postgres/`)

Berikut adalah daftar repository yang dibutuhkan dan status pengerjaannya:

**1. Modul: core_tenant_and_organization**

- `CompanyRepository`: Mengelola data perusahaan (tenant).
  - **Status: ✅ SELESAI & TERUJI**
- `BusinessLineRepository`: Mengelola lini bisnis yang dimiliki setiap perusahaan.
  - **Status: ✅ SELESAI & TERUJI**
- `StoreRepository`: Mengelola data toko, termasuk hierarkinya (pusat, cabang, ranting).
  - **Status: ✅ SELESAI & TERUJI**

**2. Modul: user_management**

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

**3. Modul: product_and_supplier_management**

- `SupplierRepository`: Mengelola data pemasok untuk setiap perusahaan.
  - **Status: ✅ SELESAI & TERUJI**
- `StoreProductRepository`: Mengelola data produk di setiap toko, termasuk harga, stok, barcode, dll.
  - **Status: ✅ SELESAI & TERUJI**
- `MasterProductRepository`: Mengelola definisi produk master yang tidak spesifik toko.
  - **Status: ✅ SELESAI & TERUJI**
- `TaxRateRepository`: Mengelola tarif pajak yang berlaku.
  - **Status: ✅ SELESAI & TERUJI**

**4. Modul: transaction_management**

- `TransactionRepository`: Mengelola data transaksi dan item-itemnya secara atomik.
  - **Status: ✅ SELESAI & TERUJI**

**5. Modul: shift_management**

- `ShiftRepository`: Mengelola jadwal shift, waktu check-in, dan check-out karyawan.
  - **Status: ✅ SELESAI & TERUJI**

**6. Modul: purchasing_and_stock_management**

- `PurchaseOrderRepository`: Mengelola pesanan pembelian ke supplier, termasuk item-itemnya.
  - **Status: ✅ SELESAI & TERUJI**
- `InternalStockTransferRepository`: Mengelola transfer barang antar toko dalam satu perusahaan.
  - **Status: ✅ SELESAI & TERUJI**
- `StockMovementRepository`: **Sangat Penting.** Mencatat setiap perubahan stok (penjualan, pembelian, transfer, penyesuaian) untuk tujuan audit dan pelacakan.
  - **Status: ✅ SELESAI & TERUJI**

**7. Modul: discount_management**

- `DiscountRepository`: Mengelola aturan diskon (misalnya, diskon 10% untuk kategori tertentu).
  - **Status: ✅ SELESAI & TERUJI**
- `AppliedTransactionDiscountRepository`: Mengelola tabel pivot untuk mencatat diskon apa yang diterapkan pada sebuah transaksi.
  - **Status: ✅ SELESAI & TERUJI**
- `AppliedItemDiscountRepository`: Mengelola tabel pivot untuk mencatat diskon apa yang diterapkan pada item transaksi.
  - **Status: ✅ SELESAI & TERUJI**

**8. Modul: auditing**

- `ActivityLogRepository`: Menyimpan log dari setiap tindakan penting yang dilakukan oleh pengguna.
  - **Status: ✅ SELESAI & TERUJI**

**9. Modul: expense_management**

- `OperationalExpenseRepository`: Mengelola data pengeluaran operasional.
  - **Status: ✅ SELESAI & TERUJI**
