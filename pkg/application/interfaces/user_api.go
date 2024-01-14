package interfaces

import (
	"context"
	"student-service/pkg/application/model"
)

type UserAPI interface {
	GetUserByUserName(context.Context, string) (model.User, error)
	Update(context.Context) error
	Get(context.Context) error
}
