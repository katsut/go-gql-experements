package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.48

import (
	"bff/graph/model"
	"context"
	"fmt"
	"log/slog"
	"strconv"
	"userservice/pb"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, newUser model.NewUser) (*model.User, error) {
	panic(fmt.Errorf("not implemented: CreateUser - createUser"))
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	slog.Info("User")
	userId, err := strconv.Atoi(id) // IDをintに変換
	if err != nil {
		return nil, err
	}
	params := &pb.GetUserParams{Id: int32(userId)}

	user, err := r.client.GetUser(ctx, params)
	if err != nil {
		return nil, err
	}

	return userToModel(user), nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
func userToModel(pbUser *pb.User) *model.User {
	return &model.User{
		ID:   strconv.Itoa(int(pbUser.Id)),
		Name: pbUser.Name,
	}
}
