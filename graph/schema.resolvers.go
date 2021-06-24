package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/blr-coder/graphql_learning/graph/generated"
	"github.com/blr-coder/graphql_learning/graph/model"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	testUser := &model.User{
		ID:   fmt.Sprintf("%d", 22),
		Name: input.Name,
		Pass: input.Pass,
	}
	return testUser, nil
}

func (r *queryResolver) GetUsers(ctx context.Context) ([]*model.User, error) {
	var users []*model.User
	testUser := model.User{
		ID:   fmt.Sprintf("T%d", 17),
		Name: "Vasia2",
		Pass: "qwerty",
	}
	users = append(users, &testUser)
	return users, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
