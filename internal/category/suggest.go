package category

import (
	"compare/components"
	"compare/internal"
	"compare/models"
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"slices"
	"strings"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type CreateRouteConfig struct {
	MaxMemory         int64  // max memory amount uploaded files will take up in memory
	ImageSaveDir      string // directory to save images in
	ImageOnlineDir    string // path to access files online
	AllowedExtensions []string
}

var maxMemory int64 = 10 << 20
var imageSaveDir = "./data/imgs"
var fileExtensions []string = []string{}

func SuggestRoute(db *sql.DB, c *CreateRouteConfig) func(r chi.Router) {
	if c != nil {
		maxMemory = c.MaxMemory
		imageSaveDir = c.ImageSaveDir
		fileExtensions = c.AllowedExtensions
	}
	err := os.MkdirAll(imageSaveDir, 0755)
	if err != nil {
		log.Fatalf("failed to create image directory. err = %s", err)
	}

	return func(r chi.Router) {
		r.Get("/", SuggestGET)
		r.Post("/", SuggestPOST(db))
	}
}

func SuggestGET(w http.ResponseWriter, r *http.Request) {
	templ.Handler(components.SuggestIndex(false, false, strings.Join(fileExtensions, ", "))).ServeHTTP(w, r)
}

func SuggestPOST(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseMultipartForm(maxMemory)
		if err != nil {
			templ.Handler(components.SuggestForm(true, true, strings.Join(fileExtensions, ", "))).ServeHTTP(w, r)
			return
		}

		name := r.PostFormValue("name")
		if len(name) == 0 {
			templ.Handler(components.SuggestForm(true, false, strings.Join(fileExtensions, ", "))).ServeHTTP(w, r)
			return
		}

		f, fh, err := r.FormFile("image")
		if err == http.ErrMissingFile {
			templ.Handler(components.SuggestForm(false, true, strings.Join(fileExtensions, ", "))).ServeHTTP(w, r)
			return
		} else if err != nil {
			http.Error(w, fmt.Sprintf("err = %s", err), http.StatusInternalServerError)
			return
		}
		defer f.Close()

		if fh.Size > maxMemory {
			templ.Handler(components.SuggestForm(false, true, strings.Join(fileExtensions, ", "))).ServeHTTP(w, r)
			return
		}

		if !slices.Contains(fileExtensions, path.Ext(fh.Filename)) {
			templ.Handler(components.SuggestForm(false, true, strings.Join(fileExtensions, ", "))).ServeHTTP(w, r)
			return
		}
		b, err := io.ReadAll(f)
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to read image file. err = %s", err), http.StatusInternalServerError)
			return
		}
		contentType := http.DetectContentType(b)
		if !strings.HasPrefix(contentType, "image/") {
			log.Printf("invalid filetype uploaded. got = %s", contentType)
			templ.Handler(components.SuggestForm(false, true, strings.Join(fileExtensions, ", "))).ServeHTTP(w, r)
			return
		}

		tx, err := db.BeginTx(r.Context(), nil)
		if err != nil {
			http.Error(w, "failed to start transaction", http.StatusInternalServerError)
			return
		}

		var imageToken string
		for {
			imageToken, err = internal.GenerateToken(10)
			if err != nil {
				tx.Rollback()
				http.Error(w, fmt.Sprintf("failed to generate image token. err = %s", err), http.StatusInternalServerError)
				return
			}
			exists, err := models.Cards(qm.Where(fmt.Sprintf("%s = ?", models.CardColumns.Token), imageToken)).Exists(r.Context(), tx)
			if err != nil {
				tx.Rollback()
				http.Error(w, fmt.Sprintf("failed to fetch tokens. err = %s", err), http.StatusInternalServerError)
				return
			}
			if !exists {
				break
			}
		}

		imageName := imageToken + path.Ext(fh.Filename)
		imageLocalPath := path.Join(imageSaveDir, imageName)

		save, err := os.Create(imageLocalPath)
		if err != nil {
			tx.Rollback()
			http.Error(w, fmt.Sprintf("failed to create file. err = %s", err), http.StatusInternalServerError)
			return
		}
		defer save.Close()
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

		card := models.Card{
			Name:     name,
			Token:    imageToken,
			Filename: imageName,
		}
		card.Insert(r.Context(), tx, boil.Infer())

		err = tx.Commit()
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to commit. err = %s", err), http.StatusInternalServerError)
			return
		}

		templ.Handler(components.SuggestSuccess(name)).ServeHTTP(w, r)
	}
}
