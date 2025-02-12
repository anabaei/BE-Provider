package data

import (
    "database/sql"
    "github.com/lib/pq"
    "time"
)

type Post struct {
    ID        int64     `json:"id"`
    Title     string    `json:"title"`
    Content   string    `json:"content"`
    Author    string    `json:"author"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
    Version   int       `json:"version"`
}

func (p *Post) GetAllPosts(db *sql.DB) ([]*Post, error) {
    query := `
        SELECT id, title, content, author, created_at, updated_at, version
        FROM posts
        ORDER BY created_at DESC`

    rows, err := db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var posts []*Post

    for rows.Next() {
        var post Post
        err := rows.Scan(
            &post.ID,
            &post.Title,
            &post.Content,
            &post.Author,
            &post.CreatedAt,
            &post.UpdatedAt,
            &post.Version,
        )
        if err != nil {
            return nil, err
        }
        posts = append(posts, &post)
    }

    if err = rows.Err(); err != nil {
        return nil, err
    }

    return posts, nil
}