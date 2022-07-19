package domain

import (
	"context"
	"fmt"

	"github.com/bjornnorgaard/laosyne/backend/graphql/graph/model"
	"github.com/bjornnorgaard/laosyne/backend/repository/database"
)

func (a API) DislikePicture(_ context.Context, pictureID int) (*model.Picture, error) {
	var pic database.Picture
	a.db.First(&pic, pictureID)
	if pic.ID == 0 {
		return nil, fmt.Errorf("no pic with id %d", pictureID)
	}

	pic.Likes--

	a.db.Save(&pic)

	dto := mapPic(pic)
	return dto, nil
}
