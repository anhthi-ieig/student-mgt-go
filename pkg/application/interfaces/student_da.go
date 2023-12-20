package interfaces

import (
	"context"
	"student-service/pkg/data-access/dto"
)

type StudentDA interface {
	GetStudents(context.Context) ([]dto.Student, error)
}
