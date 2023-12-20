package dto

import (
	"time"

	"github.com/uptrace/bun"
)

type Student struct {
	bun.BaseModel `bun:"table:student"`
	ID            string    `bun:"id,pk"`
	Name          string    `bun:"name"`
	CreatedAt     time.Time `bun:"created_at,default:current_timestamp"`
}
