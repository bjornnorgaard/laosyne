package domain

import (
	"github.com/bjornnorgaard/laosyne/backend/graphql/graph/model"
	"github.com/bjornnorgaard/laosyne/backend/repository/database"
)

func mapPic(pic database.Picture) *model.Picture {
	dto := &model.Picture{
		ID:        int(pic.ID),
		Path:      pic.Path,
		Ext:       pic.Ext,
		Views:     pic.Views,
		Likes:     pic.Likes,
		Rating:    pic.Rating,
		Deviation: pic.Deviation,
		Wins:      pic.Wins,
		Losses:    pic.Losses,
		CreatedAt: pic.CreatedAt.String(),
		UpdatedAt: pic.UpdatedAt.String(),
	}
	return dto
}
