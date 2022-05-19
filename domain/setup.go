package domain

import (
	"github.com/bjornnorgaard/laosyne/graphql/graph/generated"
	"github.com/bjornnorgaard/laosyne/repository/database"
)

func (a Api) Mutation() generated.MutationResolver {
	return a
}

func (a Api) Query() generated.QueryResolver {
	return a
}

func NewApi(database database.Queries) *Api {
	return &Api{db: database}
}
