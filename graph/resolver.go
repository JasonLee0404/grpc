package graph

import (
	"context"
	"fmt"
	"grpc/mod/graph/model"
)

type Resolver struct {
	Todos []*model.Todo
	Users map[string]*model.User
}

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	user, exist := r.Resolver.Users[input.UserID]
	if !exists {
		return nil, fmt.Errorf("user with ID %s not found", input.UserID)
	}

	// Create a new Todo
	todo := &model.Todo{
		ID:   fmt.Sprintf("%d", len(r.Resolver.Todos)+1),
		Text: input.Text,
		Done: false,
		User: user,
	}

	// Store the Todo in memory
	r.Resolver.Todos = append(r.Resolver.Todos, todo)
	return todo, nil
}

// Todos is the resolver for the todos query.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return r.Resolver.Todos, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
