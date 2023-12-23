package interfaces

import (
	"context"
	"student-service/pkg/application/model"
)

type ClassService interface {
	List(context.Context) ([]model.Class, error)
	Get(context.Context, int) (model.Class, error)
	Create(context.Context, model.Class) (model.Class, error)
	Update(context.Context, int, model.Class) (model.Class, error)
	Delete(context.Context, int) error
	ListStudents(context.Context, int) ([]model.User, error)
	ListTeachers(context.Context, int) ([]model.User, error)
	AddStudent(context.Context, int, int) error
	AddTeacher(context.Context, int, int) error
	RemoveStudent(context.Context, int, int) error
	RemoveTeacher(context.Context, int, int) error
}
