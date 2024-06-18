package service

import (
	"context"
	"userservice/pb"
)

type UserServiceServer struct{}

// Constructor
func NewUserServiceServer() *UserServiceServer {
	return &UserServiceServer{}
}

func (s *UserServiceServer) GetUser(_ context.Context, _ *pb.GetUserParams) (*pb.User, error) {
	user := pb.User{
		Id:   "uuid",
		Name: "長尾景虎",
		Age:  25,
	}
	return &user, nil
}
