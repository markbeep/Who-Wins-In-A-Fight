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
	"sync"
	"time"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

const DEBUG = false

// TODO: clear out old battles after some time
var battleCache = map[string]internal.Battle{}
var cacheMutex = sync.RWMutex{}

func getRandomBattle(ctx context.Context, db *sql.DB) (*internal.Battle, error) {
	card1, err := models.Cards(
		qm.Where(fmt.Sprintf("%s = true", models.CardColumns.Accepted)),
		qm.OrderBy("RANDOM()"),
	).One(ctx, db)
	if err != nil {
		return nil, err
	}

	card2, err := models.Cards(
		qm.Where(fmt.Sprintf("%s != ? AND %s = true", models.CardColumns.ID, models.CardColumns.Accepted), card1.ID),
		qm.OrderBy("RANDOM()"),
	).One(ctx, db)
	if err != nil {
		return nil, err
	}

	// Make sure card1.ID < card2.ID
	if card1.ID > card2.ID {
		card1, card2 = card2, card1
	}

	battle, err := models.Battles(
		qm.Where(fmt.Sprintf("%s = ? AND %s = ?", models.BattleColumns.Card1ID, models.BattleColumns.Card2ID), card1.ID, card2.ID),
	).One(ctx, db)
	if err == sql.ErrNoRows {
		battle = &models.Battle{
			Card1ID: card1.ID,
			Card2ID: card2.ID,
		}
		if err = battle.Insert(ctx, db, boil.Infer()); err != nil {
			return nil, err
		}
		if err = battle.Reload(ctx, db); err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err
	}

	var token string
	for {
		token, err = internal.GenerateToken(10)
		if err != nil {
			return nil, err
		}
		cacheMutex.RLock()
		_, ok := battleCache[token]
		cacheMutex.RUnlock()
		if !ok {
			break
		}
	}

	total := float32(battle.Card1Wins + battle.Card2Wins)
	if total == 0.0 {
		total++
	}
	localBattle := internal.Battle{
		ID:          battle.ID,
		Card1:       card1,
		Card2:       card2,
		Card1Chance: float32(battle.Card1Wins) / total,
		Card2Chance: float32(battle.Card2Wins) / total,
		Token:       token,
		Start:       time.Now(),
	}
	cacheMutex.Lock()
	battleCache[token] = localBattle
	cacheMutex.Unlock()

	return &localBattle, nil
}

func BattleGET(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		battle, err := getRandomBattle(r.Context(), db)
		if err == sql.ErrNoRows {
			templ.Handler(components.Index(nil, DEBUG)).ServeHTTP(w, r)
			return
		} else if err != nil {
			http.Error(w, fmt.Sprintf("failed to create battle. err = %s", err), http.StatusInternalServerError)
			return
		}
		templ.Handler(components.Index(battle, DEBUG)).ServeHTTP(w, r)
	}
}

func BattlePOST(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		token := chi.URLParam(r, "token")
		index := chi.URLParam(r, "index")

		if index != "0" && index != "1" {
			http.Error(w, "invalid index", http.StatusBadRequest)
			return
		}

		// Check if its a valid battle
		cacheMutex.RLock()
		localBattle, ok := battleCache[token]
		cacheMutex.RUnlock()

		// Silently fail if token is invalid and start a new battle
		if ok {
			delete(battleCache, token)

			tx, err := db.BeginTx(r.Context(), nil)
			if err != nil {
				http.Error(w, fmt.Sprintf("failed to start transaction. err = %s", err), http.StatusInternalServerError)
				return
			}

			battle, err := models.Battles(
				qm.Load(models.BattleRels.Card1),
				qm.Load(models.BattleRels.Card2),
				qm.Where(fmt.Sprintf("%s = ?", models.BattleColumns.ID), localBattle.ID),
			).One(r.Context(), tx)
			if err != nil {
				http.Error(w, fmt.Sprintf("failed to retreive battle. err = %s", err), http.StatusInternalServerError)
				return
			}

			card1 := battle.R.Card1
			card2 := battle.R.Card2
			card1.Battles++
			card2.Battles++
			if index == "0" {
				card1.Wins++
				battle.Card1Wins++
			} else {
				card2.Wins++
				battle.Card2Wins++
			}
			card1.Update(r.Context(), tx, boil.Infer())
			card2.Update(r.Context(), tx, boil.Infer())
			battle.Update(r.Context(), tx, boil.Infer())

			err = tx.Commit()
			if err != nil {
				http.Error(w, fmt.Sprintf("failed to commit. err = %s", err), http.StatusInternalServerError)
				return
			}
		}

		// Start a new battle
		newBattle, err := getRandomBattle(r.Context(), db)
		if err != nil {
			log.Printf("getRandomBattle err = %s", err)
			http.Error(w, "failed starting a new battle", http.StatusBadRequest)
			return
		}
		templ.Handler(components.Battle(newBattle, DEBUG)).ServeHTTP(w, r)
	}
}
