package storage

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	"github.com/sirupsen/logrus"

	// required drivers for pg db
	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/lib/pq"
)

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
		return fmt.Errorf("failed to open driver. err = %s", err)
	}

	logger.Printf("running migrations")
	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", migrationPath), "postgres", driver,
	)
	if err != nil {
		return fmt.Errorf("failed to create new database instance. err = %s", err)
	}
	m.Log = logger
	if err = m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed to migrate. err = %s", err)
	}
	if err == migrate.ErrNoChange {
		logger.Printf("no migrations to apply")
	} else {
		logger.Printf("migrations successfully applied")
	}
	return nil
}
