package interfaces

import (
	"context"
	"student-service/pkg/application/model"
)

type StudentAPI interface {
	List(context.Context) ([]model.Student, error)
}
