package repository

import (
	"fmt"
	"log"
	"os"

	"github.com/bjornnorgaard/laosyne/backend/repository/database"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repository struct {
	*gorm.DB
}

func NewRepository() Repository {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	name := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, pass, name, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(fmt.Errorf("failed to open database connection, %w", err))
	}

	err = db.AutoMigrate(&database.Picture{}, &database.Path{}, &database.Match{})
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
