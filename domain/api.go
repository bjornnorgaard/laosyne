package domain

import (
	"context"

	"github.com/bjornnorgaard/laosyne/graphql/graph/model"
	"github.com/bjornnorgaard/laosyne/repository/database"
	"github.com/cockroachdb/errors"
)

type Api struct {
	db database.Queries
}

func (a Api) Picture(ctx context.Context, filter string) (*model.Picture, error) {
	picture, err := a.db.GetPictureByFilter(ctx, filter)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get pic by filter: '%s'", filter)
	}

	dto := &model.Picture{
		ID:        int(picture.ID),
		Path:      picture.Path,
		Ext:       picture.Ext,
		Views:     int(picture.Views),
		Likes:     int(picture.Likes),
		Rating:    picture.Rating,
		Deviation: picture.Deviation,
		Wins:      int(picture.Wins),
		Losses:    int(picture.Losses),
		Created:   picture.Created.String(),
		Updated:   picture.Updated.String(),
	}

	return dto, nil
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

func (a Api) ScanPath(ctx context.Context) (bool, error) {
	// TODO implement me
	panic("implement me")
}
