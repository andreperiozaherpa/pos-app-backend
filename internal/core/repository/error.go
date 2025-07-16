package repository

import "errors"

// ErrNotFound adalah error standar saat data tidak ditemukan
var ErrNotFound = errors.New("record not found")

// ErrUnauthorized adalah error standar saat autentikasi gagal
var ErrUnauthorized = errors.New("unauthorized")
