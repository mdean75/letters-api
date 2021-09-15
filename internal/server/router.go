package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"letters-api/internal/letter"
)

func router(c *letter.Controller) *mux.Router {
	r := mux.NewRouter()
	addRoutes(r, c)

	return r
}

func addRoutes(r *mux.Router, c *letter.Controller) {
	r.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})

	r.HandleFunc("/users/{user}/letter", letter.GetAllForUser(c)).Methods(http.MethodGet)
	r.HandleFunc("/letter/{id}", letter.GetLetterById(c)).Methods(http.MethodGet)
	r.HandleFunc("/letter", letter.InsertLetter(c)).Methods(http.MethodPost)
}
