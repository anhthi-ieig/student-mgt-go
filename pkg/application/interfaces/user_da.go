package interfaces

import (
	"context"
	"student-service/pkg/data-access/dto"
)

type UserDA interface {
	GetUserByUserName(context.Context, string) (dto.User, error)
	Update(context.Context, int, dto.User) (dto.User, error)
	Get(context.Context, int) (dto.User, error)
}
