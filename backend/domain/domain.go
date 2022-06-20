package domain

import (
	"github.com/bjornnorgaard/laosyne/backend/graphql/graph/generated"
	"github.com/bjornnorgaard/laosyne/backend/repository"
)

type API struct {
	db repository.Repository
}

func (a API) Mutation() generated.MutationResolver {
	return a
}

func (a API) Query() generated.QueryResolver {
	return a
}

func NewApi(r repository.Repository) *API {
	return &API{
		db: r,
	}
}
