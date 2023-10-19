package entity

import (
	"time"

	"github.com/uptrace/bun"
)

// Gathering is an entity that represents a gathering.
type Gathering struct {
	bun.BaseModel `bun:"table:gatherings"`

	ID          UUID      `json:"id" bun:",pk"`
	Name        string    `json:"name"`
	Location    string    `json:"location"`
	CreatorID   UUID      `json:"creator_id"`
	TypeID      UUID      `json:"type_id"`
	ScheduledAt time.Time `json:"scheduled_at"`
	Auditable

	Creator *Member        `json:"creator" bun:"rel:belongs-to,join:creator_id=id"`
	Type    *GatheringType `json:"type" bun:"rel:belongs-to,join:type_id=id"`
}
