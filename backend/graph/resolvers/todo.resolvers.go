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
	return todos, nil
}

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
// func (r *mutationResolver) UpdateTodo(ctx context.Context, input customTypes.TodoInput) (*customTypes.Todo, error) {
// 	context := common.GetContext(ctx)
// 	todo := &customTypes.Todo{
// 		ID:   input.ID,
// 		Text: input.Text,
// 		Done: input.Done,
// 	}
// 	err := context.Database.Save(&todo).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return todo, nil
// }
// func (r *mutationResolver) DeleteTodo(ctx context.Context, todoID string) (*customTypes.Todo, error) {
// 	context := common.GetContext(ctx)
// 	var todo *customTypes.Todo
// 	err := context.Database.Where("id = ?", todoID).Delete(&todo).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return todo, nil
// }
func (r *queryResolver) GetTodo(ctx context.Context, todoID string) (*customTypes.Todo, error) {
	// context := common.GetContext(ctx)
	var todo *customTypes.Todo
	// err := context.Database.Where("id = ?", todoID).Find(&todo).Error
	// if err != nil {
	// 	return nil, err
	// }
	return todo, nil
}
