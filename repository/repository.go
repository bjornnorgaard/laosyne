package repository

import (
	"database/sql"
	"log"

	"github.com/cockroachdb/errors"
	"github.com/golang-migrate/migrate/v4"
	pq "github.com/golang-migrate/migrate/v4/database/postgres"
)

const (
	// errNoChange returned when migrations contain no changes.
	errNoChange = "no change"
)

type repository struct {
	*sql.DB
}

func NewRepository() *repository {
	db, err := sql.Open("postgres", "user=postgres password=changeme dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal(errors.Wrap(err, "application failed to start"))
	}

	err = migrateDatabase(db)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to migrate database"))
	}

	return &repository{DB: db}
}

func migrateDatabase(db *sql.DB) error {
	driver, err := pq.WithInstance(db, &pq.Config{})
	if err != nil {
		return errors.Wrap(err, "failed to get repository instance")
	}

	migrator, err := migrate.NewWithDatabaseInstance("file://repository/migrations", "repository", driver)
	if err != nil {
		return errors.Wrap(err, "failed to create migrator")
	}

	err = migrator.Up()
	if err != nil {
		if err.Error() != errNoChange {
			return errors.Wrap(err, "failed to apply migrations")
		}
	}

	return nil
}
