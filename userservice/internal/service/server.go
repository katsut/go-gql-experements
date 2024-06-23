package service

import (
	"context"
	"userservice/internal/dao/query"
	"userservice/pb"
)

type UserServiceServer struct {
}

// Constructor
func NewUserServiceServer() *UserServiceServer {
	return &UserServiceServer{}
}

func (s *UserServiceServer) GetUser(c context.Context, p *pb.GetUserParams) (*pb.User, error) {
	u := query.User
	user, err := query.User.WithContext(c).Where(u.ID.Eq(p.Id)).First()
	if err != nil {
		return nil, err
	}
	return &pb.User{
		Id:   user.ID,
		Name: user.Name,
		Age:  user.Age,
	}, nil

}
