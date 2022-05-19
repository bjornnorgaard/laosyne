package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/bjornnorgaard/laosyne/api/graphql/graph/generated"
	"github.com/bjornnorgaard/laosyne/api/graphql/graph/model"
)

func (r *mutationResolver) AddMediaPath(ctx context.Context, input model.NewMediaPath) (*model.MediaPath, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) MediaPaths(ctx context.Context) ([]*model.MediaPath, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
