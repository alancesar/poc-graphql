package main

import (
	"author-service/core/api"
	"author-service/core/repository"
	"author-service/core/service"
	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "blog"
	session, _ := cluster.CreateSession()
	defer session.Close()

	repository := repository.NewAuthorRepository(session)
	service := service.NewAuthorService(repository)

	s := r.PathPrefix("/api/authors").Subrouter()
	s.HandleFunc("", api.ListAuthorHandler(service)).Methods(http.MethodGet)
	s.HandleFunc("", api.AddAuthorHandler(service)).Methods(http.MethodPost)
	s.HandleFunc("/{id}", api.AuthorHandler(service)).Methods(http.MethodGet)
	s.HandleFunc("/{id}", api.UpdateAuthorHandler(service)).Methods(http.MethodPut)

	http.ListenAndServe(":8081", s)
}

