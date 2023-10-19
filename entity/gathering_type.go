package entity

import (
	"github.com/uptrace/bun"
)

// GatheringType is an entity that represents a gathering type.
type GatheringType struct {
	bun.BaseModel `bun:"table:gathering_types"`

	ID   UUID   `json:"id" bun:",pk"`
	Name string `json:"name"`
	Auditable
}
