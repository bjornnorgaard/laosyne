package domain

import (
	"github.com/bjornnorgaard/laosyne/graphql/graph/generated"
	"github.com/bjornnorgaard/laosyne/repository"
	"github.com/bjornnorgaard/laosyne/repository/database"
)

type Api struct {
	db   database.Queries
	gorm repository.GormContext
}

func (a Api) Mutation() generated.MutationResolver {
	return a
}

func (a Api) Query() generated.QueryResolver {
	return a
}

func NewApi(database database.Queries, gormContext repository.GormContext) *Api {
	return &Api{
		db:   database,
		gorm: gormContext,
	}
}
