package main

import (
	"compare/internal/category"
	"compare/internal/storage"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

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

	// Creation envs
	maxFileSize       = flag.String("max-filesize", os.Getenv("MAX_FILESIZE"), "maximum filesize to store in memory (default: 10MB)")
	imageSaveDir      = flag.String("image-dir", os.Getenv("IMAGE_SAVE_DIR"), "directory to store images in (default: ./data/imgs)")
	allowedExtensions = flag.String("allowed-extensions", os.Getenv("ALLOWED_EXTENSIONS"), "the allowed image extensions to be submitted (default: .png .jpg .jpeg)")
)

func main() {
	// Handle envs
	flag.Parse()
	if *port == "" {
		*port = "3000"
	}
	createConfig := category.CreateRouteConfig{
		MaxMemory:         10 << 20, // 10MB
		ImageSaveDir:      "./static/imgs",
		AllowedExtensions: []string{".jpg", ".jpeg", ".png"},
	}
	if *maxFileSize != "" {
		val, err := strconv.ParseInt(*maxFileSize, 10, 64)
		if err != nil {
			panic(fmt.Sprintf("maxFileSize is not an int: err = %s", err))
		}
		createConfig.MaxMemory = val
	}
	if *imageSaveDir != "" {
		createConfig.ImageSaveDir = *imageSaveDir
	}
	if *allowedExtensions != "" {
		createConfig.AllowedExtensions = strings.Split(*allowedExtensions, " ")
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

	r.Get("/", category.BattleGET(db))
	r.Get("/static/*", Static)
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("healthy")) })
	r.Post("/card/{token:[\\w-]+}/{index:\\d+}", category.BattlePOST(db))
	r.Get("/leaderboard", category.LeaderboardGET(db))
	r.Route("/suggest", category.SuggestRoute(db, &createConfig))
	r.Route("/admin", category.AdminRoute(db, &createConfig))

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
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	http.ServeFile(w, r, path)
}
