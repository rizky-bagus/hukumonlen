package entity

import (
	"github.com/uptrace/bun"
)

// InvitationStatus is an entity that represents an invitation status.
type InvitationStatus struct {
	bun.BaseModel `bun:"table:invitation_statuses"`

	ID   UUID   `json:"id" bun:",pk"`
	Name string `json:"name"`
	Auditable
}
