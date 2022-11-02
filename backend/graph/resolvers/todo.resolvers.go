package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/saki-oikawa-lvgs/gqlgen_tutorial/server/graph/common"
	"github.com/saki-oikawa-lvgs/gqlgen_tutorial/server/graph/customTypes"
	"github.com/saki-oikawa-lvgs/gqlgen_tutorial/server/graph/generated"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, text string) (*customTypes.Todo, error) {
	context := common.GetContext(ctx)
	todo := &customTypes.Todo{
		ID:   fmt.Sprintf("T%d", rand.Int()),
		Text: text,
		Done: false,
	}
	err := context.Database.Create(&todo).Error
	if err != nil {
		return nil, err
	}
	return todo, nil
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
		ID: obj.ID,
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

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *todoResolver) Text(ctx context.Context, obj *customTypes.Todo) (*customTypes.Post, error) {
	panic(fmt.Errorf("not implemented: Content - content"))
}
func (r *queryResolver) GetTodo(ctx context.Context, todoID string) (*customTypes.Todo, error) {
	// context := common.GetContext(ctx)
	var todo *customTypes.Todo
	// err := context.Database.Where("id = ?", todoID).Find(&todo).Error
	// if err != nil {
	// 	return nil, err
	// }
	return todo, nil
}
