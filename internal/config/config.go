package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// AppConfig menampung semua konfigurasi aplikasi.
type AppConfig struct {
	AppPort     int
	AppName     string
	DatabaseURL string
	JWTSecret   string
	Environment string // "development", "staging", "production"
}

// LoadConfig memuat konfigurasi aplikasi dari environment variables.
// Jika ENVIRONMENT adalah "development", ia juga akan mencoba memuat dari file .env.
func LoadConfig() (*AppConfig, error) {
	// Coba muat file .env jika ada (berguna untuk development)
	// Di production, environment variables biasanya diatur langsung di server/container.
	environment := os.Getenv("ENVIRONMENT")
	if environment == "" || environment == "development" {
		err := godotenv.Load() // Memuat dari .env di root proyek
		if err != nil {
			log.Println("Peringatan: Tidak dapat memuat file .env:", err)
			// Tidak mengembalikan error jika .env tidak ada, karena mungkin disengaja.
		}
	}

	appPortStr := os.Getenv("APP_PORT")
	if appPortStr == "" {
		appPortStr = "8080" // Default port
	}
	appPort, err := strconv.Atoi(appPortStr)
	if err != nil {
		return nil, fmt.Errorf("APP_PORT tidak valid: %w", err)
	}

	cfg := &AppConfig{
		AppPort:     appPort,
		AppName:     os.Getenv("APP_NAME"),
		DatabaseURL: os.Getenv("DATABASE_URL"),
		JWTSecret:   os.Getenv("JWT_SECRET"),
		Environment: environment,
	}

	// Anda bisa menambahkan validasi di sini untuk memastikan variabel penting tidak kosong
	// Misalnya: if cfg.DatabaseURL == "" { return nil, errors.New("DATABASE_URL harus diatur") }

	return cfg, nil
}
