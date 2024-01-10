package main

import (
	"compare/components"
	storage "compare/internal"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httplog"
)

var (
	port = flag.String("port", os.Getenv("PORT"), "port to host the website at")
)

var cards = []storage.BattleCard{
	{Url: "/static/chuck.png", ID: 0, Name: "Chuck Norris"},
	{Url: "/static/superman.jpg", ID: 1, Name: "Superman"},
	{Url: "/static/kermit.jpeg", ID: 2, Name: "Kermit the Gangsta Frog"},
}

func main() {
	flag.Parse()
	if *port == "" {
		*port = "3000"
	}

	r := chi.NewRouter()

	r.Get("/", BattleGET)
	r.Post("/card/{id1:\\d+}/{id2:\\d+}", BattlePOST)
	r.Get("/static/*", Static)

	host := fmt.Sprintf(":%s", *port)
	log.Printf("listening on %s", host)
	log.Fatal(http.ListenAndServe(host, r))
}

func MiddlewareLogging(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/health" { // don't log /health
			next.ServeHTTP(w, r)
		} else {
			logger := httplog.NewLogger("htmx-blog", httplog.Options{
				LogLevel: "warn",
			})
			httplog.RequestLogger(logger)(next).ServeHTTP(w, r)
		}
	}
	return http.HandlerFunc(fn)
}

func Static(w http.ResponseWriter, r *http.Request) {
	path := strings.Trim(r.URL.Path, "/")
	fileInfo, err := os.Stat(path)
	if err != nil || fileInfo.IsDir() {
		w.Write([]byte("404"))
		return
	}
	http.ServeFile(w, r, path)
}

func getRandomBattle() (storage.BattleCard, storage.BattleCard) {
	index1 := rand.Intn(len(cards))
	card1 := cards[index1]
	for {
		index2 := rand.Intn(len(cards))
		if index2 != index1 {
			return card1, cards[index2]
		}
	}
}

func BattleGET(w http.ResponseWriter, r *http.Request) {
	// TODO: get random cards from db
	card1, card2 := getRandomBattle()
	templ.Handler(components.Index("Who's stronger?", card1, card2, cards)).ServeHTTP(w, r)
}

func BattlePOST(w http.ResponseWriter, r *http.Request) {
	// TODO: remove this. simulates a slow query
	time.Sleep(500 * time.Millisecond)

	// Log results
	id, err := strconv.Atoi(chi.URLParam(r, "id1"))
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	// TODO: make this fetch and update the db entry
	var winningCard *storage.BattleCard
	found := false
	for i := range cards {
		if cards[i].ID == id {
			winningCard = &cards[i]
			found = true
		}
	}
	if !found {
		log.Println("ERROR! Winning card not found")
		return
	}

	// update loser
	id, err = strconv.Atoi(chi.URLParam(r, "id2"))
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
	}
	var losingCard *storage.BattleCard
	found = false
	for i := range cards {
		if cards[i].ID == id {
			losingCard = &cards[i]
			found = true
		}
	}
	if !found {
		log.Println("ERROR! Losing card not found")
		return
	}

	winningCard.Wins += 1
	log.Printf("%d won!\n", winningCard.ID)
	losingCard.Losses += 1
	log.Printf("%d lost!\n", losingCard.ID)

	// Start a new battle
	// TODO: get random cards from db
	card1, card2 := getRandomBattle()
	templ.Handler(components.Battle(card1, card2)).ServeHTTP(w, r)
}

func Leaderboard(w http.ResponseWriter, r *http.Request) {
	slices.SortFunc(cards, func(a, b storage.BattleCard) int {
		return 0
	})
}
