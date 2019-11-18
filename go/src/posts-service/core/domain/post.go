package domain

import "time"

type Post struct {
	PostID      string `json:"postId"`
	AuthorID    string `json:"authorId"`
	Title       string `json:"title"`
	Slug        string `json:"slug"`
	Body        string `json:"body"`
	Description string `json:"description"`
	Categories  []string `json:"categories"`
	CreatedAt   *time.Time `json:"createdAt"`
	PublishedAt *time.Time `json:"publishedAt,omitempty"`
}