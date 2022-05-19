package repository

import (
	"database/sql"

	"github.com/cockroachdb/errors"
	"github.com/golang-migrate/migrate/v4"
	pq "github.com/golang-migrate/migrate/v4/database/postgres"
)

const (
	// errNoChange returned when migrations contain no changes.
	errNoChange = "no change"
)

func Migrate(db *sql.DB) error {
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
