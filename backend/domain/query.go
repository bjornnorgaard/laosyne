package domain

import (
	"fmt"
	"strings"

	"github.com/bjornnorgaard/laosyne/backend/graphql/graph/model"
	"gorm.io/gorm"
)

func (a API) buildQuery(input *model.SearchFilter) *gorm.DB {
	query := a.db.QueryPictures().Limit(100)

	if input == nil {
		return query
	}

	if input.SortOrder != nil {
		switch *input.SortOrder {
		case model.SortOrderID:
			query = query.Order("id asc")
		case model.SortOrderRandom:
			query = query.Order("RANDOM()")
		case model.SortOrderRatingDesc:
			query = query.Order("rating desc")
		case model.SortOrderRatingAsc:
			query = query.Order("rating asc")
		case model.SortOrderViewsDesc:
			query = query.Order("views desc")
		case model.SortOrderViewsAsc:
			query = query.Order("views asc")
		case model.SortOrderLikesDesc:
			query = query.Order("likes desc")
		case model.SortOrderLikesAsc:
			query = query.Order("likes asc")
		case model.SortOrderCreatedAtDesc:
			query = query.Order("created_at desc")
		case model.SortOrderCreatedAtAsc:
			query = query.Order("created_at asc")
		case model.SortOrderUpdatedAtDesc:
			query = query.Order("updated_at desc")
		case model.SortOrderUpdatedAtAsc:
			query = query.Order("updated_at asc")
		}
	}

	if input.PathContains != nil {
		for _, s := range strings.Split(*input.PathContains, " ") {
			query = query.Where("path LIKE ?", fmt.Sprintf("%%%s%%", s))
		}
	}

	if input.LowerRating != nil {
		query = query.Where("? <= rating", input.LowerRating)
	}

	if input.UpperRating != nil {
		query = query.Where("rating <= ?", input.UpperRating)
	}

	if input.Skip != nil {
		query = query.Offset(*input.Skip)
	}

	if input.Take != nil {
		query = query.Limit(*input.Take)
	}

	return query
}
