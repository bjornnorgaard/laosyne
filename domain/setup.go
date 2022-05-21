package domain

import (
	"github.com/bjornnorgaard/laosyne/graphql/graph/generated"
	"github.com/bjornnorgaard/laosyne/repository"
)

type Api struct {
	db repository.Repository
}

func (a Api) Mutation() generated.MutationResolver {
	return a
}

func (a Api) Query() generated.QueryResolver {
	return a
}

func NewApi(r repository.Repository) *Api {
	return &Api{
		db: r,
	}
}
