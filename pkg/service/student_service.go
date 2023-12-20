package service

import (
	"context"
	"student-service/pkg/application/interfaces"
	"student-service/pkg/application/model"
)

type studentService struct {
	db interfaces.StudentDA
}

func (s *studentService) GetStudents(c context.Context) ([]model.Student, error) {
	list, err := s.db.GetStudents(c)
	if err != nil {
		return nil, err
	}

	// map from dto to model
	result := make([]model.Student, len(list))
	for i, v := range list {
		result[i] = model.Student{
			ID:        v.ID,
			Name:      v.Name,
			CreatedAt: v.CreatedAt,
		}
	}
	return result, nil
}

func NewStudentService(db interfaces.StudentDA) *studentService {
	return &studentService{db}
}
