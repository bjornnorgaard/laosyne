package domain

import (
	"context"

	"github.com/bjornnorgaard/laosyne/graphql/graph/generated"
	"github.com/bjornnorgaard/laosyne/graphql/graph/model"
	"github.com/bjornnorgaard/laosyne/repository/database"
	"github.com/cockroachdb/errors"
)

type Api struct {
	db database.Queries
}

func (a Api) Mutation() generated.MutationResolver {
	return a
}

func (a Api) Query() generated.QueryResolver {
	return a
}

func NewApi(database database.Queries) *Api {
	return &Api{db: database}
}

func (a Api) AddPath(ctx context.Context, input model.NewPath) (*model.Path, error) {
	created, err := a.db.CreatePath(ctx, input.Path)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create path")
	}

	dto := &model.Path{
		ID:   int(created.ID),
		Path: created.Path,
	}

	return dto, nil
}

func (a Api) Paths(ctx context.Context) ([]*model.Path, error) {
	mediaPaths, err := a.db.GetPaths(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get media paths")
	}

	var dto []*model.Path
	for _, mp := range mediaPaths {
		dto = append(dto, &model.Path{
			ID:      int(mp.ID),
			Path:    mp.Path,
			Created: mp.Created.String(),
			Updated: mp.Updated.String(),
		})
	}

	return dto, nil
}

func (a Api) DeletePath(ctx context.Context, input model.DeletePath) (bool, error) {

	err := a.db.DeletePath(ctx, int64(input.PathID))
	if err != nil {
		return false, errors.Wrap(err, "failed to delete path")
	}
	return true, nil
}
