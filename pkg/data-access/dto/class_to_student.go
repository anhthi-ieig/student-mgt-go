package dto

import "github.com/uptrace/bun"

type ClassToStudent struct {
	bun.BaseModel `bun:"table:class_to_student"`
	ClassID       int    `bun:",pk"`
	Class         *Class `bun:"rel:belongs-to,join:class_id=id"`
	StudentID     int    `bun:",pk"`
	Student       *User  `bun:"rel:belongs-to,join:student_id=id"`
}
