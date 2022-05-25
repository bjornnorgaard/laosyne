package domain

import (
	"context"
	"fmt"

	"github.com/bjornnorgaard/laosyne/backend/graphql/graph/model"
	"github.com/bjornnorgaard/laosyne/backend/repository/database"
)

func (a Api) GetPicture(_ context.Context, input *model.SearchFilter) (*model.Picture, error) {
	var pic database.Picture
	a.buildQuery(input).Limit(1).First(&pic)

	if pic.ID == 0 {
		return nil, fmt.Errorf("no picture matches filter: '%s'", *input.PathContains)
	}

	dto := mapPic(pic)
	return dto, nil
}
