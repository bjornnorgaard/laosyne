package domain

import (
	"context"

	"github.com/bjornnorgaard/laosyne/backend/graphql/graph/model"
	"github.com/bjornnorgaard/laosyne/backend/repository/database"
)

func (a API) AddPath(_ context.Context, input model.NewPath) (*model.Path, error) {
	path := database.Path{Path: input.Path}
	a.db.Create(&path)

	dto := &model.Path{
		ID:        int(path.ID),
		Path:      path.Path,
		CreatedAt: path.CreatedAt.String(),
	}

	return dto, nil
}
