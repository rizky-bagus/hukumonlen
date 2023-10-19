package fixture

import (
	"time"

	"github.com/DATA-DOG/go-sqlmock"

	entity "hukum-onlen-go/entity"
)

var (
	gatheringTypeColumns = []string{
		"id",
		"name",
		"created_at",
		"updated_at",
		"deleted_at",
	}
)

// NewGatheringType creates a gathering type entity.
func NewGatheringType(name string) *entity.GatheringType {
	now := time.Now()
	return &entity.GatheringType{
		ID:   entity.NewUUID(),
		Name: name,
		Auditable: entity.Auditable{
			CreatedAt: now,
			UpdatedAt: now,
		},
	}
}

// NewGatheringTypeRows returns sqlmock rows from gathering type entity.
func NewGatheringTypeRows(mock sqlmock.Sqlmock, ents ...*entity.GatheringType) *sqlmock.Rows {
	rows := mock.NewRows(gatheringTypeColumns)
	for _, ent := range ents {
		rows.AddRow(
			ent.ID,
			ent.Name,
			ent.CreatedAt,
			ent.UpdatedAt,
			ent.DeletedAt,
		)
	}
	return rows
}
