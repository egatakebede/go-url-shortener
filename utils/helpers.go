package utils

import (
	"math/rand"
	"time"
)

// charset defines characters we use for short codes
const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// GenerateShortCode creates a random string of given length
func GenerateShortCode(length int) string {
	// Create a local random generator seeded with current time
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	code := make([]byte, length)
	for i := range code {
		code[i] = charset[r.Intn(len(charset))]
	}
	return string(code)
}
