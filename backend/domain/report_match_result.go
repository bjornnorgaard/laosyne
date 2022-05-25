package domain

import (
	"context"
	"fmt"

	"github.com/bjornnorgaard/laosyne/backend/graphql/graph/model"
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

	skills := []trueskill.Player{playerWinner, playerLoser}
	newSkills, _ := ts.AdjustSkills(skills, false)

	playerWinner = newSkills[0]
	playerLoser = newSkills[1]

	winner.Rating = playerWinner.Mean()
	winner.Deviation = playerWinner.Sigma()
	winner.Wins++

	loser.Rating = playerLoser.Mean()
	loser.Deviation = playerLoser.Sigma()
	loser.Losses++

	a.db.Save(&winner)
	a.db.Save(&loser)

	return true, nil
}
