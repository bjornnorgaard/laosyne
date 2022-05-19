package domain

import (
	"context"

	"github.com/bjornnorgaard/laosyne/graphql/graph/generated"
	"github.com/bjornnorgaard/laosyne/graphql/graph/model"
	"github.com/bjornnorgaard/laosyne/postgres/database"
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

func (a Api) AddMediaPath(ctx context.Context, input model.NewMediaPath) (*model.MediaPath, error) {
	created, err := a.db.CreateMediaPath(ctx, input.Path)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create media path")
	}

	dto := &model.MediaPath{
		ID:   created.ID.String(),
		Path: created.Path,
	}

	return dto, nil
}

func (a Api) MediaPaths(ctx context.Context) ([]*model.MediaPath, error) {
	mediaPaths, err := a.db.GetMediaPaths(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get media paths")
	}

	var dto []*model.MediaPath
	for _, mp := range mediaPaths {
		dto = append(dto, &model.MediaPath{
			ID:   mp.ID.String(),
			Path: mp.Path,
		})
	}

	return dto, nil
}
