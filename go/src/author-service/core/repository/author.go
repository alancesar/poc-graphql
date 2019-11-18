package repository

import (
	"author-service/core/domain"
	"github.com/gocql/gocql"
)

type Author interface {
	GetById(id string) (domain.Author, bool, error)
	GetAll() []domain.Author
	Add(author domain.Author) error
	Update(author domain.Author) error
}

type cassandra struct {
	session *gocql.Session
}

func NewAuthorRepository(session *gocql.Session) Author {
	return cassandra{session: session}
}

func (repository cassandra) GetById(id string) (domain.Author, bool, error) {
	m := map[string]interface{}{}
	query := "SELECT id, name, email, username FROM blog.authors WHERE id = ?"
	err := repository.session.Query(query, id).MapScan(m)

	if err != nil {
		return domain.Author{}, false, err
	}

	return repository.unmarshal(m), true, err
}

func (repository cassandra) GetAll() []domain.Author {
	query := "SELECT id, name, email, username FROM blog.authors"
	iter := repository.session.Query(query).Iter()

	var posts []domain.Author
	m := map[string]interface{}{}

	for iter.MapScan(m) {
		posts = append(posts, repository.unmarshal(m))
		m = map[string]interface{}{}
	}

	return posts
}

func (repository cassandra) Add(author domain.Author) error {
	query := "INSERT INTO blog.authors (id, name, email, username) VALUES (?, ?, ?, ?)"
	return repository.session.Query(query, author.ID, author.Name, author.Email, author.Username).Exec()
}

func (repository cassandra) Update(author domain.Author) error {
	query := "UPDATE blog.authors SET name = ?, email = ?, username = ? WHERE id = ?"
	return repository.session.Query(query, author.Name, author.Email, author.Username, author.ID).Exec()
}

func (cassandra) unmarshal(author map[string]interface{}) domain.Author {
	safeStringUnmarshal := func(input interface{}) string {
		if input != nil {
			return input.(string)
		}

		return ""
	}

	return domain.Author{
		ID:       author["id"].(gocql.UUID).String(),
		Name:     safeStringUnmarshal(author["name"]),
		Email:    safeStringUnmarshal(author["email"]),
		Username: safeStringUnmarshal(author["username"]),
	}
}
