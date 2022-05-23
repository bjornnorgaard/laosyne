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

func (r *mutationResolver) DeletePath(ctx context.Context, input model.DeletePath) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) ScanPath(ctx context.Context) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AddToRating(ctx context.Context, pictureID int) (*model.Picture, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateMatch(ctx context.Context, input *model.SearchFilter) (*model.Match, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) ReportMatchResult(ctx context.Context, input model.MatchResult) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetPaths(ctx context.Context) ([]*model.Path, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetPicture(ctx context.Context, input *model.SearchFilter) (*model.Picture, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetPictures(ctx context.Context, input *model.SearchFilter) ([]*model.Picture, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
