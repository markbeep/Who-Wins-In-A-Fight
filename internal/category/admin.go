package category

import (
	"compare/components"
	"compare/models"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func AdminGET(w http.ResponseWriter, r *http.Request) {
	templ.Handler(components.AdminTokenIndex()).ServeHTTP(w, r)
}

func AdminTokenGET(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		token := chi.URLParam(r, "token")

		exists, err := models.UserTokens(qm.Where(fmt.Sprintf("%s = ?", models.UserTokenColumns.Token), token)).Exists(r.Context(), db)
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to check if token is valid. err = %s", err), http.StatusInternalServerError)
			return
		} else if !exists {
			http.Error(w, "Invalid token", http.StatusBadRequest)
			return
		}

		// TODO: show actual page
	}
}

func AdminPOST(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		token := r.PostFormValue("token")
		if token == "" {
			http.Error(w, "Invalid token", http.StatusBadRequest)
			return
		}

		exists, err := models.UserTokens(qm.Where(fmt.Sprintf("%s = ?", models.UserTokenColumns.Token), token)).Exists(r.Context(), db)
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to check if token is valid. err = %s", err), http.StatusInternalServerError)
			return
		} else if !exists {
			http.Error(w, "Invalid token", http.StatusBadRequest)
			return
		}

		w.Header().Add("HX-Push-Url", fmt.Sprintf("/admin/%s", token))

		// TODO: show actual page
	}
}
