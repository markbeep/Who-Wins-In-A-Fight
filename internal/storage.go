package storage

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	"github.com/sirupsen/logrus"

	// required drivers for pg db
	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/lib/pq"
)

type BattleCard struct {
	Url    string
	ID     int
	Wins   int
	Losses int
	Name   string
}

type Battle struct {
	Start         time.Time
	Card1         BattleCard
	Card2         BattleCard
	Token         string
	CategoryToken string
}

type Category struct {
	Token              string
	Title              string
	AllCards           map[int]*BattleCard
	AllCardsMutex      sync.RWMutex
	ActiveBattles      map[string]Battle
	ActiveBattlesMutex sync.RWMutex
}

// Connects to the database and returns the instance
func GetDB(dbUser, dbPassword, dbName, dbHost, dbPort string) (*sql.DB, error) {
	connStr := fmt.Sprintf("user='%s' password='%s' dbname='%s' host='%s' port='%s' sslmode=disable", dbUser, dbPassword, dbName, dbHost, dbPort)
	log.Printf("connect to database %s@%s:%s/%s", dbUser, dbHost, dbPort, dbName)
	return sql.Open("postgres", connStr)
}

// golang-migrate/migrate requires a logger with a Verbose() function
type MigrationLogger struct {
	Logger   *logrus.Logger
	LogLevel string
}

func (l MigrationLogger) Printf(format string, v ...interface{}) {
	l.Logger.Printf(format, v...)
}
func (l MigrationLogger) Verbose() bool {
	return l.LogLevel == "INFO" || l.LogLevel == "DEBUG" || l.LogLevel == "TRACE"
}

// Applies the relevant migration files over the database
func MigrateDatabase(db *sql.DB, logger migrate.Logger, migrationPath string) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}

	logger.Printf("running migrations")
	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", migrationPath), "postgres", driver,
	)
	if err != nil {
		return err
	}
	m.Log = logger
	if err = m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}
	if err == migrate.ErrNoChange {
		logger.Printf("no migrations to apply")
	} else {
		logger.Printf("migrations successfully applied")
	}
	return nil
}
