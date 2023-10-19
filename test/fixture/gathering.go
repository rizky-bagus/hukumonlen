package fixture

import (
	"time"

	"github.com/DATA-DOG/go-sqlmock"

	entity "hukum-onlen-go/entity"
)

var (
	gatheringColumns = []string{
		"id",
		"name",
		"location",
		"creator_id",
		"type_id",
		"scheduled_at",
		"created_at",
		"updated_at",
		"deleted_at",
	}
)

// NewGathering creates a gathering entity.
func NewGathering(name, location string) *entity.Gathering {
	now := time.Now()
	return &entity.Gathering{
		ID:          entity.NewUUID(),
		Name:        name,
		Location:    location,
		CreatorID:   entity.NewUUID(),
		TypeID:      entity.NewUUID(),
		ScheduledAt: now,
		Auditable: entity.Auditable{
			CreatedAt: now,
			UpdatedAt: now,
		},
	}
}

// NewGatheringRows returns sqlmock rows from gathering entity.
func NewGatheringRows(mock sqlmock.Sqlmock, ents ...*entity.Gathering) *sqlmock.Rows {
	rows := mock.NewRows(gatheringColumns)
	for _, ent := range ents {
		rows.AddRow(
			ent.ID,
			ent.Name,
			ent.Location,
			ent.CreatorID,
			ent.TypeID,
			ent.ScheduledAt,
			ent.CreatedAt,
			ent.UpdatedAt,
			ent.DeletedAt,
		)
	}
	return rows
}
