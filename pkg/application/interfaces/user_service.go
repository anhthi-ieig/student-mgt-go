package interfaces

import (
	"context"
	"student-service/pkg/application/model"
)

type UserService interface {
	GetUserByUserName(context.Context, string) (model.User, error)
}
