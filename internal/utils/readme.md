# /internal/utils

Direktori ini berisi berbagai fungsi utilitas (_helper/utility_) generik yang dapat digunakan ulang di berbagai bagian aplikasi.
Fungsi di sini **tidak spesifik modul bisnis**, tetapi menyediakan fungsionalitas umum yang dibutuhkan banyak fitur.

---

## 📋 Daftar Utility

| File          | Fungsi                                      | Deskripsi Singkat                                | Status              |
| ------------- | ------------------------------------------- | ------------------------------------------------ | ------------------- |
| `password.go` | `HashPassword`, `CheckPasswordHash`,        | Utility hash password aman dengan bcrypt         | ✅ Selesai & Teruji |
|               | `CheckPasswordHashWithError`, `GetHashCost` | -                                                |                     |
| `jwt.go`      | `GenerateToken`, `ValidateToken`            | Utility membuat & memverifikasi JWT (login/auth) | ✅ Selesai & Teruji |

---

## Penjelasan Utility

- **password.go**

  - Hash password plaintext ke bcrypt, dengan cost configurable.
  - Verifikasi password user saat login, aman dan tidak rawan rainbow table.
  - Ada fungsi cek hash error & audit cost parameter (untuk audit keamanan).

- **jwt.go**
  - Membuat JWT token menggunakan `HS256` untuk otentikasi user (login, API).
  - Verifikasi signature, cek expiry, dan extract user ID dari token.
  - Aman untuk production, secret harus panjang dan random!

---

## Contoh Penggunaan

```go
// Hash password user
hash, err := utils.HashPassword("plainpassword")

// Validasi login user
valid := utils.CheckPasswordHash("inputPassword", hash)

// Generate JWT token
token, err := utils.GenerateToken("user-id", time.Hour*24, jwtSecret)
```

---

> Jika menambah file utility baru di folder ini, update README dan jelaskan fungsinya untuk memudahkan tim lain!

**Terakhir update:** [isi tanggal & inisial updater]

---
