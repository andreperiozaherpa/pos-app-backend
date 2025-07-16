# /Users/andre/Programing/aplikasi perusahaan/pos-backend/internal/config/readme.md

Direktori ini berisi kode dan definisi struktur untuk memuat dan mengelola konfigurasi aplikasi. Tujuannya adalah untuk menyediakan cara yang terpusat dan terstruktur untuk mengakses pengaturan aplikasi, seperti kredensial database, kunci API, port server, dan parameter lingkungan lainnya.

Manfaat utama folder ini adalah:

- **Pemisahan Konfigurasi**: Memisahkan pengaturan yang dapat berubah dari kode sumber, memungkinkan perubahan konfigurasi tanpa perlu mengkompilasi ulang aplikasi.
- **Manajemen Lingkungan**: Memudahkan pengelolaan konfigurasi yang berbeda untuk lingkungan pengembangan, pengujian, dan produksi.
- **Keamanan**: Memungkinkan penggunaan variabel lingkungan atau sistem manajemen rahasia untuk kredensial sensitif.

---

## Daftar File Konfigurasi

| File        | Fungsi                                                                        | Status              |
| ----------- | ----------------------------------------------------------------------------- | ------------------- |
| `config.go` | Definisi struct `AppConfig`, fungsi `LoadConfig()` untuk load env & file .env | ✅ Selesai & Teruji |

---

## Contoh .env

Berikut contoh file `.env` yang biasa digunakan di aplikasi ini:

```env
APP_NAME=pos-backend
APP_PORT=8080
DATABASE_URL=postgres://user:password@localhost:5432/dbname
JWT_SECRET=your-super-secret-key
ENVIRONMENT=development
```

---
