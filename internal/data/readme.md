# Dokumentasi Folder internal/data

Folder `internal/data` adalah tempat implementasi lapisan infrastruktur aplikasi yang berhubungan langsung dengan penyimpanan data, seperti database, cache, atau storage lainnya.

## Fungsi Utama

- Menyimpan implementasi repository yang berinteraksi dengan database atau storage.
- Memisahkan interface repository yang ada di `internal/core/repository` dengan implementasi aktual.
- Menjaga agar lapisan domain tetap bersih dan tidak tergantung pada teknologi penyimpanan data tertentu.

## Struktur Folder

Direkomendasikan untuk membuat subfolder berdasarkan jenis teknologi atau sistem penyimpanan data yang digunakan, misalnya:

- `/internal/data/postgres/` untuk implementasi repository PostgreSQL.
- `/internal/data/mongo/` untuk implementasi repository MongoDB (jika digunakan).
- `/internal/data/redis/` untuk implementasi cache Redis (jika digunakan).

## Panduan Implementasi

- Implementasi repository harus memenuhi kontrak interface yang sudah didefinisikan di `internal/core/repository`.
- Jangan mengimpor package domain di dalam implementasi repository agar menjaga independensi domain.
- Gunakan dependency injection untuk menghubungkan implementasi repository dengan service pada lapisan application.

## Catatan

- Folder ini bisa berkembang seiring bertambahnya teknologi penyimpanan yang digunakan dalam aplikasi.
- Jika ada service lain yang berhubungan dengan data storage selain repository, bisa dipertimbangkan untuk ditempatkan di sini juga.
