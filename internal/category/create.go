package category

import (
	"compare/components"
	"net/http"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
)

func CreateRoute() http.Handler {
	r := chi.NewRouter()
	r.Get("/", createGET)
	r.Post("/", createPOST)
	return r
}

func createGET(w http.ResponseWriter, r *http.Request) {
	templ.Handler(components.CreateIndex()).ServeHTTP(w, r)
}

func createPOST(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("HX-Push-Url", "/")
	token, err := GenerateToken(10)
	if err != nil {
		http.Error(w, "failed to generate token", http.StatusInternalServerError)
		return
	}
	_ = token
	// generate new category
	templ.Handler(components.CreateIndex())
}
