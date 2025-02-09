package data

import (
    "database/sql"
	"time"
    "github.com/lib/pq"
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

type BookModel struct{
    DB *sql.DB
}


func (b BookModel) Insert(book *book) error {
    query := `INSERT INTO books (title, published, pages
    , genres, rating) VALUES($1, $2, $3, $4, $5) 
    RETURNING id, created_at, version`

    args := []interface{}(book.Title, book.Published, book.Pages, pq.Array(book.Genres), book.rating)
    return b.DB.queryRow(query, args...).Scan(&book.ID, &book.CreatedAt, &book.Version)
} 