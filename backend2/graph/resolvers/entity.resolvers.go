package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/saki-oikawa-lvgs/sample-project/backend2/graph/customTypes"
	"github.com/saki-oikawa-lvgs/sample-project/backend2/graph/generated"
)

// FindPostByID is the resolver for the findPostByID field.
func (r *entityResolver) FindPostByID(ctx context.Context, id int) (*customTypes.Post, error) {
	panic(fmt.Errorf("not implemented: FindPostByID - findPostByID"))
}

// FindTodoByID is the resolver for the findTodoByID field.
func (r *entityResolver) FindTodoByID(ctx context.Context, id int) (*customTypes.Todo, error) {
	return &customTypes.Todo{
		ID: id,
	}, nil
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
