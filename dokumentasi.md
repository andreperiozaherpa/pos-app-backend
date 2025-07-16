# Dokumentasi Proyek Backend POS

Dokumen ini menyediakan tinjauan lengkap tentang aplikasi backend Point of Sale (POS), termasuk visi, kasus penggunaan, arsitektur, struktur proyek, dan status kemajuan.

## Daftar Isi

- [Visi Aplikasi](#visi-aplikasi)
- [Kasus Penggunaan (Use Cases)](#kasus-penggunaan-use-cases)
  - [Auth & User Management](#1-auth--user-management)
  - [Organization & Store Management](#2-organization--store-management)
  - ...
- [Struktur Proyek](#struktur-proyek)
- [Status Kemajuan Proyek](#status-kemajuan-proyek)

---

## Teknologi yang Digunakan

- Bahasa Pemrograman: Go 1.22+
- Database: PostgreSQL 15
- Migrasi: Goose / SQLC (atau alat lain yang kamu pakai)
- Routing: Chi / Gin (sebutkan jika sudah ada)
- Auth: JWT (jika pakai)
- Env Management: godotenv
- Logger: Zerolog / Logrus (sebutkan jika ada)
- Testing: `testing`, `testify`, dsb.

---

## Visi Aplikasi

Aplikasi ini adalah sistem Point of Sale (POS) multi-tenant yang dirancang untuk mendukung berbagai jenis perusahaan dengan struktur bisnis yang kompleks. Tujuannya adalah menyediakan platform yang fleksibel, andal, dan dapat diskalakan untuk mengelola operasi penjualan, inventaris, karyawan, dan pelanggan.

---

## Arsitektur Umum

Aplikasi ini mengikuti pendekatan Clean Architecture dengan pembagian tanggung jawab sebagai berikut:

- `internal/api`: HTTP handler dan route
- `internal/core`: logika bisnis dan interface use case
- `internal/data`: akses database
- `internal/models`: entity/domain model

Semua dependensi mengarah ke `core`, menjaga isolasi logika bisnis.

---

## Kasus Penggunaan (Use Cases)

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

---

## Struktur Proyek

Untuk detail lengkap mengenai struktur folder dan file dalam proyek ini, silakan merujuk ke dokumen terpisah: [**Struktur Proyek**](./struktur.md).

---

## Status Kemajuan Proyek

Berikut adalah status kemajuan untuk setiap komponen utama aplikasi:

### Lapisan Aplikasi

- **Lapisan Model (`internal/models`)**:
  - **Status: ✅ SELESAI** (Semua model yang dibutuhkan telah dibuat dan terdokumentasi).
- **Lapisan Repository (`internal/data/postgres`)**:
  - **Status: ✅ SELESAI & TERUJI** (Semua repository untuk PostgreSQL telah diimplementasikan dan diuji).
- **Lapisan Service (`internal/core/services`)**:
  - **Status: ⬜ SEDANG BERLANGSUNG** (Fokus pengembangan saat ini. `AuthService`, `UserService`, `EmployeeService`, dan `CustomerService` telah selesai dan teruji).
- **Lapisan API (`internal/api`)**:
  - **Status: ⬜ TO-DO** (Akan diimplementasikan setelah lapisan service cukup matang).

### Komponen Pendukung

- **Konfigurasi (`internal/config`)**:
  - **Status: ✅ SELESAI & TERUJI**
- **Koneksi Database (`internal/database`)**:
  - **Status: ✅ SELESAI & TERUJI**
- **Utilitas (`internal/utils`)**:
  - **Status: ✅ SELESAI & TERUJI**

### Infrastruktur & Lainnya

- **Migrasi Database (`migrations`)**:
  - **Status: ⬜ TO-DO**
- **Testing**:
  - Repository: **✅ SELESAI & TERUJI**
  - Services: **⬜ SEDANG BERLANGSUNG**
  - API: **⬜ TO-DO**
- **Dependency Injection (DI)**:
  - **Status: ⬜ TO-DO**
- **Penanganan Error Global**:
  - **Status: ⬜ TO-DO**
- **Logging dan Monitoring**:
  - **Status: ⬜ TO-DO**
- **Dokumentasi API (Swagger/OpenAPI)**:
  - **Status: ⬜ TO-DO**
