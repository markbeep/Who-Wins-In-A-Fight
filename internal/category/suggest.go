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

	"github.com/a-h/templ"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type CreateRouteConfig struct {
	MaxMemory      int64  // max memory amount uploaded files will take up in memory
	ImageSaveDir   string // directory to save images in
	ImageOnlineDir string // path to access files online
}

func SuggestGET(w http.ResponseWriter, r *http.Request) {
	templ.Handler(components.SuggestIndex()).ServeHTTP(w, r)
}

func SuggestPOST(db *sql.DB, c *CreateRouteConfig) func(w http.ResponseWriter, r *http.Request) {
	var maxMemory int64 = 32 << 20 // 32MB
	imageSaveDir := "./data/imgs"
	if c != nil {
		maxMemory = c.MaxMemory
		imageSaveDir = c.ImageSaveDir
	}
	err := os.MkdirAll(imageSaveDir, 0755)
	if err != nil {
		log.Fatalf("failed to create image directory. err = %s", err)
	}

	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseMultipartForm(maxMemory)
		if err != nil {
			http.Error(w, "Failed to parse form", http.StatusBadRequest)
			return
		}

		name := r.PostFormValue("name")
		if len(name) == 0 {
			http.Error(w, "Missing name value", http.StatusBadRequest)
			return
		}

		f, fh, err := r.FormFile("image")
		if err == http.ErrMissingFile {
			http.Error(w, "No image uploaded", http.StatusBadRequest)
			return
		} else if err != nil {
			http.Error(w, fmt.Sprintf("err = %s", err), http.StatusInternalServerError)
			return
		}
		defer f.Close()

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
			exists, err := models.Cards(qm.Where(fmt.Sprintf("%s = ?", models.CardColumns.Token), imageToken)).Exists(r.Context(), db)
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
		b, err := io.ReadAll(f)
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

		card := models.Card{
			Name:     name,
			Token:    imageToken,
			Filename: imageName,
			Accepted: true, // TODO: remove this line
		}
		card.Insert(r.Context(), db, boil.Infer())

		tx.Commit()

		templ.Handler(components.SuggestSuccess()).ServeHTTP(w, r)
	}
}
