package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/saki-oikawa-lvgs/sample-project/backend2/graph/common"
	"github.com/saki-oikawa-lvgs/sample-project/backend2/graph/customTypes"
	"github.com/saki-oikawa-lvgs/sample-project/backend2/graph/generated"
)

// Todo is the resolver for the todo field.
func (r *postResolver) Todo(ctx context.Context, obj *customTypes.Post) (*customTypes.Todo, error) {
	return &customTypes.Todo{
		ID: 888,
	}, nil
}

// GetPosts is the resolver for the getPosts field.
func (r *queryResolver) GetPosts(ctx context.Context) ([]*customTypes.Post, error) {
	context := common.GetContext(ctx)
	var posts []*customTypes.Post
	err := context.Database.Find(&posts).Error
	if err != nil {
		return nil, err
	}
	return posts, nil
}

// Post returns generated.PostResolver implementation.
func (r *Resolver) Post() generated.PostResolver { return &postResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type postResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
