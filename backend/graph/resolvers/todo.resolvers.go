package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/saki-oikawa-lvgs/gqlgen_tutorial/server/graph/customTypes"
	"github.com/saki-oikawa-lvgs/gqlgen_tutorial/server/graph/generated"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, text string) (*customTypes.Todo, error) {
	panic(fmt.Errorf("not implemented: CreateTodo - createTodo"))
}

// GetTodos is the resolver for the getTodos field.
func (r *queryResolver) GetTodos(ctx context.Context) ([]*customTypes.Todo, error) {
	panic(fmt.Errorf("not implemented: GetTodos - getTodos"))
}

// Post is the resolver for the post field.
func (r *todoResolver) Post(ctx context.Context, obj *customTypes.Todo) (*customTypes.Post, error) {
	panic(fmt.Errorf("not implemented: Post - post"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
