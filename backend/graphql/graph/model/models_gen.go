// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type DeletePath struct {
	PathID int `json:"pathId"`
}

type Match struct {
	PlayerOne *Picture `json:"playerOne"`
	PlayerTwo *Picture `json:"playerTwo"`
}

type MatchResult struct {
	WinnerID int `json:"winnerId"`
	LoserID  int `json:"loserId"`
}

type NewPath struct {
	Path string `json:"path"`
}

type Path struct {
	ID        int    `json:"id"`
	Path      string `json:"path"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type Picture struct {
	ID        int     `json:"id"`
	Path      string  `json:"path"`
	Ext       string  `json:"ext"`
	Views     int     `json:"views"`
	Likes     int     `json:"likes"`
	Rating    float64 `json:"rating"`
	Deviation float64 `json:"deviation"`
	Wins      int     `json:"wins"`
	Losses    int     `json:"losses"`
	CreatedAt string  `json:"createdAt"`
	UpdatedAt string  `json:"updatedAt"`
}

type SearchFilter struct {
	Take         *int       `json:"take"`
	Skip         *int       `json:"skip"`
	PathContains *string    `json:"pathContains"`
	UpperRating  *int       `json:"upperRating"`
	LowerRating  *int       `json:"lowerRating"`
	SortOrder    *SortOrder `json:"sortOrder"`
}

type SortOrder string

const (
	SortOrderID            SortOrder = "ID"
	SortOrderRandom        SortOrder = "RANDOM"
	SortOrderRatingDesc    SortOrder = "RATING_DESC"
	SortOrderRatingAsc     SortOrder = "RATING_ASC"
	SortOrderViewsDesc     SortOrder = "VIEWS_DESC"
	SortOrderViewsAsc      SortOrder = "VIEWS_ASC"
	SortOrderLikesDesc     SortOrder = "LIKES_DESC"
	SortOrderLikesAsc      SortOrder = "LIKES_ASC"
	SortOrderWinsDesc      SortOrder = "WINS_DESC"
	SortOrderWinsAsc       SortOrder = "WINS_ASC"
	SortOrderLossesDesc    SortOrder = "LOSSES_DESC"
	SortOrderLossesAsc     SortOrder = "LOSSES_ASC"
	SortOrderCreatedAtDesc SortOrder = "CREATED_AT_DESC"
	SortOrderCreatedAtAsc  SortOrder = "CREATED_AT_ASC"
	SortOrderUpdatedAtDesc SortOrder = "UPDATED_AT_DESC"
	SortOrderUpdatedAtAsc  SortOrder = "UPDATED_AT_ASC"
)

var AllSortOrder = []SortOrder{
	SortOrderID,
	SortOrderRandom,
	SortOrderRatingDesc,
	SortOrderRatingAsc,
	SortOrderViewsDesc,
	SortOrderViewsAsc,
	SortOrderLikesDesc,
	SortOrderLikesAsc,
	SortOrderWinsDesc,
	SortOrderWinsAsc,
	SortOrderLossesDesc,
	SortOrderLossesAsc,
	SortOrderCreatedAtDesc,
	SortOrderCreatedAtAsc,
	SortOrderUpdatedAtDesc,
	SortOrderUpdatedAtAsc,
}

func (e SortOrder) IsValid() bool {
	switch e {
	case SortOrderID, SortOrderRandom, SortOrderRatingDesc, SortOrderRatingAsc, SortOrderViewsDesc, SortOrderViewsAsc, SortOrderLikesDesc, SortOrderLikesAsc, SortOrderWinsDesc, SortOrderWinsAsc, SortOrderLossesDesc, SortOrderLossesAsc, SortOrderCreatedAtDesc, SortOrderCreatedAtAsc, SortOrderUpdatedAtDesc, SortOrderUpdatedAtAsc:
		return true
	}
	return false
}

func (e SortOrder) String() string {
	return string(e)
}

func (e *SortOrder) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SortOrder(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SortOrder", str)
	}
	return nil
}

func (e SortOrder) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
