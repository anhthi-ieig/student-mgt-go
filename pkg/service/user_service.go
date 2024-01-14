package service

import (
	"context"
	"student-service/pkg/application/interfaces"
	"student-service/pkg/application/model"
	"student-service/pkg/data-access/dto"
)

type UserService struct {
	db interfaces.UserDA
}

func (s *UserService) Update(ctx context.Context, id int, request model.User) (model.User, error) {
	user, err := s.db.Update(ctx, id, modelToDtoUser(request))
	if err != nil {
		return model.User{}, err
	}
	return dtoToModelUser(user), nil
}

func (s *UserService) Get(ctx context.Context, id int) (model.User, error) {
	user, err := s.db.Get(ctx, id)
	if err != nil {
		return model.User{}, err
	}
	return dtoToModelUser(user), nil
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

func dtoToModelUser(d dto.User) model.User {
	m := model.User{
		Username: d.Username,
		Password: d.Password,
		Role:     d.Role,
		Name:     d.Name,
	}
	return m
}

func modelToDtoUser(m model.User) dto.User {
	d := dto.User{
		Username: m.Username,
		Password: m.Password,
		Role:     m.Role,
		Name:     m.Name,
	}
	return d
}

func NewUserService(db interfaces.UserDA) *UserService {
	return &UserService{db}
}
