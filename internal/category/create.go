package category

import (
	"compare/components"
	"compare/models"
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type CreateRouteConfig struct {
	MaxMemory    int64  // max memory amount uploaded files will take up in memory
	ImageSaveDir string // directory to save images in
}

var (
	config = CreateRouteConfig{}
)

func CreateRoute(db *sql.DB, c *CreateRouteConfig) http.Handler {
	if c != nil {
		config = *c
		err := os.MkdirAll(c.ImageSaveDir, 0755)
		if err != nil {
			panic("failed to create image directory")
		}
	}
	r := chi.NewRouter()
	r.Get("/", createGET)
	r.Post("/", createPOST(db))
	return r
}

func createGET(w http.ResponseWriter, r *http.Request) {
	templ.Handler(components.CreateIndex()).ServeHTTP(w, r)
}

func createPOST(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseMultipartForm(config.MaxMemory)
		if err != nil {
			http.Error(w, "Failed to parse form", http.StatusBadRequest)
			return
		}

		title := r.PostFormValue("title")
		if len(title) == 0 {
			http.Error(w, "Missing title value", http.StatusBadRequest)
			return
		}
		description := r.PostFormValue("description")
		if len(description) == 0 {
			http.Error(w, "Missing description value", http.StatusBadRequest)
			return
		}
		files, ok := r.MultipartForm.File["images"]
		if !ok || len(files) == 0 {
			http.Error(w, "No files given", http.StatusBadRequest)
			return
		}

		tx, err := db.BeginTx(r.Context(), nil)
		if err != nil {
			http.Error(w, "failed to start transaction", http.StatusInternalServerError)
			return
		}

		var categoryToken, adminToken string

		for {
			categoryToken, err = GenerateToken(10)
			if err != nil {
				tx.Rollback()
				http.Error(w, fmt.Sprintf("failed to generate category token. err = %s", err), http.StatusInternalServerError)
				return
			}
			exists, err := models.Categories(qm.Where(fmt.Sprintf("%s = ?", models.CategoryColumns.Token), categoryToken)).Exists(r.Context(), db)
			if err != nil {
				tx.Rollback()
				http.Error(w, fmt.Sprintf("failed to check if category token already exists. err = %s", err), http.StatusInternalServerError)
				return
			}
			if !exists {
				break
			}
		}

		for {
			adminToken, err = GenerateToken(10)
			if err != nil {
				tx.Rollback()
				http.Error(w, fmt.Sprintf("failed to generate admin token. err = %s", err), http.StatusInternalServerError)
				return
			}
			exists, err := models.UserTokens(qm.Where(fmt.Sprintf("%s = ?", models.UserTokenColumns.Token), adminToken)).Exists(r.Context(), db)
			if err != nil {
				tx.Rollback()
				http.Error(w, fmt.Sprintf("failed to check if admin token already exists. err = %s", err), http.StatusInternalServerError)
				return
			}
			if !exists {
				break
			}
		}

		// generate new category and admin token for it
		category := models.Category{
			Token:       categoryToken,
			Title:       title,
			Description: description,
		}
		userToken := models.UserToken{
			Token: adminToken,
		}
		category.Insert(r.Context(), db, boil.Infer())
		category.AddUserTokens(r.Context(), db, true, &userToken)

		for _, f := range files {
			fo, err := f.Open()
			if err != nil {
				tx.Rollback()
				http.Error(w, fmt.Sprintf("failed to open file. err = %s", err), http.StatusInternalServerError)
				return
			}
			defer fo.Close()

			// Generate unique token
			var imageToken string
			var imageName string
			var imagePath string
			for {
				imageToken, err = GenerateToken(10)
				if err != nil {
					tx.Rollback()
					http.Error(w, fmt.Sprintf("failed to generate admin token. err = %s", err), http.StatusInternalServerError)
					return
				}
				imageName = imageToken + path.Ext(f.Filename)
				imagePath = path.Join(config.ImageSaveDir, imageName)
				if _, err := os.Stat(imagePath); err != nil {
					// File doesn't exist
					break
				}
			}

			// Here we save the file to the defined location
			save, err := os.Create(imagePath)
			if err != nil {
				tx.Rollback()
				http.Error(w, fmt.Sprintf("failed to create file. err = %s", err), http.StatusInternalServerError)
				return
			}
			defer save.Close()
			b, err := io.ReadAll(fo)
			if err != nil {
				tx.Rollback()
				http.Error(w, fmt.Sprintf("failed to read image file. err = %s", err), http.StatusInternalServerError)
				return
			}
			// TODO: compress image
			writeLen, err := save.Write(b)
			if err != nil {
				tx.Rollback()
				http.Error(w, fmt.Sprintf("failed to write new image file. err = %s", err), http.StatusInternalServerError)
				return
			}

			if writeLen == 0 {
				tx.Rollback()
				http.Error(w, "Written file length is 0", http.StatusInternalServerError)
				return
			}

			card := models.Card{Filename: imageName}
			category.AddCards(r.Context(), db, true, &card)
		}

		tx.Commit()

		w.Header().Add("HX-Push-Url", fmt.Sprintf("/category/%s/edit/%s", categoryToken, adminToken))

		// TODO: return EDIT page
		templ.Handler(components.CreateIndex()).ServeHTTP(w, r)
	}
}
