package models

import "time"

// URL represents a shortened URL
type URL struct {
	OriginalURL string
	ShortCode   string
	CreatedAt   time.Time
}
