package repository

import (
	"log"

	"github.com/bjornnorgaard/laosyne/repository/database"
	"github.com/cockroachdb/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repository struct {
	*gorm.DB
}

func NewRepository() Repository {
	dsn := "host=localhost user=postgres password=changeme dbname=postgres port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to open database connection"))
	}

	err = db.AutoMigrate(&database.Picture{}, &database.Path{})
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to migrate database"))
	}

	return Repository{db}
}
