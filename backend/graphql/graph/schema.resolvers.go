package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/bjornnorgaard/laosyne/backend/graphql/graph/generated"
	"github.com/bjornnorgaard/laosyne/backend/graphql/graph/model"
)

func (r *mutationResolver) AddPath(ctx context.Context, input model.NewPath) (*model.Path, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeletePath(ctx context.Context, pathID int) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) ScanPaths(ctx context.Context) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AddToRating(ctx context.Context, pictureID int) (*model.Picture, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) LikePicture(ctx context.Context, pictureID int) (*model.Picture, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DislikePicture(ctx context.Context, pictureID int) (*model.Picture, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) ReportMatchResult(ctx context.Context, input model.MatchResult) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Paths(ctx context.Context) ([]*model.Path, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Picture(ctx context.Context, pictureID int) (*model.Picture, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Pictures(ctx context.Context, input *model.SearchFilter) ([]*model.Picture, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Match(ctx context.Context, input *model.SearchFilter) (*model.Match, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *queryResolver) GetPaths(ctx context.Context) ([]*model.Path, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *queryResolver) GetPicture(ctx context.Context, pictureID int) (*model.Picture, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *queryResolver) GetPictures(ctx context.Context, input *model.SearchFilter) ([]*model.Picture, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *queryResolver) CreateMatch(ctx context.Context, input *model.SearchFilter) (*model.Match, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *queryResolver) GetLeaderboard(ctx context.Context, input *model.SearchFilter) ([]*model.Picture, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *mutationResolver) CreateMatch(ctx context.Context, input *model.SearchFilter) (*model.Match, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *mutationResolver) ScanPath(ctx context.Context) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}
