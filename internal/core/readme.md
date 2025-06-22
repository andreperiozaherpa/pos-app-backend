# /Users/andre/Programing/aplikasi perusahaan/pos-backend/internal/core/readme.md

Direktori ini merepresentasikan **Lapisan Core** atau **Domain Layer** dari aplikasi. Ini adalah jantung dari aplikasi yang berisi semua logika bisnis inti (use cases) dan aturan domain.

Tujuan utama folder ini adalah:

- **Enkapsulasi Logika Bisnis**: Semua aturan dan proses bisnis yang kompleks diimplementasikan di sini.
- **Independensi dari Infrastruktur**: Kode dalam lapisan ini tidak boleh memiliki pengetahuan langsung tentang detail implementasi infrastruktur (seperti HTTP, database, atau sistem eksternal lainnya).
- **Ketergantungan pada Abstraksi**: Lapisan ini hanya boleh bergantung pada interface (abstraksi) yang didefinisikan di lapisan yang lebih rendah (misalnya, interface repository dari lapisan `data/`). Ini memungkinkan fleksibilitas dan kemudahan pengujian.

---

## Struktur Subdirektori

Direktori ini berisi subdirektori berikut:

- `services/`: Berisi file service yang mengimplementasikan logika bisnis tertentu.
  - **Status: ⬜ SEDANG BERLANGSUNG** (Ini adalah fokus pengembangan saat ini. Implementasi penuh akan mencakup logika bisnis, validasi, dan orkestrasi repository).

---

## Catatan

Lapisan `core` adalah lapisan yang paling stabil dan paling sedikit berubah dalam arsitektur aplikasi, karena ia hanya berurusan dengan aturan bisnis yang mendasar. Perubahan pada detail implementasi (misalnya, mengganti database) tidak akan memengaruhi lapisan ini.
