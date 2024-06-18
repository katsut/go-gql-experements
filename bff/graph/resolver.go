package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
import (
	"userservice/pb"
)

type Resolver struct {
	client pb.UserServiceClient
}

func NewResolver(client pb.UserServiceClient) *Resolver {
	return &Resolver{client: client}
}
