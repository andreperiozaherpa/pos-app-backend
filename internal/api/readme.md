# /internal/api

Direktori ini merepresentasikan **Lapisan API** (API Layer / Presentation Layer) dari aplikasi.  
Semua permintaan eksternal ke backend masuk lewat lapisan ini.

---

## 📦 Tanggung Jawab Utama

- **Mendefinisikan Rute API:**  
  Mengatur endpoint-endpoint REST (misal: `/login`, `/products`, dst).
- **Menangani Permintaan/Respons HTTP:**  
  Parsing request (body, query, path), format response (JSON).
- **Orkestrasi Middleware:**  
  Autentikasi, otorisasi, logging, CORS, dll.
- **Delegasi ke Service/Usecase:**  
  Semua logika bisnis harus di-_delegasi_ ke lapisan core/service/usecase.

---

## 🗂️ Struktur Subdirektori & Status

| Folder/File   | Deskripsi                                         | Status         |
| ------------- | ------------------------------------------------- | -------------- |
| `handlers/`   | Handler HTTP per modul fungsional                 | ⬜ BELUM MULAI |
| `middleware/` | Fungsi middleware reusable (auth, log, CORS, dll) | ⬜ TO-DO       |
| `router.go`   | Daftar semua route & inisialisasi router utama    | ⬜ TO-DO       |

---

## 🏗️ Contoh Struktur Ideal

```
api/
├── handlers/
│   ├── user_handler.go
│   ├── product_handler.go
│   └── ...
├── middleware/
│   ├── auth.go
│   ├── logger.go
│   └── ...
├── router.go
└── readme.md
```

---

## 🚩 Best Practice

- **Jaga API Layer Tetap Tipis:**  
  Semua business logic, validasi berat, dan proses data harus di service/usecase.  
  Handler hanya parsing-request, delegasi, dan kirim response.

- **Uniform Response:**  
  Pakai pola response JSON seragam, misal:
  ```json
  {
    "success": true,
    "data": {},
    "message": "ok"
  }
  ```
- **Error Handling:**  
  Semua error harus di-handle dan response tidak bocor detail internal (error DB, dsb).

---

**Terakhir update:** [isi tanggal & inisial updater]

---
