package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq" // Driver PostgreSQL, underscore berarti diimpor untuk side effect (registrasi driver)
)

// DB adalah instance koneksi database global (atau bisa di-pass melalui dependency injection)
var DB *sql.DB

// ConnectDB membuat koneksi ke database PostgreSQL menggunakan DSN (Data Source Name)
// dan mengatur beberapa parameter koneksi.
func ConnectDB(dataSourceName string) (*sql.DB, error) {
	var err error
	DB, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("gagal membuka koneksi database: %w", err)
	}

	// Atur parameter koneksi (opsional tapi direkomendasikan)
	DB.SetMaxOpenConns(25)                 // Jumlah maksimum koneksi terbuka
	DB.SetMaxIdleConns(25)                 // Jumlah maksimum koneksi idle
	DB.SetConnMaxLifetime(5 * time.Minute) // Waktu maksimum koneksi bisa digunakan kembali

	if err = DB.Ping(); err != nil {
		return nil, fmt.Errorf("gagal melakukan ping ke database: %w", err)
	}

	log.Println("Berhasil terhubung ke database PostgreSQL!")
	return DB, nil
}
