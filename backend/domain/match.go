package domain

import (
	"context"
	"math/rand"

	"github.com/bjornnorgaard/laosyne/backend/graphql/graph/model"
	"github.com/bjornnorgaard/laosyne/backend/repository/database"
)

func (a API) Match(_ context.Context, input *model.SearchFilter) (*model.Match, error) {
	var challenger database.Picture

	a.buildQuery(input).
		Where("rating <> 0").
		Where("wins < 20").
		Where("losses < 20").
		Where("2 < deviation").
		Order("RANDOM()").
		First(&challenger)

	if challenger.ID == 0 {
		a.buildQuery(input).
			Where("rating <> 0").
			Order("RANDOM()").
			First(&challenger)
	}

	opponent := a.findOpponent(challenger)

	match := &model.Match{
		PlayerOne: mapPic(challenger),
		PlayerTwo: mapPic(opponent),
	}

	if rand.Intn(100)%2 == 0 {
		tempOne := match.PlayerOne
		match.PlayerOne = match.PlayerTwo
		match.PlayerTwo = tempOne
	}

	return match, nil
}

func (a API) findOpponent(challenger database.Picture) database.Picture {
	var opponent database.Picture

	a.db.QueryPictures().
		Where("id != ?", challenger.ID).
		Where("rating <> 0").
		Order("RANDOM()").
		First(&opponent)

	return opponent
}
