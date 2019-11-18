package service

import (
	"posts-service/core/domain"
	"posts-service/core/repository"
)

type Post interface {
	GetAll() []domain.Post
	GetById(id string) domain.Post
	GetByAuthorId(authorId string) []domain.Post
}

type service struct {
	repository repository.Post
}

func NewPostService(repository repository.Post) Post {
	return service{repository: repository}
}

func (service service) GetAll() []domain.Post {
	return service.repository.GetAll()
}

func (service service) GetById(id string) domain.Post {
	return service.repository.GetById(id)
}

func (service service) GetByAuthorId(authorID string) []domain.Post {
	return service.repository.GetByAuthorId(authorID)
}
