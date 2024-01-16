package category

import (
	"compare/components"
	"compare/internal"
	"compare/models"
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func AdminRoute(db *sql.DB, c *CreateRouteConfig) func(r chi.Router) {
	if c != nil {
		imageSaveDir = c.ImageSaveDir
	}

	// Get superuser token or generate one if there's none
	tokenModel, err := models.UserTokens(qm.Where(fmt.Sprintf("%s = true", models.UserTokenColumns.Superuser))).One(context.Background(), db)
	var token string
	if err == sql.ErrNoRows {
		for {
			token, err = internal.GenerateToken(10)
			if err != nil {
				panic(err)
			}
			exists, err := models.UserTokens(qm.Where(fmt.Sprintf("%s = ?", models.UserTokenColumns.Token), token)).Exists(context.Background(), db)
			if err != nil {
				panic(err)
			}
			if !exists {
				break
			}
		}
		tokenModel = &models.UserToken{
			Token:     token,
			Superuser: true,
		}
		tokenModel.Insert(context.Background(), db, boil.Infer())
	} else if err != nil {
		panic(err)
	} else {
		token = tokenModel.Token
	}

	fmt.Printf("---------------\nADMIN TOKEN: %s\n---------------\n", token)

	return func(r chi.Router) {
		r.Get("/", AdminGET)
		r.Post("/", AdminPOST(db))
		r.Get("/{token:[\\w-]+}", AdminDashboardGET(db))
		r.Get("/{token:[\\w-]+}/review", CardGET(db, true))
		r.Patch("/{token:[\\w-]+}/review/{id:\\d+}", CardPATCH(db, true))
		r.Delete("/{token:[\\w-]+}/review/{id:\\d+}", CardDELETE(db, true))
		r.Get("/{token:[\\w-]+}/edit", CardGET(db, false))
		r.Patch("/{token:[\\w-]+}/edit/{id:\\d+}", CardPATCH(db, false))
		r.Delete("/{token:[\\w-]+}/edit/{id:\\d+}", CardDELETE(db, false))
	}
}

func AdminGET(w http.ResponseWriter, r *http.Request) {
	templ.Handler(components.AdminTokenIndex(false)).ServeHTTP(w, r)
}

func AdminDashboardGET(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		token := chi.URLParam(r, "token")

		valid, err := validateToken(token, r.Context(), db)
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to check if token is valid. err = %s", err), http.StatusInternalServerError)
			return
		} else if !valid {
			templ.Handler(components.AdminForm(true)).ServeHTTP(w, r)
			return
		}

		reviewCount, err := getPendingCards(r.Context(), db)
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to count reviews. err = %s", err), http.StatusInternalServerError)
			return
		}
		templ.Handler(components.AdminDashboard(token, reviewCount)).ServeHTTP(w, r)
	}
}

func AdminPOST(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		token := r.PostFormValue("token")
		valid, err := validateToken(token, r.Context(), db)
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to check if token is valid. err = %s", err), http.StatusInternalServerError)
			return
		} else if !valid {
			templ.Handler(components.AdminForm(true)).ServeHTTP(w, r)
			return
		}

		w.Header().Add("HX-Location", fmt.Sprintf("/admin/%s", token))
		w.Write([]byte("200"))
	}
}

func CardGET(db *sql.DB, isReview bool) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		token := chi.URLParam(r, "token")

		valid, err := validateToken(token, r.Context(), db)
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to check if token is valid. err = %s", err), http.StatusInternalServerError)
			return
		} else if !valid {
			templ.Handler(components.AdminForm(true)).ServeHTTP(w, r)
			return
		}

		// TODO: make this more efficient/paginated if there are a lot of cards to review
		cards, err := models.Cards(qm.Where(fmt.Sprintf("%s = ?", models.CardColumns.Accepted), !isReview)).All(r.Context(), db)
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to fetch cards. err = %s", err), http.StatusInternalServerError)
			return
		}

		if isReview {
			templ.Handler(components.ReviewIndex(token, cards)).ServeHTTP(w, r)
		} else {
			templ.Handler(components.EditIndex(token, cards)).ServeHTTP(w, r)
		}

	}
}

func CardPATCH(db *sql.DB, isReview bool) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		token := chi.URLParam(r, "token")

		valid, err := validateToken(token, r.Context(), db)
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to check if token is valid. err = %s", err), http.StatusInternalServerError)
			return
		} else if !valid {
			templ.Handler(components.AdminForm(true)).ServeHTTP(w, r)
			return
		}

		r.ParseForm()
		name := r.FormValue("name")
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, "invalid url parameter", http.StatusBadRequest)
			return
		}

		tx, err := db.BeginTx(r.Context(), nil)
		if err != nil {
			http.Error(w, fmt.Sprintf("transaction failed. err = %s", err), http.StatusInternalServerError)
			return
		}

		card, err := models.Cards(qm.Where(fmt.Sprintf("%s = ? AND %s = ?", models.CardColumns.Accepted, models.CardColumns.ID), !isReview, id)).One(r.Context(), tx)
		if err != nil {
			tx.Rollback()
			http.Error(w, fmt.Sprintf("unable to fetch card. err = %s", err), http.StatusInternalServerError)
			return
		}
		if name != "" {
			card.Name = name
		}
		card.Accepted = true
		card.Update(r.Context(), tx, boil.Infer())

		err = tx.Commit()
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to commit. err = %s", err), http.StatusInternalServerError)
			return
		}

		var count int64
		if isReview {
			count, err = getPendingCards(r.Context(), db)
			if err != nil {
				if err != nil {
					http.Error(w, fmt.Sprintf("failed to count reviews. err = %s", err), http.StatusInternalServerError)
					return
				}
			}
		} else {
			count, err = getReviewedCards(r.Context(), db)
			if err != nil {
				if err != nil {
					http.Error(w, fmt.Sprintf("failed to count reviews. err = %s", err), http.StatusInternalServerError)
					return
				}
			}
		}

		if count == 0 {
			w.Header().Add("HX-Refresh", "true")
		}

		if isReview {
			w.Write([]byte("200"))
		} else {
			templ.Handler(components.EditCard(token, card, isReview)).ServeHTTP(w, r)
		}

	}
}

func CardDELETE(db *sql.DB, isReview bool) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		token := chi.URLParam(r, "token")

		valid, err := validateToken(token, r.Context(), db)
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to check if token is valid. err = %s", err), http.StatusInternalServerError)
			return
		} else if !valid {
			templ.Handler(components.AdminForm(true)).ServeHTTP(w, r)
			return
		}

		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, "invalid url parameter", http.StatusBadRequest)
			return
		}

		tx, err := db.BeginTx(r.Context(), nil)
		if err != nil {
			http.Error(w, fmt.Sprintf("transaction failed. err = %s", err), http.StatusInternalServerError)
			return
		}

		card, err := models.Cards(qm.Where(fmt.Sprintf("%s = ? AND %s = ?", models.CardColumns.Accepted, models.CardColumns.ID), !isReview, id)).One(r.Context(), tx)
		if err != nil {
			tx.Rollback()
			http.Error(w, fmt.Sprintf("unable to delete card. err = %s", err), http.StatusInternalServerError)
			return
		}
		card.Delete(r.Context(), tx)

		err = tx.Commit()
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to commit. err = %s", err), http.StatusInternalServerError)
			return
		}

		// delete file locally. Ignore if something fails
		path := path.Join(imageSaveDir, card.Filename)
		err = os.Remove(path)
		if err != nil {
			log.Printf("unable to delete file. err = %s", err)
		}

		var count int64
		if isReview {
			count, err = getPendingCards(r.Context(), db)
			if err != nil {
				if err != nil {
					http.Error(w, fmt.Sprintf("failed to count reviews. err = %s", err), http.StatusInternalServerError)
					return
				}
			}
		} else {
			count, err = getReviewedCards(r.Context(), db)
			if err != nil {
				if err != nil {
					http.Error(w, fmt.Sprintf("failed to count reviews. err = %s", err), http.StatusInternalServerError)
					return
				}
			}
		}

		if count == 0 {
			w.Header().Add("HX-Refresh", "true")
		}

		w.Write([]byte("200"))
	}
}

func EditGET(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

// Validates if a given token is valid. Returns true if the token is valid.
func validateToken(token string, c context.Context, db *sql.DB) (bool, error) {
	if token == "" {
		return false, nil
	}
	exists, err := models.UserTokens(qm.Where(fmt.Sprintf("%s = ?", models.UserTokenColumns.Token), token)).Exists(c, db)
	if err == sql.ErrNoRows {
		return false, nil
	} else if err != nil {
		return false, err
	} else if !exists {
		return false, nil
	}
	return true, nil
}

func getReviewedCards(c context.Context, db *sql.DB) (int64, error) {
	return models.Cards(qm.Where(fmt.Sprintf("%s = true", models.CardColumns.Accepted))).Count(c, db)
}

func getPendingCards(c context.Context, db *sql.DB) (int64, error) {
	return models.Cards(qm.Where(fmt.Sprintf("%s = false", models.CardColumns.Accepted))).Count(c, db)
}
