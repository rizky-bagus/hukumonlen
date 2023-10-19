package entity

import (
	"github.com/uptrace/bun"
)

// Member is an entity that represents a member.
type Member struct {
	bun.BaseModel `bun:"table:members"`

	ID        UUID   `json:"id" bun:",pk"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Auditable
}
