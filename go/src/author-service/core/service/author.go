package service

import (
	"author-service/core/domain"
	"author-service/core/repository"
	"fmt"
	"github.com/google/uuid"
)

type Author interface {
	GetById(id string) (domain.Author, bool, error)
	Add(author domain.Author) (domain.Author, error)
	Update(id string, author domain.Author) (domain.Author, bool, error)
}

type service struct {
	repository repository.AuthorRepository
}

var (
	emptyAuthor = domain.Author{}
)

func NewAuthorService(repository repository.AuthorRepository) Author {
	return service{repository: repository}
}

func (service service) GetById(id string) (domain.Author, bool, error) {
	author, exists, err := service.repository.GetById(id)

	if !exists {
		return emptyAuthor, exists, fmt.Errorf(fmt.Sprintf("author %s not found", id))
	}

	return author, exists, err
}

func (service service) Add(author domain.Author) (domain.Author, error) {
	uuid, err := uuid.NewRandom()

	if err != nil {
		return emptyAuthor, err
	}

	author.ID = uuid.String()
	return author, service.repository.Add(author)
}

func (service service) Update(id string, author domain.Author) (domain.Author, bool, error) {
	_, exists, err := service.repository.GetById(id)

	if !exists {
		return emptyAuthor, exists, fmt.Errorf("not found")
	}

	if err != nil {
		return emptyAuthor, exists, err
	}

	author.ID = id
	return author, exists, service.repository.Update(author)
}