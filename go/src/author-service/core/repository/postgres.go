package repository

import (
	"author-service/core/domain"
	"github.com/jinzhu/gorm"
)

type postgres struct {
	db *gorm.DB
}

type author struct {
	gorm.Model
	UUID     string
	Name     string
	Username string
	Email    string
}

var (
	emptyResponse = domain.Author{}
)

func (postgres postgres) GetById(uuid string) (domain.Author, bool, error) {
	output := author{}
	if query := postgres.db.Where("uuid = ?", uuid).First(&output); query.Error != nil {
		return emptyResponse, false, query.Error
	}

	if output.UUID == "" {
		return emptyResponse, false, nil
	}

	return postgres.parseToDomain(output), true, nil
}

func (postgres postgres) Add(author domain.Author) error {
	entity := postgres.parseFromDomain(author)
	if query := postgres.db.Create(&entity); query.Error != nil {
		return query.Error
	}

	return nil
}

func (postgres postgres) Update(input domain.Author) error {
	query := postgres.db.Where("uuid = ?", input.ID).Updates(author{
		Name:     input.Name,
		Username: input.Username,
		Email:    input.Username,
	})
	return query.Error
}

func (postgres postgres) parseFromDomain(domain domain.Author) author {
	return author{
		UUID:     domain.ID,
		Name:     domain.Name,
		Username: domain.Username,
		Email:    domain.Email,
	}
}

func (postgres postgres) parseToDomain(entity author) domain.Author {
	return domain.Author{
		ID:       entity.UUID,
		Name:     entity.Name,
		Username: entity.Username,
		Email:    entity.Email,
	}
}
