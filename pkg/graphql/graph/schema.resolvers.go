package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.21 DO NOT EDIT.

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"

	"github.com/ncostamagna/gokit-graphql/pkg/graphql/graph/model"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	rand, _ := rand.Int(rand.Reader, big.NewInt(100))

	var user *model.User

	for _, u := range r.users {
		if input.UserID == u.ID {
			user = u
		}
	}

	if user == nil {
		return nil, errors.New("User doesn't exist")
	}
	todo := &model.Todo{
		Text: input.Text,
		ID:   fmt.Sprintf("T%d", rand),
		User: user,
	}
	r.todos = append(r.todos, todo)
	return todo, nil
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	rand, _ := rand.Int(rand.Reader, big.NewInt(100))
	user := &model.User{
		Name:  input.Name,
		ID:    fmt.Sprintf("T%d", rand),
		Email: input.Email,
	}
	r.users = append(r.users, user)
	return user, nil
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return r.todos, nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	return r.users, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
