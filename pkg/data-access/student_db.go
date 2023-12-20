package dataaccess

import (
	"context"
	"student-service/pkg/data-access/dto"
	"student-service/pkg/utils"
	"time"

	"github.com/uptrace/bun"
)

type studentDA struct {
	dbc *bun.DB
}

func (s *studentDA) GetStudents(c context.Context) ([]dto.Student, error) {
	// See more: https://bun.uptrace.dev/guide/
	// var list []dto.Student
	// s.dbc.NewSelect().Model(&list).Scan(c)

	return []dto.Student{
		{
			ID:        "123",
			Name:      utils.SampleUtil(),
			CreatedAt: time.Now(),
		},
	}, nil
}

// NewStudentDA creates a new Student Data Access
func NewStudentDA(dbc *bun.DB) *studentDA {
	return &studentDA{dbc}
}
