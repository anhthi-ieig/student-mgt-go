package dataaccess

import (
	"context"
	"student-service/pkg/data-access/dto"

	"github.com/uptrace/bun"
)

type userDA struct {
	dbc *bun.DB
}

func (s *userDA) GetUserByUserName(c context.Context, username string) (dto.User, error) {

	var user dto.User
	err := s.dbc.NewSelect().Model(&user).Where("username = ?", username).Limit(1).Scan(c)

	if err != nil {
		panic(err)
	}

	return dto.User{
		ID:       user.ID,
		Username: user.Username,
		Password: user.Password,
		Role:     user.Role,
		Name:     user.Name,
	}, err
}

// NewStudentDA creates a new Student Data Access
func NewUserDA(dbc *bun.DB) *userDA {
	return &userDA{dbc}
}
