package data

import (
	"time"
)

type Book struct {
    ID        int64     `json:"id"`
    Title     string    `json:"-"`
    CreatedAt time.Time `json:"created_at, omitempty"`
    Published time.Time `json:"published"`
    Pages     int       `json:"pages"`
    Genres    []string  `json:"genres"`
    Rating    float64   `json:"rating"`
    Version   int       `json:"version"`
}