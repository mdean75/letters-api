package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"letters-api/internal/letter"
)

func router(c *letter.Controller) *mux.Router {
	r := mux.NewRouter()
	addRoutes(r, c)

	return r
}

func routerWithCors(c *letter.Controller) http.Handler {
	methods := []string{"*"}
	origins := []string{"letters2lostlovedones.com", "localhost:4200"}
	crs := cors.New(cors.Options{AllowedMethods: methods, AllowedOrigins: origins})

	return crs.Handler(router(c))
}

func addRoutes(r *mux.Router, c *letter.Controller) {
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	r.HandleFunc("/users/{user}/letter", letter.GetAllForUser(c)).Methods(http.MethodGet)
	r.HandleFunc("/letter/{id}", letter.GetLetterById(c)).Methods(http.MethodGet)
	r.HandleFunc("/letter", letter.InsertLetter(c)).Methods(http.MethodPost)
}
