package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/saki-oikawa-lvgs/gqlgen_tutorial/server/graph/customTypes"
	"github.com/saki-oikawa-lvgs/gqlgen_tutorial/server/graph/generated"
)

// FindTodoByID is the resolver for the findTodoByID field.
func (r *entityResolver) FindTodoByID(ctx context.Context, id string) (*customTypes.Todo, error) {
	panic(fmt.Errorf("not implemented: FindTodoByID - findTodoByID"))
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *entityResolver) FindPostByID(ctx context.Context, id string) (*customTypes.Post, error) {
	panic(fmt.Errorf("not implemented: FindPostByID - findPostByID"))
}
