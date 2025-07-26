package users

import (
	"crypto/rand"
	"encoding/base64"
	"log"
)

// Create JWT key
func NewJwtKey(size int) string {
	key := make([]byte, size)
	if _, err := rand.Read(key); err != nil {
		log.Fatalf("failed to generate jwt key: %v", err)
	}
	return base64.StdEncoding.EncodeToString(key)
}
