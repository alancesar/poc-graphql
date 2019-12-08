package repository

import (
	"fmt"
	"github.com/gocql/gocql"
	"posts-service/core/domain"
	"time"
)

const (
	allFieldsPosts = "post_id,author_id,title,slug,body,description,categories,created_at,published_at"
)

type Post interface {
	GetAll() []domain.Post
	GetById(id string) domain.Post
	GetByAuthorId(authorID string) []domain.Post
}

type cassandra struct {
	session *gocql.Session
}

func NewPostRepository(session *gocql.Session) Post {
	return cassandra{session: session}
}

func (repository cassandra) GetAll() []domain.Post {
	query := fmt.Sprintf("SELECT %s FROM blog.posts", allFieldsPosts)
	iter := repository.session.Query(query).Iter()

	var posts []domain.Post
	m := map[string]interface{}{}

	for iter.MapScan(m) {
		posts = append(posts, repository.unmarshal(m))
		m = map[string]interface{}{}
	}

	return posts
}

func (repository cassandra) GetById(id string) domain.Post {
	m := map[string]interface{}{}
	query := fmt.Sprintf("SELECT %s FROM blog.posts WHERE post_id = ?", allFieldsPosts)
	err := repository.session.Query(query, id).MapScan(m)

	if err != nil {
		fmt.Println(fmt.Sprintf("error getting post %s", id))
		return domain.Post{}
	}

	return repository.unmarshal(m)
}

func (repository cassandra) GetByAuthorId(authorID string) []domain.Post {
	query := fmt.Sprintf("SELECT %s FROM blog.posts WHERE author_id = ? ALLOW FILTERING", allFieldsPosts)
	iter := repository.session.Query(query, authorID).Iter()

	var posts []domain.Post
	m := map[string]interface{}{}

	for iter.MapScan(m) {
		posts = append(posts, repository.unmarshal(m))
		m = map[string]interface{}{}
	}

	return posts
}

func (cassandra) unmarshal(post map[string]interface{}) domain.Post {
	safeStringUnmarshal := func(input interface{}) string {
		if input != nil {
			return input.(string)
		}

		return ""
	}

	timeUnmarshal := func(input interface{}) *time.Time {
		if input == nil || input.(time.Time).IsZero() {
			return nil
		}

		time := input.(time.Time)
		return &time
	}

	return domain.Post{
		PostID:      post["post_id"].(gocql.UUID).String(),
		AuthorID:    post["author_id"].(gocql.UUID).String(),
		Title:       safeStringUnmarshal(post["title"]),
		Slug:        safeStringUnmarshal(post["slug"]),
		Body:        safeStringUnmarshal(post["body"]),
		Description: safeStringUnmarshal(post["description"]),
		Categories:  post["categories"].([]string),
		CreatedAt:   timeUnmarshal(post["created_at"]),
		PublishedAt: timeUnmarshal(post["published_at"]),
	}
}
