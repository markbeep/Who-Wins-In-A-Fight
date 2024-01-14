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
	"time"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

const DEBUG = false

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
		token, err := internal.GenerateToken(10)
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

func BattleGET(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		battle, err := getRandomBattle(r.Context(), db)
		if err == sql.ErrNoRows {
			templ.Handler(components.Index(nil, DEBUG)).ServeHTTP(w, r)
			return
		} else if err != nil {
			log.Printf("getRandomBattle err = %s", err)
			http.Error(w, "failed to create battle", http.StatusInternalServerError)
			return
		}
		templ.Handler(components.Index(battle, DEBUG)).ServeHTTP(w, r)

		log.Println("TIME TAKEN", time.Now().Sub(start))
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
		).One(r.Context(), tx)
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
		card1.Update(r.Context(), tx, boil.Infer())
		card2.Update(r.Context(), tx, boil.Infer())

		tx.Commit()

		// Start a new battle
		newBattle, err := getRandomBattle(r.Context(), db)
		if err != nil {
			log.Printf("getRandomBattle err = %s", err)
			http.Error(w, "failed starting a new battle", http.StatusBadRequest)
			return
		}
		templ.Handler(components.Battle(*newBattle, DEBUG)).ServeHTTP(w, r)
	}
}
