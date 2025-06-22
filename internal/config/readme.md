# /Users/andre/Programing/aplikasi perusahaan/pos-backend/internal/config/readme.md

Direktori ini berisi kode dan definisi struktur untuk memuat dan mengelola konfigurasi aplikasi. Tujuannya adalah untuk menyediakan cara yang terpusat dan terstruktur untuk mengakses pengaturan aplikasi, seperti kredensial database, kunci API, port server, dan parameter lingkungan lainnya.

Manfaat utama folder ini adalah:

- **Pemisahan Konfigurasi**: Memisahkan pengaturan yang dapat berubah dari kode sumber, memungkinkan perubahan konfigurasi tanpa perlu mengkompilasi ulang aplikasi.
- **Manajemen Lingkungan**: Memudahkan pengelolaan konfigurasi yang berbeda untuk lingkungan pengembangan, pengujian, dan produksi.
- **Keamanan**: Memungkinkan penggunaan variabel lingkungan atau sistem manajemen rahasia untuk kredensial sensitif.

---

## Daftar File Konfigurasi

Berikut adalah daftar file konfigurasi yang ada di direktori ini beserta fungsinya:

- `config.go`: Mendefinisikan struktur `AppConfig` yang menampung semua pengaturan aplikasi, serta fungsi `LoadConfig()` yang bertanggung jawab untuk memuat nilai-nilai konfigurasi dari variabel lingkungan dan/atau file `.env`.
  - **Status: ✅ SELESAI & TERUJI**
