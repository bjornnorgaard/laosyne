package main

import (
	"github.com/bjornnorgaard/laosyne/domain"
	"github.com/bjornnorgaard/laosyne/graphql"
	"github.com/bjornnorgaard/laosyne/repository"
	"github.com/bjornnorgaard/laosyne/repository/database"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {
	repo := repository.NewRepository()
	gormContext := repository.NewGormContext(repo.DB)
	queries := database.New(repo)
	api := domain.NewApi(*queries, gormContext)
	graphql.Start(api)
}
