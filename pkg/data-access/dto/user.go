package dto

import (
	"github.com/uptrace/bun"
)

type Role string

const (
	Role_Student Role = "student"
	Role_Teacher Role = "teacher"
	Role_Admin   Role = "admin"
)

type User struct {
	bun.BaseModel `bun:"table:users"`
	ID            int     `bun:"id,pk,autoincrement"`
	Username      string  `bun:"username,notnull"`
	Password      string  `bun:"password"`
	Name          string  `bun:"name"`
	Role          Role    `bun:"role,notnull"`
	Teaches       []Class `bun:"m2m:class_to_teacher,join:Teacher=Class"`
	Studies       []Class `bun:"m2m:class_to_student,join:Student=Class"`
}
