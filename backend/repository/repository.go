package repository

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/bjornnorgaard/laosyne/backend/repository/database"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repository struct {
	*gorm.DB
}

func NewRepository() Repository {
	dsn := os.ExpandEnv("host=${DB_HOST} user=${DB_USER} password=${DB_PASS} dbname=${DB_NAME} port=${DB_PORT} sslmode=disable")
	db, err := openConnection(dsn)
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&database.Picture{}, &database.Path{}, &database.Match{})
	if err != nil {
		log.Fatal(err)
	}

	return Repository{db}
}

func openConnection(dsn string) (*gorm.DB, error) {
	open, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	retries := 0
	for err != nil {
		time.Sleep(time.Second * 1)
		retries++
		if 2 <= retries {
			break
		}

		open, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}
	}

	return open, err
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
