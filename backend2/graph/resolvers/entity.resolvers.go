package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/saki-oikawa-lvgs/sample-project/backend2/graph/customTypes"
	"github.com/saki-oikawa-lvgs/sample-project/backend2/graph/generated"
	"github.com/saki-oikawa-lvgs/sample-project/backend2/service"
)

// FindPostByID is the resolver for the findPostByID field.
func (r *entityResolver) FindPostByID(ctx context.Context, id int) (*customTypes.Post, error) {
	return service.PostFindByID(id)
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
