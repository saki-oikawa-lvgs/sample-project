package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/saki-oikawa-lvgs/sample-project/backend/graph/customTypes"
	"github.com/saki-oikawa-lvgs/sample-project/backend/graph/generated"
	"github.com/saki-oikawa-lvgs/sample-project/backend/service"
)

// CreateTodo is the resolver for the create_todo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input customTypes.NewTodo) (*customTypes.Todo, error) {
	return service.TodoCreate(input)
}

// Posts is the resolver for the posts field.
func (r *postResolver) Posts(ctx context.Context, obj *customTypes.Post) ([]*customTypes.Post, error) {
	panic(fmt.Errorf("not implemented: GetTodos - getTodos"))
}

// GetTodos is the resolver for the getTodos field.
func (r *queryResolver) GetTodos(ctx context.Context) ([]*customTypes.Todo, error) {
	return service.TodosGet()
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
