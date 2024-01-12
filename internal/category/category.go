package category

import (
	"compare/components"
	"compare/models"
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// Generate a random token of a given length.
// len is the amount of random bytes to generate.
func GenerateToken(len int) (string, error) {
	b := make([]byte, len)
	if _, err := rand.Read(b); err != nil {
		return "", fmt.Errorf("failed generating a random token. err = %s", err)
	}
	return strings.ReplaceAll(base64.URLEncoding.EncodeToString(b), "=", ""), nil
}

func getRandomBattle(ctx context.Context, db *sql.DB) (*models.Battle, error) {
	card1, err := models.Cards(
		qm.OrderBy("RANDOM()"),
	).One(ctx, db)
	if err != nil {
		return nil, err
	}

	card2, err := models.Cards(
		qm.Where(fmt.Sprintf("%s != ?", models.CardColumns.ID), card1.ID),
		qm.OrderBy("RANDOM()"),
	).One(ctx, db)
	if err != nil {
		return nil, err
	}

	for {
		token, err := GenerateToken(10)
		if err != nil {
			return nil, err
		}
		exists, err := models.Battles(qm.Where(fmt.Sprintf("%s = ?", models.BattleColumns.Token), token)).Exists(ctx, db)
		if err != nil {
			return nil, err
		}
		if !exists {
			battle := models.Battle{
				Card1ID: card1.ID,
				Card2ID: card2.ID,
				Token:   token,
			}
			if err = battle.Insert(ctx, db, boil.Infer()); err != nil {
				return nil, err
			}
			if err = battle.Reload(ctx, db); err != nil {
				return nil, err
			}
			if err = battle.SetCard1(ctx, db, false, card1); err != nil {
				return nil, err
			}
			if err = battle.SetCard2(ctx, db, false, card2); err != nil {
				return nil, err
			}
			return &battle, nil
		}
	}
}

// func getSortedCards(c *storage.Category) []storage.BattleCard {
// 	sortedCards := []storage.BattleCard{}
// 	c.AllCardsMutex.RLock()
// 	for _, v := range c.AllCards {
// 		sortedCards = append(sortedCards, *v)
// 	}
// 	c.AllCardsMutex.RUnlock()

// 	slices.SortStableFunc(sortedCards, func(a, b storage.BattleCard) int {
// 		return b.Wins - a.Wins
// 	})
// 	return sortedCards
// }

func BattleGET(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: get random cards from db
		battle, err := getRandomBattle(r.Context(), db)
		if err != nil {
			if strings.Contains(err.Error(), "no rows in result set") {
				templ.Handler(components.EmptyIndex()).ServeHTTP(w, r)
			}
			log.Printf("getRandomBattle err = %s", err)
			http.Error(w, "failed to create battle", http.StatusInternalServerError)
			return
		}
		templ.Handler(components.Index(*battle)).ServeHTTP(w, r)
	}
}

func BattlePOST(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: remove this. simulates a slow query
		time.Sleep(200 * time.Millisecond)

		token := chi.URLParam(r, "token")
		index := chi.URLParam(r, "index")

		if index != "0" && index != "1" {
			http.Error(w, "invalid index", http.StatusBadRequest)
			return
		}

		tx, err := db.BeginTx(r.Context(), nil)
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to start transaction. err = %s", err), http.StatusBadRequest)
			return
		}

		// Check if its a valid battle
		battle, err := models.Battles(
			qm.Load(models.BattleRels.Card1),
			qm.Load(models.BattleRels.Card2),
			qm.Where(fmt.Sprintf("%s = ?", models.BattleColumns.Token), token),
		).One(r.Context(), db)
		if err != nil || battle == nil {
			tx.Rollback()
			http.Error(w, "invalid token", http.StatusBadRequest)
			return
		}

		card1 := battle.R.Card1
		card2 := battle.R.Card2
		card1.Battles++
		card2.Battles++
		if index == "0" {
			card1.Wins++
		} else {
			card2.Wins++
		}
		card1.Update(r.Context(), db, boil.Infer())
		card2.Update(r.Context(), db, boil.Infer())

		tx.Commit()

		// Start a new battle
		newBattle, err := getRandomBattle(r.Context(), db)
		if err != nil {
			log.Printf("getRandomBattle err = %s", err)
			http.Error(w, "failed starting a new battle", http.StatusBadRequest)
			return
		}
		templ.Handler(components.Battle(*newBattle)).ServeHTTP(w, r)
	}
}

// func leaderboard(c *storage.Category) func(w http.ResponseWriter, r *http.Request) {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		sortedCards := getSortedCards(c)
// 		templ.Handler(components.Leaderboard(sortedCards)).ServeHTTP(w, r)
// 	}
// }
