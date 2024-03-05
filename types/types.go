package types

import "github.com/google/uuid"

type ShortUrlModel struct {
	ID       uuid.UUID `json:"id"`
	URL      string    `json:"url"`
	ShortUrl uuid.UUID `json:"short_url"`
	Clicks   int       `json:"clicks"`
}

type CreateShortUrl struct {
	URL string `json:"url"`
}
