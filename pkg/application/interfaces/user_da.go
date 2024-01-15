package interfaces

import (
	"context"
	"student-service/pkg/data-access/dto"
)

type UserDA interface {
	List(context.Context) ([]dto.User, error)
	GetUserByUserName(context.Context, string) (dto.User, error)
	Get(context.Context, int) (dto.User, error)
	Create(context.Context, dto.User) (dto.User, error)
	Update(context.Context, int, dto.User) (dto.User, error)
}
