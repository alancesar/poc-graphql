package infra

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func NewPostgresConnection(host, database, username, password string, port int, ssl bool) (*gorm.DB, error) {
	connection := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s",
		host, port, username, database, password)

	if !ssl {
		connection = connection + " sslmode=disable"
	}

	return gorm.Open("postgres", connection)
}