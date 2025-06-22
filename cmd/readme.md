# /Users/andre/Programing/aplikasi perusahaan/pos-backend/cmd/readme.md

Direktori ini berisi semua aplikasi utama yang dapat dieksekusi (executable applications) dari proyek. Setiap subdirektori di dalam `cmd/` harus merepresentasikan satu aplikasi yang dapat dikompilasi menjadi sebuah executable.

Tujuan utama folder ini adalah untuk:

- **Titik Masuk Aplikasi**: Menyediakan titik masuk (`main.go`) untuk setiap layanan atau alat yang akan dijalankan.
- **Pemisahan Tanggung Jawab**: Memisahkan kode yang bertanggung jawab untuk inisialisasi aplikasi dan memulai server dari logika bisnis inti.
- **Fleksibilitas Deployment**: Memungkinkan deployment aplikasi yang berbeda secara independen (misalnya, server API, worker antrean, alat CLI).

---

## Daftar Aplikasi yang Dapat Dieksekusi

Berikut adalah daftar aplikasi yang akan ada di direktori ini beserta fungsinya:

- `api/`: Berisi kode untuk server API utama aplikasi POS. Ini akan menjadi executable yang melayani permintaan HTTP dari klien.
  - **Status: ✅ SELESAI & TERUJI** (File `main.go` sudah ada dan siap untuk inisialisasi).

---

## Catatan

Setiap subdirektori di `cmd/` harus berisi file `main.go` dan tidak boleh ada kode lain di dalamnya selain yang diperlukan untuk menginisialisasi dan menjalankan aplikasi. Logika bisnis dan fungsionalitas lainnya harus berada di direktori `internal/`.
