package domain

import (
	"fmt"
	"strings"

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
		for _, s := range strings.Split(*input.PathContains, " ") {
			query = query.Where("path LIKE ?", fmt.Sprintf("%%%s%%", s))
		}
	}

	if input.LowerRating != nil {
		query = query.Where("? < rating", input.LowerRating)
	}

	if input.UpperRating != nil {
		query = query.Where("rating < ?", input.UpperRating)
	}

	return query
}
