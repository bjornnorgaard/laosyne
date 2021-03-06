package domain

import (
	"context"

	"github.com/bjornnorgaard/laosyne/backend/graphql/graph/model"
	"github.com/bjornnorgaard/laosyne/backend/repository/database"
)

func (a API) Pictures(_ context.Context, input *model.SearchFilter) ([]*model.Picture, error) {
	var pics []database.Picture
	a.buildQuery(input).Find(&pics)
	if len(pics) == 0 {
		return []*model.Picture{}, nil
	}

	var dto []*model.Picture
	for _, p := range pics {
		dto = append(dto, &model.Picture{
			ID:        int(p.ID),
			Path:      p.Path,
			Ext:       p.Ext,
			Views:     p.Views,
			Likes:     p.Likes,
			Rating:    p.Rating,
			Deviation: p.Deviation,
			Wins:      p.Wins,
			Losses:    p.Losses,
			CreatedAt: p.CreatedAt.String(),
			UpdatedAt: p.UpdatedAt.String(),
		})
	}

	return dto, nil
}
