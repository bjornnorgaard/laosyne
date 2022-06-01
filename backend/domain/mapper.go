package domain

import (
	"fmt"

	"github.com/bjornnorgaard/laosyne/backend/graphql/graph/model"
	"github.com/bjornnorgaard/laosyne/backend/repository/database"
	"gorm.io/gorm"
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

func (a API) buildQuery(input *model.SearchFilter) *gorm.DB {
	query := a.db.QueryPictures()

	if input == nil {
		return query
	}

	if input.PathContains != nil {
		query = query.Where("path LIKE ?", fmt.Sprintf("%%%s%%", *input.PathContains))
	}

	return query
}
