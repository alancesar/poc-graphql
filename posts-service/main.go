package main

import (
	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
	"net/http"
	"posts-service/core/api"
	"posts-service/core/repository"
	"posts-service/core/service"
)

func main() {
	r := mux.NewRouter()

	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "blog"
	session, _ := cluster.CreateSession()
	defer session.Close()

	repository := repository.NewPostRepository(session)
	service := service.NewPostService(repository)

	s := r.PathPrefix("/api/posts").Subrouter()
	s.HandleFunc("", api.ListHandler(service))
	s.HandleFunc("/search", api.PostByAuthorHandler(service))
	s.HandleFunc("/{id}", api.PostHandler(service))

	http.ListenAndServe(":8080", s)
}
