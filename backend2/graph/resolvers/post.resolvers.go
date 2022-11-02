package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/saki-oikawa-lvgs/gqlgen_tutorial/server/graph/customTypes"
	"github.com/saki-oikawa-lvgs/gqlgen_tutorial/server/graph/generated"
)

// CreatePost is the resolver for the createPost field.
func (r *mutationResolver) CreatePost(ctx context.Context, text string) (*customTypes.Post, error) {
	panic(fmt.Errorf("not implemented: CreatePost - createPost"))
}

// Todo is the resolver for the todo field.
func (r *postResolver) Todo(ctx context.Context, obj *customTypes.Post) (*customTypes.Todo, error) {
	panic(fmt.Errorf("not implemented: Todo - todo"))
}

// GetPosts is the resolver for the getPosts field.
func (r *queryResolver) GetPosts(ctx context.Context) ([]*customTypes.Post, error) {
	panic(fmt.Errorf("not implemented: GetPosts - getPosts"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Post returns generated.PostResolver implementation.
func (r *Resolver) Post() generated.PostResolver { return &postResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type postResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
