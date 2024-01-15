package interfaces

import (
	"context"
	"student-service/pkg/application/model"
)

type UserAPI interface {
	GetUserByUserName(context.Context, string) (model.User, error)
	Get(context.Context) error
	List(context.Context) ([]model.User, error)
	Create(context.Context) (model.User, error)
	Update(context.Context) error
}
