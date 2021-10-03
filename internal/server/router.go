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
	methods := []string{http.MethodPost, http.MethodGet, http.MethodOptions}
	origins := []string{"http://localhost:4200", "https://letters2lostlovedones.com"}
	headers := []string{"Content-Type"}

	opts := cors.Options{
		AllowedMethods:     methods,
		AllowedOrigins:     origins,
		AllowedHeaders:     headers,
		OptionsPassthrough: true,
		Debug:              true,
	}
	crs := cors.New(opts)

	return crs.Handler(router(c))
}

func addRoutes(r *mux.Router, c *letter.Controller) {
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	r.HandleFunc("/users/{user}/letter", letter.GetAllForUser(c)).Methods(http.MethodGet, http.MethodOptions)
	//r.HandleFunc("/letter/{id}/meta", letter.GetLetterById(c)).Methods(http.MethodGet, http.MethodOptions)
	//r.HandleFunc("/letter", letter.InsertLetter(c)).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/letters/{id}/content", letter.GetLetterContentById(c)).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/letter/{id}/meta", letter.GetMetaForLetter(c)).Methods(http.MethodGet, http.MethodOptions)

	r.HandleFunc("/users/{user}/metadata", letter.GetMetaForUser(c)).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/letter", letter.InsertLetterHTML(c)).Methods(http.MethodPost, http.MethodOptions)

	r.HandleFunc("/token/profile", letter.GetLoginRadiusUserDetails(c)).Methods(http.MethodGet, http.MethodOptions)
}
