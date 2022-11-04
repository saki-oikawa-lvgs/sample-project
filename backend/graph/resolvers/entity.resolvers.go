package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/saki-oikawa-lvgs/sample-project/backend/graph/customTypes"
	"github.com/saki-oikawa-lvgs/sample-project/backend/graph/generated"
)

// FindPostByID is the resolver for the findPostByID field.
func (r *entityResolver) FindPostByID(ctx context.Context, id int) (*customTypes.Post, error) {
	return &customTypes.Post{
		ID: id,
	}, nil
}

// FindTodoByID is the resolver for the findTodoByID field.
func (r *entityResolver) FindTodoByID(ctx context.Context, id int) (*customTypes.Todo, error) {
	panic(fmt.Errorf("not implemented: FindTodoByID - findTodoByID"))
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
