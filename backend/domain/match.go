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

	opponent := a.findWorthyOpponent(challenger)

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

func (a API) findWorthyOpponent(challenger database.Picture) database.Picture {
	var (
		lastWonMatch  database.Match
		lastLostMatch database.Match
		opponent      database.Picture
	)

	a.db.Where("winner_id = ?", challenger.ID).Order("created_at desc").First(&lastWonMatch)
	a.db.Where("loser_id = ?", challenger.ID).Order("created_at desc").First(&lastLostMatch)

	opponentQuery := a.db.Where("rating <> 0").Where("likes >= 0")

	// Never played.
	if lastWonMatch.ID == 0 && lastLostMatch.ID == 0 {
		return a.findOpponent(challenger)
	}
	// Never won a match.
	if lastWonMatch.ID == 0 && lastLostMatch.ID != 0 {
		opponentQuery = opponentQuery.Where("rating < ?", challenger.Rating)
	}
	// Never lost a match.
	if lastLostMatch.ID == 0 && lastWonMatch.ID != 0 {
		opponentQuery = opponentQuery.Where("? < rating", challenger.Rating)
	}
	// Won and lost matches.
	if lastWonMatch.ID != 0 && lastLostMatch.ID != 0 {
		// Last match was won.
		if lastWonMatch.CreatedAt.After(lastLostMatch.CreatedAt) {
			opponentQuery = opponentQuery.Where("? < rating", challenger.Rating)
		} else {
			opponentQuery = opponentQuery.Where("rating < ?", challenger.Rating)
		}
	}

	opponentQuery.First(&opponent)
	if opponent.ID == 0 {
		return a.findOpponent(challenger)
	}

	return opponent
}
