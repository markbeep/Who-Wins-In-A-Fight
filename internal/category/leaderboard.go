package category

import (
	"compare/components"
	"compare/models"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func LeaderboardGET(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		cards, err := models.Cards(
			qm.Where(fmt.Sprintf("%s = true", models.CardColumns.Accepted)),
			qm.OrderBy(fmt.Sprintf("%s DESC", models.CardColumns.Wins)),
			qm.Limit(20),
		).All(r.Context(), db)

		if err != nil {
			http.Error(w, fmt.Sprintf("unable to fetch cards. err = %s", err), http.StatusInternalServerError)
			return
		}

		templ.Handler(components.LeaderboardIndex(cards)).ServeHTTP(w, r)
	}
}
