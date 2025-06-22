# /Users/andre/Programing/aplikasi perusahaan/pos-backend/internal/data/readme.md

Direktori ini merepresentasikan **Lapisan Akses Data** atau **Infrastructure Layer** dari aplikasi. Ini adalah tempat di mana implementasi konkret dari mekanisme persistensi data berada.

Tanggung jawab utama folder ini adalah:

- **Interaksi Database**: Berisi kode yang secara langsung berinteraksi dengan database (misalnya, PostgreSQL, MySQL, SQLite).
- **Implementasi Repository**: Mengimplementasikan interface repository yang didefinisikan di lapisan `core/` atau di sini, menyediakan metode untuk operasi CRUD (Create, Read, Update, Delete) dan query data.
- **Abstraksi Detail Database**: Menyembunyikan detail spesifik database dari lapisan di atasnya (seperti lapisan `core/services`), sehingga logika bisnis tidak terikat pada jenis database tertentu.

---

## Struktur Subdirektori

- `postgres/`: Berisi implementasi repository spesifik untuk PostgreSQL.
  - **Status: ✅ SELESAI & TERUJI** (Semua repository untuk PostgreSQL telah diimplementasikan dan diuji).
- `mysql/`: Berisi implementasi repository spesifik untuk MySQL (jika diperlukan di masa depan).
  - **Status: ⬜ TO-DO**
- `sqlite/`: Berisi implementasi repository spesifik untuk SQLite (jika diperlukan di masa depan).
  - **Status: ⬜ TO-DO**

---

## Catatan

Lapisan `data` adalah lapisan yang paling dekat dengan infrastruktur persistensi. Perubahan pada teknologi database (misalnya, dari PostgreSQL ke MySQL) hanya akan memengaruhi implementasi di dalam subdirektori yang relevan di sini, tanpa memengaruhi lapisan `core/services` atau `api/handlers`.
