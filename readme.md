# pos-app/backend

Ini adalah direktori root untuk aplikasi Backend POS.

### Visi Aplikasi

Aplikasi ini adalah sistem Point of Sale (POS) multi-tenant yang dirancang untuk mendukung berbagai jenis perusahaan dengan struktur bisnis yang kompleks.

### Use Cases (Kasus Penggunaan)

Aplikasi POS ini dirancang untuk menangani kebutuhan sistem penjualan (Point of Sale) multi-tenant dengan berbagai skenario bisnis berikut:

- **Auth & User Management:**  
  Registrasi, login, manajemen user (karyawan & customer/member), pengelolaan role & permission, reset password, dua faktor autentikasi, serta pencatatan aktivitas user.

- **Organization & Store Management:**  
  Pengelolaan perusahaan multi-lini (pusat, cabang, ranting), pengelolaan toko, struktur organisasi, serta manajemen karyawan per toko atau perusahaan.

- **Product & Inventory Management:**  
  Manajemen data master produk, produk per toko, harga beli/jual/grosir, stok, suplier, barcode, serta mutasi dan histori stok.

- **Transaction Management:**  
  Penanganan transaksi penjualan: pencatatan, pemberian diskon, penambahan pajak, pencetakan struk/barcode, refund, audit trail, serta pencatatan pembayaran dan pelacakan status transaksi.

- **Purchasing & Stock Transfer:**  
  Pembuatan purchase order, penerimaan barang, transfer stok antar toko/cabang, serta pencatatan mutasi stok masuk/keluar.

- **Shift Management:**  
  Pengaturan dan pencatatan jadwal shift karyawan, absensi, cuti, serta riwayat kehadiran.

- **Auditing (Activity Log):**  
  Pencatatan seluruh aktivitas penting user dan sistem untuk keperluan audit dan keamanan.

- **Import/Export Data:**  
  Fitur import/export data massal untuk master data (produk, customer, suplier) maupun transaksi/laporan.

- **Reporting (Laporan/Export):**  
  Penyediaan laporan penjualan, stok, laba rugi, aktivitas karyawan, dsb. dalam berbagai format (Excel/PDF).

- **File Storage (Upload Gambar/Dokumen):**  
  Penyimpanan dan manajemen file seperti gambar produk, dokumen pendukung, dan file hasil export/import.

- **RBAC (Role Based Access Control):**  
  Pengelolaan hak akses aplikasi berbasis role dan permission secara fleksibel.

- **Notification (Email/SMS/WA):**  
  Fitur notifikasi ke user untuk berbagai event penting, baik internal (in-app) maupun eksternal (email, WhatsApp, SMS).

---

### Checklist Progres Use Case Utama

| Use Case                              | Status     | Catatan                                         |
| ------------------------------------- | ---------- | ----------------------------------------------- |
| Auth & User Management                | ✅ Selesai | Service, model, repository, interface           |
| Organization & Store Management       | ✅ Selesai | Service, model, repository, interface           |
| Product & Inventory Management        | ✅ Selesai | Service, model, repository, interface           |
| Transaction Management                | ✅ Selesai | Service, model, repository, interface           |
| Purchasing & Stock Transfer           | ✅ Selesai | Service, model, repository, interface           |
| Shift Management                      | ✅ Selesai | Service, model, repository, interface           |
| Auditing (Activity Log)               | ✅ Selesai | Service, model, repository, interface           |
| Import/Export Data                    | ✅ Selesai | Service, model, repository, interface, opsional |
| Reporting (Laporan/Export)            | ✅ Selesai | Service, model, repository, interface, opsional |
| File Storage (Upload Gambar/Dokumen)  | ✅ Selesai | Service, model, repository, interface, opsional |
| RBAC (Role Based Access Control)      | ✅ Selesai | Service, model, repository, interface, opsional |
| Notification (Email/SMS/WA, Opsional) | ✅ Selesai | Service, model, repository, interface, opsional |

---

### Arsitektur Aplikasi

Aplikasi ini menerapkan pola **Clean Architecture** dengan pemisahan yang jelas antara domain/business logic (core), data access (repository), application service, serta API handler/controller.  
Struktur ini memudahkan maintainability, pengujian, dan scale-up aplikasi.

**Lapisan utama arsitektur:**

- **Core Domain/Business Logic:**  
  Seluruh interface service (usecase) & repository didefinisikan tanpa ketergantungan pada teknologi/infrastruktur spesifik.
- **Model (Entity & Struct):**  
  Seluruh entitas utama & supporting (advance model) yang dibutuhkan dalam bisnis.
- **Infrastruktur Data/Repository:**  
  Implementasi akses data nyata (database/ORM, dsb) yang memenuhi kontrak repository.
- **Application Service (Usecase):**  
  Implementasi logic nyata (service) yang mengatur workflow aplikasi.
- **API Handler/Controller:**  
  Handler untuk request HTTP/gRPC, dependency ke application service.
- **Integration Test & Mock:**  
  Mock dan test untuk memastikan reliability dan robustness aplikasi.
- **Dependency Injection:**  
  Mekanisme wiring dependency untuk menjaga loose coupling.
- **Penanganan Error Global, Logging & Monitoring:**  
  Menjamin reliability, traceability, serta observability aplikasi.

---

Checklist progres dan usecase akan terus diupdate mengikuti kebutuhan bisnis, penambahan fitur, serta proses pengembangan berikutnya.

### Arsitektur Aplikasi

| Komponen/Lapisan              | Status     | Catatan                                                            |
| ----------------------------- | ---------- | ------------------------------------------------------------------ |
| Core Domain/Business Logic    | ✅ Selesai | `/internal/core` (service & repository interface sudah lengkap)    |
| Model (Entity & Struct)       | ✅ Selesai | `/internal/models` semua entity utama & advanced sudah ada         |
| Infrastruktur Data/Repository | ⬜ Belum   | Implementasi repository infra, contoh: Gorm, Postgres, dsb (TO-DO) |
| Application Service (Usecase) | ⬜ Belum   | Implementasi service nyata, wiring DI (TO-DO)                      |
| API Handler/Controller        | ⬜ Belum   | Handler REST/HTTP/gRPC, dependency ke service (TO-DO)              |
| Integration Test & Mock       | ⬜ Belum   | `/internal/mock` & `/internal/test` (TO-DO)                        |
| Dependency Injection          | ⬜ Belum   | DI container/manual/wiring (TO-DO)                                 |
| Penanganan Error Global       | ⬜ Belum   | Middleware/error handler API (TO-DO)                               |
| Logging & Monitoring          | ⬜ Belum   | Setup logging (zap/logrus), APM/monitoring (TO-DO)                 |
| Dokumentasi API (OpenAPI)     | ⬜ Belum   | Swagger, Postman, dsb (TO-DO)                                      |
