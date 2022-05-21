package main

import (
	"github.com/bjornnorgaard/laosyne/domain"
	"github.com/bjornnorgaard/laosyne/graphql"
	"github.com/bjornnorgaard/laosyne/repository"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {
	r := repository.NewRepository()
	api := domain.NewApi(r)
	graphql.Start(api)
}
