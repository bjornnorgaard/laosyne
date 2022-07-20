package main

import (
	"github.com/bjornnorgaard/laosyne/backend/domain"
	"github.com/bjornnorgaard/laosyne/backend/graphql"
	"github.com/bjornnorgaard/laosyne/backend/repository"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	_ = godotenv.Load()

	db := repository.NewRepository()
	api := domain.NewApi(db)
	graphql.Start(api)
}
