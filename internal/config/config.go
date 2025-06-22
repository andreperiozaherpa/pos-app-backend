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
	// Coba muat file .env terlebih dahulu. Ini akan memuat variabel ke environment proses,
	// tapi tidak akan menimpa variabel yang sudah ada di environment sistem.
	err := godotenv.Load()
	if err != nil {
		log.Println("Peringatan: Tidak dapat memuat file .env:", err)
		// Tidak mengembalikan error jika .env tidak ada, karena mungkin disengaja,
		// terutama di lingkungan non-development.
	} else {
		log.Println(".env file loaded (if it existed).")
	}

	// Ambil nilai ENVIRONMENT setelah mencoba memuat .env
	// Jika tidak ada di env sistem atau .env, default ke "development"
	environment := os.Getenv("ENVIRONMENT")
	if environment == "" {
		environment = "development" // Default ke development jika tidak ada sama sekali
	}
	log.Println("ENVIRONMENT yang digunakan:", environment)

	// Coba muat file .env jika ada (berguna untuk development)
	// Di production, environment variables biasanya diatur langsung di server/container.

	appPortStr := os.Getenv("APP_PORT")
	if appPortStr == "" {
		appPortStr = "8080" // Default port
	}
	appPort, err := strconv.Atoi(appPortStr)
	if err != nil {
		return nil, fmt.Errorf("APP_PORT tidak valid: %w", err)
	}

	// Debugging: Cetak nilai variabel setelah mencoba memuat .env
	log.Println("APP_NAME:", os.Getenv("APP_NAME"))
	log.Println("APP_PORT:", os.Getenv("APP_PORT"))
	log.Println("DATABASE_URL:", os.Getenv("DATABASE_URL"))

	cfg := &AppConfig{
		AppPort:     appPort,
		AppName:     os.Getenv("APP_NAME"),
		DatabaseURL: os.Getenv("DATABASE_URL"),
		JWTSecret:   os.Getenv("JWT_SECRET"),
		Environment: environment, // Gunakan variabel 'environment' yang sudah diproses
	}

	// Anda bisa menambahkan validasi di sini untuk memastikan variabel penting tidak kosong
	// Misalnya: if cfg.DatabaseURL == "" { return nil, errors.New("DATABASE_URL harus diatur") }

	return cfg, nil
}
