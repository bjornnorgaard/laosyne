package domain

import (
	"context"
	"fmt"

	"github.com/bjornnorgaard/laosyne/backend/repository/database"
)

func (a API) LikePicture(_ context.Context, pictureID int) (bool, error) {
	var pic database.Picture
	a.db.First(&pic, pictureID)
	if pic.ID == 0 {
		return false, fmt.Errorf("no pic with id %d", pictureID)
	}

	pic.Likes++

	a.db.Save(&pic)
	return true, nil
}
