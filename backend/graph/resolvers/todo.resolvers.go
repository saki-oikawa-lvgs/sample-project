package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/saki-oikawa-lvgs/sample-project/backend/graph/common"
	"github.com/saki-oikawa-lvgs/sample-project/backend/graph/customTypes"
	"github.com/saki-oikawa-lvgs/sample-project/backend/graph/generated"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, text string) (*customTypes.Todo, error) {
	panic(fmt.Errorf("not implemented: CreateTodo - createTodo"))
}

// GetTodos is the resolver for the getTodos field.
func (r *queryResolver) GetTodos(ctx context.Context) ([]*customTypes.Todo, error) {
	context := common.GetContext(ctx)
	var todos []*customTypes.Todo
	err := context.Database.Find(&todos).Error
	if err != nil {
		return nil, err
	}
	//ここに追加する
	return todos, nil
}

// Post is the resolver for the post field.
func (r *todoResolver) Post(ctx context.Context, obj *customTypes.Todo) (*customTypes.Post, error) {
	return &customTypes.Post{
		ID: obj.PostID,
	}, nil
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
