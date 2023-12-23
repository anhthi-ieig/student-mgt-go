package model

import (
	"time"
)

type Class struct {
	ID        int       `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	StartDate time.Time `json:"startDate,omitempty"`
	EndDate   time.Time `json:"endDate,omitempty"`
	Subjects  string    `json:"subjects,omitempty"`
	Students  []User    `json:"students,omitempty"`
	Teachers  []User    `json:"teachers,omitempty"`
}
