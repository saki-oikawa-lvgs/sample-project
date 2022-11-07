package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/saki-oikawa-lvgs/sample-project/backend2/graph/customTypes"
	"github.com/saki-oikawa-lvgs/sample-project/backend2/graph/generated"
	"github.com/saki-oikawa-lvgs/sample-project/backend2/service"
)

// CreatePost is the resolver for the create_post field.
func (r *mutationResolver) CreatePost(ctx context.Context, input customTypes.NewPost) (*customTypes.Post, error) {
	return service.PostCreate(input)
}

// GetPosts is the resolver for the getPosts field.
func (r *queryResolver) GetPosts(ctx context.Context) ([]*customTypes.Post, error) {
	return service.PostsGet()
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
