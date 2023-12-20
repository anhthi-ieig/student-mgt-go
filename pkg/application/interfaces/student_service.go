package interfaces

import (
	"context"
	"student-service/pkg/application/model"
)

type StudentService interface {
	GetStudents(context.Context) ([]model.Student, error)
}
