package entity

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

// Auditable is a base entity that contains basic auditable fields.
type Auditable struct {
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

var _ bun.BeforeAppendModelHook = (*Auditable)(nil)

// BeforeAppendModel adds timestamps before model is inserted to the database.
func (m *Auditable) BeforeAppendModel(_ context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		m.CreatedAt = time.Now()
		m.UpdatedAt = time.Now()
	case *bun.UpdateQuery:
		m.UpdatedAt = time.Now()
	}
	return nil
}
