# /Users/andre/Programing/aplikasi perusahaan/pos-backend/internal/api/readme.md

Direktori ini merepresentasikan **Lapisan API** atau **Lapisan Presentasi** dari aplikasi. Ini adalah titik masuk utama untuk semua permintaan eksternal ke backend dan bertanggung jawab untuk menangani interaksi HTTP.

Tujuan utama folder ini adalah untuk:

- **Mendefinisikan Rute API**: Mengatur endpoint-endpoint yang tersedia (misalnya, `/login`, `/products`).
- **Menangani Permintaan/Respons HTTP**: Mem-parsing data dari permintaan (request body, query params) dan memformat data untuk respons (biasanya dalam format JSON).
- **Orkestrasi Middleware**: Menerapkan middleware untuk fungsionalitas lintas-sisi seperti autentikasi, otorisasi, logging, dan CORS.
- **Delegasi ke Lapisan Service**: Memanggil metode yang sesuai di lapisan service untuk menjalankan logika bisnis, dan kemudian mengembalikan hasilnya kepada klien.

---

## Struktur Subdirektori

Direktori ini diatur ke dalam subdirektori berikut untuk memisahkan tanggung jawab:

- `handlers/`: Berisi implementasi handler HTTP untuk setiap modul fungsional.
  - **Status: ⬜ BELUM DIMULAI** (Akan diimplementasikan setelah lapisan service cukup matang).
- `middleware/`: Berisi fungsi-fungsi middleware HTTP yang dapat digunakan kembali (misalnya, untuk autentikasi, logging).
  - **Status: ⬜ TO-DO**
- `router.go` (atau file serupa): Berisi definisi semua rute API dan menghubungkannya ke handler yang sesuai.
  - **Status: ⬜ TO-DO**

---

## Catatan

Lapisan API harus tetap "tipis" (thin). Logika bisnis yang kompleks tidak boleh berada di sini; sebaliknya, logika tersebut harus didelegasikan ke lapisan `core/services`.
