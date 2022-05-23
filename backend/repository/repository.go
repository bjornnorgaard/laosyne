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

func (r Repository) QueryPictures() *gorm.DB {
	return r.Session(&gorm.Session{})
}

func (r Repository) FindByID(id int) (*database.Picture, error) {
	var pic database.Picture

	r.DB.First(&pic, id)
	if pic.ID == 0 {
		return nil, fmt.Errorf("no picture with ID %d", id)
	}

	return &pic, nil
}
