package interfaces

import (
	"context"
)

type ClassAPI interface {
	List(context.Context) error
	Get(context.Context) error
	Create(context.Context) error
	Update(context.Context) error
	Delete(context.Context) error
	ListStudents(context.Context) error
	ListTeachers(context.Context) error
	AddStudent(context.Context) error
	AddTeacher(context.Context) error
	RemoveStudent(context.Context) error
	RemoveTeacher(context.Context) error
}
