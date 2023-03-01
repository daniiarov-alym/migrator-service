package db

import (
	"context"
	"database/sql"
	"github.com/daniiarov-alym/migrator-service/src/config"
	"fmt"
	"strings"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/jackc/pgx/v4/pgxpool"
	logger "github.com/sirupsen/logrus"
)

func Run(ctx context.Context) {
	createDatabase(ctx)
	migrateDb(ctx)
}


func createDatabase(ctx context.Context) {
	conf := config.Conf()
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/postgres?sslmode=disable",
		conf.User, conf.Password,
		conf.Host, conf.Port)
	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		logger.WithError(err).Fatal("Unable to parse database config")
	}
	repo, err := pgxpool.ConnectConfig(ctx, config)
	if err != nil {
		logger.Fatalf("Failed to initialise connection to Postgres Database: %s", err)
	}
	_, err = repo.Exec(ctx, "CREATE DATABASE \""+conf.Database+"\"")
	if err != nil && !strings.Contains(err.Error(), "(SQLSTATE 42P04)") {
		logger.Fatal("Failed to create a database: " + err.Error())
	}
}


func migrateDb(ctx context.Context) {
	conf := config.Conf()
	logger.Trace("Starting migrations")
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		conf.User, conf.Password,
		conf.Host, conf.Port, conf.Database)
	dbCon, err := sql.Open("postgres", dsn)
	if err != nil {
		logger.Fatal("failed to connect to db :: ", err)
	}

	driver, err := postgres.WithInstance(dbCon, &postgres.Config{})
	if err != nil {
		logger.Fatal("failed to create postgres abstraction :: ", err)
	}

	srcUrl := "file://migrations"
	migrator, err := migrate.NewWithDatabaseInstance(srcUrl, conf.Database, driver)
	if err != nil {
		logger.Fatal("failed to create migrator instance :: ", err)
	}

	err = migrator.Up()
	if err != nil && err.Error() != "no change" {
		logger.Fatal("failed to apply migrations :: ", err)
	}
	logger.Trace("Migration finished successfully")
}
