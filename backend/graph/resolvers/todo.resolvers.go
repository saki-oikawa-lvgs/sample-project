package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	customTypes1 "github.com/saki-oikawa-lvgs/sample-project/backend/graph/customTypes"
	generated1 "github.com/saki-oikawa-lvgs/sample-project/backend/graph/generated"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, text string) (*customTypes1.Todo, error) {
	panic(fmt.Errorf("not implemented: CreateTodo - createTodo"))
}

// GetTodos is the resolver for the getTodos field.
func (r *queryResolver) GetTodos(ctx context.Context) ([]*customTypes1.Todo, error) {
	panic(fmt.Errorf("not implemented: GetTodos - getTodos"))
}

// Post is the resolver for the post field.
func (r *todoResolver) Post(ctx context.Context, obj *customTypes1.Todo) (*customTypes1.Post, error) {
	panic(fmt.Errorf("not implemented: Post - post"))
}

// Mutation returns generated1.MutationResolver implementation.
func (r *Resolver) Mutation() generated1.MutationResolver { return &mutationResolver{r} }

// Query returns generated1.QueryResolver implementation.
func (r *Resolver) Query() generated1.QueryResolver { return &queryResolver{r} }

// Todo returns generated1.TodoResolver implementation.
func (r *Resolver) Todo() generated1.TodoResolver { return &todoResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
