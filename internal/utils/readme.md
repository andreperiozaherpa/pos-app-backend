# /Users/andre/Programing/aplikasi perusahaan/pos-backend/internal/utils/readme.md

Direktori ini berisi kumpulan fungsi utilitas internal yang bersifat generik dan dapat digunakan kembali di berbagai bagian aplikasi. Fungsi-fungsi di sini tidak spesifik untuk modul bisnis tertentu (misalnya, manajemen pengguna atau produk), melainkan menyediakan fungsionalitas umum yang dibutuhkan di banyak tempat.

Tujuan utama folder ini adalah untuk:

- **Meningkatkan Reusabilitas**: Menghindari duplikasi kode dengan menyediakan fungsi yang dapat dipanggil dari mana saja.
- **Memisahkan Tanggung Jawab**: Menjaga logika bisnis tetap bersih dengan memindahkan fungsionalitas pendukung ke tempat yang tepat.
- **Mempermudah Pemeliharaan**: Perubahan pada utilitas hanya perlu dilakukan di satu tempat.

---

## Daftar Utilitas

Berikut adalah daftar file utilitas yang ada di direktori ini beserta fungsinya:

- `password.go`: Menyediakan fungsi untuk mengelola password, seperti hashing password menggunakan bcrypt dan memverifikasi password terhadap hash-nya.
  - **Status: ✅ SELESAI & TERUJI**
- `jwt.go`: Menyediakan fungsi untuk mengelola JSON Web Tokens (JWT), seperti membuat token JWT baru dan memvalidasi token yang sudah ada.
  - **Status: ✅ SELESAI & TERUJI**
