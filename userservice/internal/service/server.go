package service

import (
	"context"
	"userservice/internal/dao/query"
	"userservice/pb"
)

type UserServiceServer struct {
	q *query.Query
}

// Constructor
func NewUserServiceServer(q *query.Query) *UserServiceServer {
	return &UserServiceServer{q}
}

func (s *UserServiceServer) GetUser(c context.Context, p *pb.GetUserParams) (*pb.User, error) {
	u := s.q.User
	user, err := u.WithContext(c).Where(u.ID.Eq(p.Id)).First()
	if err != nil {
		return nil, err
	}
	return &pb.User{
		Id:   user.ID,
		Name: user.Name,
		Age:  user.Age,
	}, nil

}
