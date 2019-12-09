package repository

import (
	"author-service/core/domain"
	"github.com/gocql/gocql"
	"github.com/jinzhu/gorm"
)

type AuthorRepository interface {
	GetById(id string) (domain.Author, bool, error)
	Add(author domain.Author) error
	Update(author domain.Author) error
}

func NewAuthorPostgresRepository(database *gorm.DB) AuthorRepository {
	database.AutoMigrate(&author{})
	return postgres{db: database}
}

func NewAuthorCassandraRepository(session *gocql.Session) AuthorRepository {
	return cassandra{session: session}
}
