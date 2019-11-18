package api

import (
	"author-service/core/domain"
	"author-service/core/service"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func AuthorHandler(service service.Author) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("content-type", "application/json")
		id := mux.Vars(r)["id"]
		author, exists, err := service.GetById(id)

		if err != nil {
			if !exists {
				fmt.Printf(fmt.Sprintf("author %s not found", id))
				w.WriteHeader(http.StatusNotFound)
				return
			}

			fmt.Printf(fmt.Sprintf("error on updating author: %v", err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if err := json.NewEncoder(w).Encode(author); err != nil {
			fmt.Printf(fmt.Sprintf("error encode author: %v", err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if err := json.NewEncoder(w).Encode(author); err != nil {
			fmt.Printf(fmt.Sprintf("error encode author: %v", err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func ListAuthorHandler(service service.Author) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("content-type", "application/json")

		authors := service.GetAll()

		if err := json.NewEncoder(w).Encode(authors); err != nil {
			fmt.Printf(fmt.Sprintf("error encode author: %v", err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func AddAuthorHandler(service service.Author) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("content-type", "application/json")
		author := domain.Author{}
		json.NewDecoder(r.Body).Decode(&author)
		author, err := service.Add(author)

		if err != nil {
			fmt.Printf(fmt.Sprintf("error on adding author: %v", err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if err := json.NewEncoder(w).Encode(author); err != nil {
			fmt.Printf(fmt.Sprintf("error encode author: %v", err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func UpdateAuthorHandler(service service.Author) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("content-type", "application/json")
		id := mux.Vars(r)["id"]

		author := domain.Author{}
		json.NewDecoder(r.Body).Decode(&author)
		author, exists, err := service.Update(id, author)

		if err != nil {
			if !exists {
				fmt.Printf(fmt.Sprintf("author %s not found", id))
				w.WriteHeader(http.StatusNotFound)
				return
			}

			fmt.Printf(fmt.Sprintf("error on updating author: %v", err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if err := json.NewEncoder(w).Encode(author); err != nil {
			fmt.Printf(fmt.Sprintf("error encode author: %v", err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}