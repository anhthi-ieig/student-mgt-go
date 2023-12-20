package dto

import "github.com/uptrace/bun"

type ClassToTeacher struct {
	bun.BaseModel `bun:"table:class_to_teacher"`
	ClassID       int    `bun:",pk"`
	Class         *Class `bun:"rel:belongs-to,join:class_id=id"`
	TeacherID     int    `bun:",pk"`
	Teacher       *User  `bun:"rel:belongs-to,join:teacher_id=id"`
}
