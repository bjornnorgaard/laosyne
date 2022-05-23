package graphql

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/bjornnorgaard/laosyne/backend/domain"
	"github.com/bjornnorgaard/laosyne/backend/graphql/graph/generated"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rs/cors"
)

const (
	defaultPort = "8080"
)

func Start(api *domain.Api) {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()
	router.Use(cors.Default().Handler, middleware.Logger, middleware.Heartbeat("/hc"), middleware.Recoverer)
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: api}))

	router.Handle("/", playground.Handler("Starwars", "/query"))
	router.Handle("/query", srv)
	router.Handle("/p", api.GetFile())

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":8080", router))
}
