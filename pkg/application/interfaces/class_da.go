package interfaces

import (
	"context"
	"student-service/pkg/data-access/dto"
)

type ClassDA interface {
	List(context.Context) ([]dto.Class, error)
	Get(context.Context, int) (dto.Class, error)
	Create(context.Context, dto.Class) (dto.Class, error)
	Update(context.Context, int, dto.Class) (dto.Class, error)
	Delete(context.Context, int) error
	ListStudents(context.Context, int) ([]dto.User, error)
	ListTeachers(context.Context, int) ([]dto.User, error)
	AddStudent(context.Context, []dto.ClassToStudent) error
	AddTeacher(context.Context, []dto.ClassToTeacher) error
	RemoveStudent(context.Context, []dto.ClassToStudent) error
	RemoveTeacher(context.Context, []dto.ClassToTeacher) error
}
