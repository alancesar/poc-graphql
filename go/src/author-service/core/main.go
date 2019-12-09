package main

import (
	"author-service/core/api"
	"author-service/core/infra"
	"author-service/core/repository"
	"author-service/core/service"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	connection, err := infra.NewPostgresConnection(
		"localhost", "blog", "blog", "blog", 5432, false)

	if err != nil {
		fmt.Errorf("error on connecting database: %v", err)
		panic(err)
	}

	repository := repository.NewAuthorPostgresRepository(connection)
	service := service.NewAuthorService(repository)

	s := r.PathPrefix("/api/authors").Subrouter()
	s.HandleFunc("", api.AddAuthorHandler(service)).Methods(http.MethodPost)
	s.HandleFunc("/{id}", api.AuthorHandler(service)).Methods(http.MethodGet)
	s.HandleFunc("/{id}", api.UpdateAuthorHandler(service)).Methods(http.MethodPut)

	http.ListenAndServe(":8081", s)
}

