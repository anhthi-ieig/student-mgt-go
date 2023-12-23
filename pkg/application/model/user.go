package model

import "student-service/pkg/data-access/dto"

type User struct {
	ID       int      `json:"id"`
	Username string   `json:"username"`
	Name     string   `json:"name"`
	Role     dto.Role `json:"role"`
}
