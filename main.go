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
	maxFileSize  = flag.String("max-filesize", os.Getenv("MAX_FILESIZE"), "maximum filesize to store in memory (default: 32MB)")
	imageSaveDir = flag.String("image-dir", os.Getenv("IMAGE_SAVE_DIR"), "directory to store images in (default: ./data/imgs)")
)

func main() {
	// Handle envs
	flag.Parse()
	if *port == "" {
		*port = "3000"
	}
	createConfig := category.CreateRouteConfig{
		MaxMemory:    32 << 20, // 32MB
		ImageSaveDir: "./static/imgs",
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
	r.Get("/suggest", category.SuggestGET)
	r.Post("/suggest", category.SuggestPOST(db, &createConfig))
	r.Get("/admin", category.AdminGET)
	r.Post("/admin", category.AdminPOST(db))
	r.Get("/admin/{token:[\\w-]+}", category.AdminTokenGET(db))

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
