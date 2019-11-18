package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"posts-service/core/service"
)

func ListHandler(service service.Post) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("content-type", "application/json")
		posts := service.GetAll()

		if err := json.NewEncoder(w).Encode(posts); err != nil {
			fmt.Printf("error")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func PostHandler(service service.Post) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("content-type", "application/json")
		id := mux.Vars(r)["id"]
		post := service.GetById(id)

		if err := json.NewEncoder(w).Encode(post); err != nil {
			fmt.Printf("error")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func PostByAuthorHandler(service service.Post) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("content-type", "application/json")
		authorID := r.URL.Query().Get("author")
		posts := service.GetByAuthorId(authorID)

		if err := json.NewEncoder(w).Encode(posts); err != nil {
			fmt.Printf("error")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}