package repository

import (
	"fmt"
	"log"

	"github.com/bjornnorgaard/laosyne/backend/repository/database"
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
		log.Fatal(fmt.Errorf("failed to open database connection, %w", err))
	}

	err = db.AutoMigrate(&database.Picture{}, &database.Path{})
	if err != nil {
		log.Fatal(fmt.Errorf("failed to migrate database: %w", err))
	}

	return Repository{db}
}
