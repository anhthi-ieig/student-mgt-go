package interfaces

import (
	"context"
	"student-service/pkg/application/model"
)

type UserService interface {
	GetUserByUserName(context.Context, string) (model.User, error)
	Get(context.Context, int) (model.User, error)
	List(context.Context) ([]model.User, error)
	Create(context.Context, model.User) (model.User, error)
	Update(context.Context, int, model.User) (model.User, error)
}
