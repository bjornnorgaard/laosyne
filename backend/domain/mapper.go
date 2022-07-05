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
	query := a.db.QueryPictures().Limit(100)

	if input == nil {
		return query
	}

	if input.SortOrder != nil {
		switch *input.SortOrder {
		case model.SortOrderRandom:
			query = query.Order("RANDOM()")
		case model.SortOrderRatingDesc:
			query = query.Order("rating desc")
		case model.SortOrderRatingAsc:
			query = query.Order("rating asc")
		}
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

	if input.Skip != nil {
		query = query.Offset(*input.Skip)
	}

	if input.Take != nil {
		query = query.Limit(*input.Take)
	}

	return query
}
