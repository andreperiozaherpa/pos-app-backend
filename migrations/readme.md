# /Users/andre/Programing/aplikasi perusahaan/pos-backend/migrations/readme.md

Direktori ini berisi semua file migrasi database. Migrasi adalah perubahan terstruktur pada skema database (dan kadang-kadang data) yang diterapkan secara berurutan.

Tujuan utama folder ini adalah untuk:

- **Manajemen Skema Terkontrol**: Memastikan bahwa perubahan skema database diterapkan secara konsisten di semua lingkungan (development, staging, production).
- **Pelacakan Perubahan**: Menyediakan riwayat yang jelas tentang bagaimana skema database telah berevolusi dari waktu ke waktu.
- **Kolaborasi Tim**: Memungkinkan beberapa pengembang untuk bekerja pada perubahan database secara bersamaan tanpa konflik yang signifikan.
- **Rollback**: Memfasilitasi kemampuan untuk mengembalikan perubahan skema jika terjadi masalah.

---

## Jenis File Migrasi

File migrasi di sini biasanya akan berupa skrip SQL (`.sql`) atau file yang spesifik untuk alat migrasi yang digunakan (misalnya, `goose`, `migrate`, `golang-migrate`).

Berikut adalah contoh jenis-jenis migrasi yang akan ada di folder ini:

- `YYYYMMDDHHMMSS_initial_schema.sql`: Migrasi awal untuk membuat semua tabel dan indeks dasar.
  - **Status: ⬜ TO-DO**
- `YYYYMMDDHHMMSS_add_new_feature_table.sql`: Migrasi untuk menambahkan tabel baru yang diperlukan oleh fitur tertentu.
  - **Status: ⬜ TO-DO**
- `YYYYMMDDHHMMSS_alter_existing_table.sql`: Migrasi untuk mengubah skema tabel yang sudah ada (misalnya, menambahkan kolom, mengubah tipe data).
  - **Status: ⬜ TO-DO**
- `YYYYMMDDHHMMSS_data_migration_for_x.sql`: Migrasi untuk memanipulasi data yang sudah ada di database (misalnya, mengisi kolom baru dengan nilai default).
  - **Status: ⬜ TO-DO**

---

## Catatan

Setiap file migrasi harus bersifat _idempotent_ (dapat dijalankan berkali-kali tanpa efek samping yang tidak diinginkan) dan _atomic_ (berhasil sepenuhnya atau gagal sepenuhnya). Penamaan file migrasi biasanya mengikuti format timestamp untuk memastikan urutan eksekusi yang benar.
