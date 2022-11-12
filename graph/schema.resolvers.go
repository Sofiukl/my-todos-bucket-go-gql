package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/sofiukl/my-todos-bucket-go-gql/graph/generated"
	"github.com/sofiukl/my-todos-bucket-go-gql/graph/model"
)

// CreateBucket is the resolver for the createBucket field.
func (r *mutationResolver) CreateBucket(ctx context.Context, name string, category *string) (*model.Bucket, error) {
	panic(fmt.Errorf("not implemented: CreateBucket - createBucket"))
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (string, error) {
	panic(fmt.Errorf("not implemented: CreateUser - createUser"))
}

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	panic(fmt.Errorf("not implemented: Login - login"))
}

// RefreshToken is the resolver for the refreshToken field.
func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	panic(fmt.Errorf("not implemented: RefreshToken - refreshToken"))
}

// RecentBuckets is the resolver for the recentBuckets field.
func (r *queryResolver) RecentBuckets(ctx context.Context, count *int, offset *int) ([]*model.Bucket, error) {
	category := "up skills"
	var buckets []*model.Bucket
	dummyBucket := model.Bucket{
		ID:       "b1",
		Name:     "golang",
		Category: &category,
	}
	buckets = append(buckets, &dummyBucket)
	return buckets, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
