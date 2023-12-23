package model

import "student-service/pkg/data-access/dto"

type User struct {
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	Password string   `json:"password"`
	Username string   `json:"username"`
	Role     dto.Role `json:"role"`
}
