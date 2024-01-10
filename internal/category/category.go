package category

import (
	"compare/components"
	storage "compare/internal"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	mrand "math/rand"
	"net/http"
	"slices"
	"time"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
)

func CategoryRouter(c *storage.Category) http.Handler {
	r := chi.NewRouter()
	r.Get("/", BattleGET(c))
	r.Post("/card/{token:\\w+}/{index:\\d+}", BattlePOST(c))
	r.Get("/leaderboard", Leaderboard(c))
	return r
}

// Generate a random token of a given length
func GenerateToken(len int) (string, error) {
	b := make([]byte, len)
	if _, err := rand.Read(b); err != nil {
		return "", fmt.Errorf("")
	}
	return hex.EncodeToString(b), nil
}

func getRandomBattle(c *storage.Category) (*storage.Battle, error) {
	c.AllCardsMutex.RLock()
	if len(c.AllCards) <= 1 {
		c.AllCardsMutex.RUnlock()
		panic("Not enough cards to choose 2 unique ones")
	}
	index1 := mrand.Intn(len(c.AllCards))
	card1 := *c.AllCards[index1]
	var card2 storage.BattleCard
	for {
		index2 := mrand.Intn(len(c.AllCards))
		if index2 != index1 {
			card2 = *c.AllCards[index2]
			break
		}
	}
	c.AllCardsMutex.RUnlock()
	for {
		token, err := GenerateToken(30)
		if err != nil {
			return nil, err
		}
		c.ActiveBattlesMutex.RLock()
		_, exists := c.ActiveBattles[token]
		c.ActiveBattlesMutex.RUnlock()
		if !exists {
			// TODO: Add battle to db to recover on restart
			battle := storage.Battle{
				Start:         time.Now(),
				Card1:         card1,
				Card2:         card2,
				Token:         token,
				CategoryToken: c.Token,
			}
			c.ActiveBattlesMutex.Lock()
			c.ActiveBattles[token] = battle
			c.ActiveBattlesMutex.Unlock()
			return &battle, nil
		}
	}
}

func getSortedCards(c *storage.Category) []storage.BattleCard {
	sortedCards := []storage.BattleCard{}
	c.AllCardsMutex.RLock()
	for _, v := range c.AllCards {
		sortedCards = append(sortedCards, *v)
	}
	c.AllCardsMutex.RUnlock()

	slices.SortStableFunc(sortedCards, func(a, b storage.BattleCard) int {
		return b.Wins - a.Wins
	})
	return sortedCards
}

func BattleGET(c *storage.Category) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: get random cards from db
		battle, err := getRandomBattle(c)
		if err != nil {
			log.Printf("getRandomBattle err = %s", err)
			return
		}
		sortedCards := getSortedCards(c)
		templ.Handler(components.CategoryIndex(c.Title, *battle, sortedCards)).ServeHTTP(w, r)
	}
}

func BattlePOST(c *storage.Category) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: remove this. simulates a slow query
		time.Sleep(500 * time.Millisecond)

		// Log results
		token := chi.URLParam(r, "token")
		index := chi.URLParam(r, "index")

		c.ActiveBattlesMutex.RLock()
		battle, ok := c.ActiveBattles[token]
		c.ActiveBattlesMutex.RUnlock()

		if index != "0" && index != "1" {
			http.Error(w, "invalid index", http.StatusBadRequest)
			return
		}

		// if !ok, token might be too old. Silently fail and start a new battle
		if ok {
			c.AllCardsMutex.Lock()
			if index == "0" {
				c.AllCards[battle.Card1.ID].Wins++
				c.AllCards[battle.Card2.ID].Losses++
			} else {
				c.AllCards[battle.Card1.ID].Losses++
				c.AllCards[battle.Card2.ID].Wins++
			}
			c.AllCardsMutex.Unlock()

			// clear battle
			c.ActiveBattlesMutex.Lock()
			delete(c.ActiveBattles, token)
			c.ActiveBattlesMutex.Unlock()
		}

		// Start a new battle
		// TODO: get random cards from db
		newBattle, err := getRandomBattle(c)
		if err != nil {
			log.Printf("getRandomBattle err = %s", err)
			http.Error(w, "failed starting a new battle", http.StatusBadRequest)
			return
		}
		templ.Handler(components.Battle(*newBattle)).ServeHTTP(w, r)
	}
}

func Leaderboard(c *storage.Category) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		sortedCards := getSortedCards(c)
		templ.Handler(components.Leaderboard(sortedCards)).ServeHTTP(w, r)
	}
}
