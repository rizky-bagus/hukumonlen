package entity

import (
	"database/sql/driver"

	"github.com/google/uuid"
)

// UUID is a special UUID type that serializes data to MySQL's BINARY(16) data type.
type UUID struct {
	uuid.UUID
}

// Value implements the driver.Valuer interface.
func (u UUID) Value() (driver.Value, error) {
	return u.MarshalBinary()
}

// NewUUID returns a random UUID.
func NewUUID() UUID {
	return UUID{uuid.New()}
}
