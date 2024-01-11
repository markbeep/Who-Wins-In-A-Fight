package main

import (
	"compare/components"
	storage "compare/internal"
	"compare/internal/category"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httplog"
	"github.com/sirupsen/logrus"
)

var (
	port = flag.String("port", os.Getenv("PORT"), "port to host the website at")

	// Logging
	logLevel = flag.String("log-level", "INFO", "one of: TRACE, DEBUG, INFO, WARN, ERROR")

	// Postgres DB
	dbUser     = flag.String("db-user", os.Getenv("POSTGRES_USER"), "user of the database")
	dbPassword = flag.String("db-password", os.Getenv("POSTGRES_PASSWORD"), "password for the database")
	dbName     = flag.String("db-name", os.Getenv("POSTGRES_DB"), "name of the database")
	dbPort     = flag.String("db-port", os.Getenv("POSTGRES_DB_PORT"), "port of the database")
	dbHost     = flag.String("db-host", os.Getenv("POSTGRES_DB_SERVER"), "host of the database")
)

// TODO: store in db
var categories = []storage.Category{
	{
		Token: "123", // TODO: this should be a token stored in the db and regeneratable (incase of leak)
		Title: "Who's stronger?",
		AllCards: map[int]*storage.BattleCard{
			0: {Url: "/static/chuck.png", ID: 0, Name: "Chuck Norris"},
			1: {Url: "/static/superman.jpg", ID: 1, Name: "Superman"},
			2: {Url: "/static/kermit.jpeg", ID: 2, Name: "Kermit the Gangsta Frog"}},
		AllCardsMutex: sync.RWMutex{},
		// TODO: Fill this db with relevant battles from the db on startup
		ActiveBattles:      map[string]storage.Battle{},
		ActiveBattlesMutex: sync.RWMutex{},
	},
}

func main() {
	flag.Parse()
	if *port == "" {
		*port = "3000"
	}

	// TODO: involve the logger more
	logger := logrus.New()
	logger.SetReportCaller(true)
	logger.SetLevel(func(l string) logrus.Level {
		lvl, err := logrus.ParseLevel(l)
		if err != nil {
			return logrus.TraceLevel
		}
		return lvl
	}(*logLevel))

	// Connects to the db and handles migrations
	db, err := storage.GetDB(*dbUser, *dbPassword, *dbName, *dbHost, *dbPort)
	if err != nil {
		logger.Fatal(err)
	}
	defer db.Close()
	err = storage.MigrateDatabase(db, storage.MigrationLogger{Logger: logger, LogLevel: *logLevel}, "migrations")
	if err != nil {
		logger.Fatal(err)
	}

	// Webserver
	r := chi.NewRouter()
	r.Use(MiddlewareLogging)

	r.Get("/", templ.Handler(components.Index()).ServeHTTP)
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("healthy")) })
	r.Get("/static/*", Static)
	for i := range categories {
		c := &categories[i]
		r.Mount("/"+c.Token, category.CategoryRouter(c))
	}
	r.Mount("/create", category.CreateRoute())

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
