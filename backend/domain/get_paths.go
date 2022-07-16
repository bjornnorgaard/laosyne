package domain

import (
	"context"

	"github.com/bjornnorgaard/laosyne/backend/graphql/graph/model"
	"github.com/bjornnorgaard/laosyne/backend/repository/database"
)

func (a API) Paths(_ context.Context) ([]*model.Path, error) {
	var paths []database.Path
	a.db.Find(&paths)

	var dto []*model.Path
	for _, mp := range paths {
		dto = append(dto, &model.Path{ID: int(mp.ID), Path: mp.Path, CreatedAt: mp.CreatedAt.String()})
	}

	return dto, nil
}
