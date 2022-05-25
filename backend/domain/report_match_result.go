package domain

import (
	"context"
	"fmt"

	"github.com/bjornnorgaard/laosyne/backend/graphql/graph/model"
	"github.com/bjornnorgaard/laosyne/backend/repository/database"
	"github.com/bjornnorgaard/laosyne/backend/trueskill"
)

func (a Api) ReportMatchResult(_ context.Context, input model.MatchResult) (bool, error) {
	winner, err := a.db.FindByID(input.WinnerID)
	if err != nil {
		return false, fmt.Errorf("no picture with winner ID %d", input.WinnerID)
	}

	loser, err := a.db.FindByID(input.LoserID)
	if err != nil {
		return false, fmt.Errorf("no picture with loser ID %d", input.WinnerID)
	}

	ts := trueskill.New(trueskill.DrawProbabilityZero())

	playerWinner := trueskill.NewPlayer(winner.Rating, winner.Deviation)
	playerLoser := trueskill.NewPlayer(loser.Rating, loser.Deviation)

	players := []trueskill.Player{playerWinner, playerLoser}
	adjustedSkills, _ := ts.AdjustSkills(players, false)

	playerWinner = adjustedSkills[0]
	playerLoser = adjustedSkills[1]

	winner.Rating = playerWinner.Mean()
	winner.Deviation = playerWinner.Sigma()
	winner.Wins++

	loser.Rating = playerLoser.Mean()
	loser.Deviation = playerLoser.Sigma()
	loser.Losses++

	result := database.Match{
		Quality:  ts.MatchQuality(players),
		WinnerID: winner.ID,
		LoserID:  loser.ID,
	}

	a.db.Create(&result)
	a.db.Save(&winner)
	a.db.Save(&loser)

	return true, nil
}
