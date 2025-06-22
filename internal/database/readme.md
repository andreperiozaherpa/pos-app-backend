# /Users/andre/Programing/aplikasi perusahaan/pos-backend/internal/database/readme.md

Direktori ini bertanggung jawab untuk menyiapkan dan mengelola koneksi ke sistem manajemen database (DBMS) yang digunakan oleh aplikasi. Ini adalah bagian dari lapisan infrastruktur yang memastikan aplikasi dapat berkomunikasi dengan database.

Tujuan utama folder ini adalah:

- **Inisialisasi Koneksi**: Menyediakan fungsi untuk membuat dan menguji koneksi ke database.
- **Manajemen Pool Koneksi**: Mengelola pool koneksi database untuk efisiensi dan performa.
- **Abstraksi Driver**: Meskipun spesifik untuk satu jenis database, tujuannya adalah menyediakan objek koneksi yang dapat digunakan oleh lapisan `data` tanpa perlu mengetahui detail driver.

---

## Daftar File Koneksi Database

Berikut adalah daftar file koneksi database yang ada di direktori ini beserta fungsinya:

- `postgres.go`: Mendefinisikan fungsi `ConnectDB()`, yang membuat koneksi ke database PostgreSQL menggunakan pengaturan konfigurasi dari `internal/config`.
  - **Status: ✅ SELESAI & TERUJI**

---

## Catatan

File-file di direktori ini sangat spesifik untuk jenis database yang digunakan. Jika aplikasi perlu mendukung database lain (misalnya, MySQL), file baru (misalnya, `mysql.go`) akan ditambahkan di sini.
