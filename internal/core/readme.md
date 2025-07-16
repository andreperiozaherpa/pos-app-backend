# internal/core

Direktori ini adalah jantung dari arsitektur _Clean Architecture_ aplikasi, yang berisi definisi kontrak (interface) logika bisnis utama (_usecase/service layer_) dan akses data (_repository layer_) untuk setiap domain utama.

---

## 📂 Struktur Direktori

| Folder      | Status     | Keterangan                                         |
| ----------- | ---------- | -------------------------------------------------- |
| /service    | ✅ Selesai | Interface usecase/business logic (sudah lengkap)   |
| /repository | ✅ Selesai | Interface repository/data access (sudah lengkap)   |
| /dto        | ⬜ Belum   | Data transfer object (opsional, belum ada)         |
| /mock       | ⬜ Belum   | Mock interface untuk testing (opsional, belum ada) |
| /test       | ⬜ Belum   | Kode unit/integrasi testing (opsional, belum ada)  |

---

## 📝 Panduan Pengembangan

1. **Definisikan interface baru** di folder `/service` (untuk business logic) dan `/repository` (untuk akses data) untuk setiap entitas/domain baru.
2. **Update README di tiap subfolder** sebagai checklist progres (status interface/method/service).
3. **Implementasikan interface** pada layer aplikasi/infrastruktur (di luar core).
4. **Pisahkan logika domain** dari detail database atau teknologi lain, agar mudah di-test dan scalable.
5. **Selalu update dokumentasi** jika ada entitas, method, atau workflow baru.

---

## 📑 Dokumentasi & Navigasi

- [Service Layer (`/service`)](./service/readme.md):  
  Daftar lengkap interface logika bisnis (dengan checklist method utama & advanced).
- [Repository Layer (`/repository`)](./repository/readme.md):  
  Daftar lengkap interface repository (dengan checklist method utama & advanced).

---

## 🚀 Catatan

- Tidak ada implementasi business logic/infrastruktur di sini, hanya kontrak (_pure interface_).
- File struct model didefinisikan di `internal/models`.
- Untuk testing/mock bisa gunakan `/mock` atau tools GoMock, dst.

---
