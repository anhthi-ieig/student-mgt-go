package service

import (
	"context"
	"student-service/pkg/application/interfaces"
	"student-service/pkg/application/model"
	"student-service/pkg/data-access/dto"
)

type classService struct {
	db interfaces.ClassDA
}

func (s *classService) List(c context.Context) ([]model.Class, error) {
	list, err := s.db.List(c)
	if err != nil {
		return nil, err
	}

	result := make([]model.Class, len(list))
	for i, v := range list {
		result[i] = dtoToModel(v)
	}
	return result, nil
}

func (s *classService) Get(c context.Context, id int) (model.Class, error) {
	class, err := s.db.Get(c, id)
	if err != nil {
		return model.Class{}, err
	}
	return dtoToModel(class), nil
}

func (s *classService) Create(c context.Context, request model.Class) (model.Class, error) {
	class, err := s.db.Create(c, modelToDto(request))
	if err != nil {
		return model.Class{}, err
	}
	return dtoToModel(class), nil
}

func (s *classService) Update(c context.Context, id int, request model.Class) (model.Class, error) {
	class, err := s.db.Update(c, id, modelToDto(request))
	if err != nil {
		return model.Class{}, err
	}
	return dtoToModel(class), nil
}

func (s *classService) Delete(c context.Context, id int) error {
	return s.db.Delete(c, id)
}

func (s *classService) ListStudents(c context.Context, id int) ([]model.User, error) {
	students, err := s.db.ListStudents(c, id)
	if err != nil {
		return nil, err
	}

	result := make([]model.User, len(students))
	for i, s := range students {
		result[i] = model.User{
			ID:       s.ID,
			Username: s.Username,
			Name:     s.Name,
			Role:     s.Role,
		}
	}

	return result, nil
}
func (s *classService) ListTeachers(c context.Context, id int) ([]model.User, error) {
	teachers, err := s.db.ListTeachers(c, id)
	if err != nil {
		return nil, err
	}

	result := make([]model.User, len(teachers))
	for i, s := range teachers {
		result[i] = model.User{
			ID:       s.ID,
			Username: s.Username,
			Name:     s.Name,
			Role:     s.Role,
		}
	}

	return result, nil
}

func (s *classService) AddStudent(c context.Context, classId int, studentId int) error {
	d := []dto.ClassToStudent{
		{
			StudentID: studentId,
			ClassID:   classId,
		},
	}
	return s.db.AddStudent(c, d)
}

func (s *classService) AddTeacher(c context.Context, classId int, teacherId int) error {
	d := []dto.ClassToTeacher{
		{
			TeacherID: teacherId,
			ClassID:   classId,
		},
	}
	return s.db.AddTeacher(c, d)
}
func (s *classService) RemoveStudent(c context.Context, classId int, studentId int) error {
	d := []dto.ClassToStudent{
		{
			StudentID: studentId,
			ClassID:   classId,
		},
	}
	return s.db.RemoveStudent(c, d)
}

func (s *classService) RemoveTeacher(c context.Context, classId int, teacherId int) error {
	d := []dto.ClassToTeacher{
		{
			TeacherID: teacherId,
			ClassID:   classId,
		},
	}
	return s.db.RemoveTeacher(c, d)
}

func dtoToModel(d dto.Class) model.Class {
	m := model.Class{
		ID:        d.ID,
		Name:      d.Name,
		StartDate: d.StartDate,
		EndDate:   d.EndDate,
		Subjects:  d.Subjects,
		Students:  make([]model.User, len(d.Students)),
		Teachers:  make([]model.User, len(d.Teachers)),
	}

	for i, s := range d.Students {
		m.Students[i] = model.User{
			ID:       s.ID,
			Username: s.Username,
			Name:     s.Name,
			Role:     s.Role,
		}
	}

	for i, t := range d.Teachers {
		m.Teachers[i] = model.User{
			ID:       t.ID,
			Username: t.Username,
			Name:     t.Name,
			Role:     t.Role,
		}
	}
	return m
}

func modelToDto(m model.Class) dto.Class {
	d := dto.Class{
		ID:        m.ID,
		Name:      m.Name,
		StartDate: m.StartDate,
		EndDate:   m.EndDate,
		Subjects:  m.Subjects,
		Students:  make([]dto.User, len(m.Students)),
		Teachers:  make([]dto.User, len(m.Teachers)),
	}

	for i, s := range m.Students {
		d.Students[i] = dto.User{
			ID:       s.ID,
			Username: s.Username,
			Name:     s.Name,
			Role:     s.Role,
		}
	}

	for i, t := range m.Teachers {
		d.Teachers[i] = dto.User{
			ID:       t.ID,
			Username: t.Username,
			Name:     t.Name,
			Role:     t.Role,
		}
	}
	return d
}

func NewClassService(db interfaces.ClassDA) *classService {
	return &classService{db}
}
