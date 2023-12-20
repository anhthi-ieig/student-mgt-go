package dto

import (
	"time"

	"github.com/uptrace/bun"
)

type Class struct {
	bun.BaseModel `bun:"table:class"`
	ID            int       `bun:"id,pk,autoincrement"`
	Name          string    `bun:"name"`
	StartDate     time.Time `bun:"start_date"`
	EndDate       time.Time `bun:"end_date"`
	Subjects      string    `bun:"subjects"`
	Students      []User    `bun:"m2m:class_to_student,join:Class=Student"`
	Teachers      []User    `bun:"m2m:class_to_teacher,join:Class=Teacher"`
}
