# /Users/andre/Programing/aplikasi perusahaan/pos-backend/internal/api/handlers/readme.md

Direktori ini berisi implementasi dari **Lapisan API / Handler HTTP**. Ini adalah titik masuk utama untuk semua permintaan eksternal ke backend aplikasi.

Setiap file handler (misalnya, `auth_handler.go`) bertanggung jawab untuk menangani permintaan HTTP yang terkait dengan modul atau sumber daya tertentu (misalnya, autentikasi, produk, transaksi).

Tanggung jawab utama handler meliputi:

- Menerima permintaan HTTP (parsing request body, query parameters, path variables).
- Memanggil metode service yang sesuai untuk menjalankan logika bisnis.
- Mengelola respons HTTP (mengatur status kode, mengembalikan data JSON, menangani error).
- Tidak mengandung logika bisnis secara langsung; hanya bertindak sebagai antarmuka tipis ke lapisan service.

Contoh file:

- `auth_handler.go`: Menangani endpoint terkait autentikasi (login, register, dll.).
- `product_handler.go`: Menangani endpoint terkait produk (daftar produk, buat produk, dll.).
- `transaction_handler.go`: Menangani endpoint terkait transaksi (buat transaksi, lihat transaksi, dll.).

---

## Daftar Handler yang Akan Diimplementasikan

Berikut adalah daftar handler yang akan diimplementasikan, yang masing-masing akan berinteraksi dengan service yang sesuai:

- `auth_handler.go`: Menangani endpoint terkait autentikasi (login, logout, refresh token).
  - **Status: ⬜ TO-DO**
- `user_handler.go`: Menangani endpoint terkait manajemen pengguna umum (misalnya, melihat/memperbarui profil pengguna).
  - **Status: ⬜ TO-DO**
- `employee_handler.go`: Menangani endpoint terkait manajemen karyawan (CRUD karyawan, penetapan peran).
  - **Status: ⬜ TO-DO**
- `customer_handler.go`: Menangani endpoint terkait manajemen pelanggan/member (CRUD pelanggan, poin).
  - **Status: ⬜ TO-DO**
- `role_handler.go`: Menangani endpoint terkait manajemen peran (CRUD peran, penetapan izin ke peran).
  - **Status: ⬜ TO-DO**
- `permission_handler.go`: Menangani endpoint terkait manajemen izin (daftar izin).
  - **Status: ⬜ TO-DO**
- `company_handler.go`: Menangani endpoint terkait manajemen perusahaan (CRUD perusahaan).
  - **Status: ⬜ TO-DO**
- `business_line_handler.go`: Menangani endpoint terkait manajemen lini bisnis (CRUD lini bisnis).
  - **Status: ⬜ TO-DO**
- `store_handler.go`: Menangani endpoint terkait manajemen toko (CRUD toko).
  - **Status: ⬜ TO-DO**
- `master_product_handler.go`: Menangani endpoint terkait manajemen master produk (CRUD master produk).
  - **Status: ⬜ TO-DO**
- `store_product_handler.go`: Menangani endpoint terkait manajemen produk di toko (CRUD produk toko, penyesuaian stok).
  - **Status: ⬜ TO-DO**
- `supplier_handler.go`: Menangani endpoint terkait manajemen supplier (CRUD supplier).
  - **Status: ⬜ TO-DO**
- `tax_rate_handler.go`: Menangani endpoint terkait manajemen tarif pajak (CRUD tarif pajak).
  - **Status: ⬜ TO-DO**
- `transaction_handler.go`: Menangani endpoint terkait transaksi penjualan (membuat transaksi, melihat detail/daftar transaksi, refund).
  - **Status: ⬜ TO-DO**
- `shift_handler.go`: Menangani endpoint terkait manajemen shift (membuat shift, check-in/out).
  - **Status: ⬜ TO-DO**
- `purchase_order_handler.go`: Menangani endpoint terkait manajemen pesanan pembelian (membuat PO, menerima PO).
  - **Status: ⬜ TO-DO**
- `internal_stock_transfer_handler.go`: Menangani endpoint terkait transfer stok internal (membuat, menyetujui, mengirim, menerima transfer).
  - **Status: ⬜ TO-DO**
- `activity_log_handler.go`: Menangani endpoint terkait log aktivitas (melihat log).
  - **Status: ⬜ TO-DO**
- `discount_handler.go`: Menangani endpoint terkait manajemen diskon (CRUD diskon, penerapan diskon).
  - **Status: ⬜ TO-DO**
- `operational_expense_handler.go`: Menangani endpoint terkait manajemen pengeluaran operasional (CRUD pengeluaran).
  - **Status: ⬜ TO-DO**
