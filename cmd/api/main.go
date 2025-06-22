package main

import (
	"fmt"
	"log"
	"os"

	"pos-app/backend/internal/config"
	"pos-app/backend/internal/database" // Import package database
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Gagal memuat konfigurasi: %v", err)
		os.Exit(1) // Keluar jika konfigurasi gagal dimuat
	}

	fmt.Printf("Backend %s sedang berjalan di port %d (Lingkungan: %s)\n", cfg.AppName, cfg.AppPort, cfg.Environment)
	fmt.Printf("Database URL: %s\n", cfg.DatabaseURL)

	// Membuat koneksi ke database
	db, err := database.ConnectDB(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Gagal terhubung ke database: %v", err)
		os.Exit(1)
	}
	defer db.Close() // Pastikan koneksi database ditutup saat fungsi main selesai
	// Di sini kita akan mulai menyiapkan server HTTP, koneksi database, dll.
}
