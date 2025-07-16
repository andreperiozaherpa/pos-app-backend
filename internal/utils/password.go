package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword returns a bcrypt hash of the password using the provided cost.
// If cost is 0, it uses bcrypt.DefaultCost.
func HashPassword(password string, cost ...int) (string, error) {
	var c int
	if len(cost) > 0 {
		c = cost[0]
	} else {
		c = bcrypt.DefaultCost
	}
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), c)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(bytes), nil
}

// CheckPasswordHash compares a bcrypt hashed password with its possible plaintext equivalent.
// Returns true if they match.
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// CheckPasswordHashWithError does the same as CheckPasswordHash but returns the error for logging/debug.
func CheckPasswordHashWithError(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

// GetHashCost extracts the bcrypt cost parameter from a hashed password.
func GetHashCost(hash string) (int, error) {
	return bcrypt.Cost([]byte(hash))
}
