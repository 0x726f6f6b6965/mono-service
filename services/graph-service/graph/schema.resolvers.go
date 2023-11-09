package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.40

import (
	"context"
	"errors"

	"github.com/0x726f6f6b6965/mono-service/services/graph-service/internal/middleware"
	"github.com/0x726f6f6b6965/mono-service/services/graph-service/internal/model"
	"github.com/0x726f6f6b6965/mono-service/services/graph-service/internal/service"
)

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, input model.LoginInput) (string, error) {
	return r.authService.Login(ctx, input)
}

// Register is the resolver for the register field.
func (r *mutationResolver) Register(ctx context.Context, input model.NewUser) (string, error) {
	return r.authService.Register(ctx, input)
}

// EditRole is the resolver for the editRole field.
func (r *mutationResolver) EditRole(ctx context.Context, input *model.EditRole) (*model.UserRole, error) {
	user := middleware.ForContext(ctx)

	allow, err := service.Auth(ctx,
		service.Request{
			Role: user.UserRole,
			Path: "editRole",
		})

	if !allow || err != nil {
		return &model.UserRole{}, errors.New("access denied")
	}
	return r.authService.EditRole(ctx, input)
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
