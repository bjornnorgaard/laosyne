package main

import (
	"database/sql"
	"log"

	"github.com/bjornnorgaard/laosyne/domain"
	"github.com/bjornnorgaard/laosyne/graphql"
	"github.com/bjornnorgaard/laosyne/postgres/database"
	"github.com/cockroachdb/errors"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "user=postgres password=changeme dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal(errors.Wrap(err, "application failed to start"))
	}

	queries := database.New(db)
	api := domain.NewApi(*queries)
	graphql.Start(api)
}
