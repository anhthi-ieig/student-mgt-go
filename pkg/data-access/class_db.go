package dataaccess

import (
	"context"
	"student-service/pkg/data-access/dto"

	"github.com/uptrace/bun"
)

type classDA struct {
	dbc *bun.DB
}

func (s *classDA) List(c context.Context) ([]dto.Class, error) {
	var classes []dto.Class

	if err := s.dbc.NewSelect().
		Model(&classes).
		Relation("Students").
		Relation("Teachers").
		Scan(c); err != nil {
		return nil, err
	}

	return classes, nil
}

func (s *classDA) Get(c context.Context, ID int) (dto.Class, error) {
	class := dto.Class{ID: ID}

	if err := s.dbc.NewSelect().
		Model(&class).
		Relation("Students").
		Relation("Teachers").
		WherePK().
		Scan(c); err != nil {
		return dto.Class{}, err
	}

	return class, nil
}

func (s *classDA) Create(c context.Context, class dto.Class) (dto.Class, error) {
	if _, err := s.dbc.NewInsert().
		Model(&class).
		On("CONFLICT (id) DO UPDATE").
		Returning("*").
		Exec(c); err != nil {
		return dto.Class{}, err
	}
	return class, nil
}

func (s *classDA) Update(c context.Context, id int, class dto.Class) (dto.Class, error) {
	if _, err := s.dbc.NewUpdate().
		Model(&class).
		Where("id = ?", id).
		Returning("*").
		Exec(c); err != nil {
		return dto.Class{}, err
	}
	return class, nil
}

func (s *classDA) Delete(c context.Context, id int) (err error) {
	class := dto.Class{ID: id}
	_, err = s.dbc.NewDelete().Model(&class).WherePK().Exec(c)
	return
}

func (s *classDA) ListStudents(c context.Context, id int) ([]dto.User, error) {
	var students []dto.User
	subq := s.dbc.NewSelect().Model((*dto.ClassToStudent)(nil)).Column("student_id").Where("class_id = ?", id)

	err := s.dbc.NewSelect().Model((*dto.User)(nil)).Where("id IN (?)", subq).Scan(c, &students)
	if err != nil {
		return nil, err
	}

	return students, nil
}

func (s *classDA) ListTeachers(c context.Context, id int) ([]dto.User, error) {
	var students []dto.User
	subq := s.dbc.NewSelect().Model((*dto.ClassToTeacher)(nil)).Column("teacher_id").Where("class_id = ?", id)

	err := s.dbc.NewSelect().Model((*dto.User)(nil)).Where("id IN (?)", subq).Scan(c, &students)
	if err != nil {
		return nil, err
	}

	return students, nil
}

func (s *classDA) AddStudent(c context.Context, d []dto.ClassToStudent) error {
	_, err := s.dbc.NewInsert().Model(&d).Exec(c)
	return err
}

func (s *classDA) AddTeacher(c context.Context, d []dto.ClassToTeacher) error {
	_, err := s.dbc.NewInsert().Model(&d).Exec(c)
	return err
}

func (s *classDA) RemoveStudent(c context.Context, d []dto.ClassToStudent) error {
	_, err := s.dbc.NewDelete().Model(&d).WherePK().Exec(c)
	return err
}

func (s *classDA) RemoveTeacher(c context.Context, d []dto.ClassToTeacher) error {
	_, err := s.dbc.NewDelete().Model(&d).WherePK().Exec(c)
	return err
}

func NewClassDA(dbc *bun.DB) *classDA {
	return &classDA{dbc}
}
