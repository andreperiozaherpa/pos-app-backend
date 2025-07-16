package postgres

import (
	"context"
	"log"
)

// LogQuery mencatat query dan durasinya.
func LogQuery(ctx context.Context, query string, durationMs int64) {
	log.Printf("[DB QUERY] %s (durasi: %d ms)\n", query, durationMs)
}

// LogError mencatat error pada repository atau DB.
func LogError(ctx context.Context, operation string, err error) {
	if err != nil {
		log.Printf("[DB ERROR] %s: %v\n", operation, err)
	}
}
