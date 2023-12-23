package service

import (
	"context"
	"student-service/pkg/application/interfaces"
	"student-service/pkg/application/model"
)

type UserService struct {
	db interfaces.UserDA
}

func (s *UserService) GetUserByUserName(c context.Context, username string) (model.User, error) {
	user, err := s.db.GetUserByUserName(c, username)

	if err != nil {
		return model.User{}, err
	}

	return model.User{
		Username: user.Username,
		Password: user.Password,
		Role:     user.Role,
		Name:     user.Name,
	}, nil
}

func NewUserService(db interfaces.UserDA) *UserService {
	return &UserService{db}
}
