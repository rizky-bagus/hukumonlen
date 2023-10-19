package fixture

import (
	"time"

	"github.com/DATA-DOG/go-sqlmock"

	entity "hukum-onlen-go/entity"
)

var (
	invitationStatusColumns = []string{
		"id",
		"name",
		"created_at",
		"updated_at",
		"deleted_at",
	}
)

// NewInvitationStatus creates an invitation status entity.
func NewInvitationStatus(name string) *entity.InvitationStatus {
	now := time.Now()
	return &entity.InvitationStatus{
		ID:   entity.NewUUID(),
		Name: name,
		Auditable: entity.Auditable{
			CreatedAt: now,
			UpdatedAt: now,
		},
	}
}

// NewInvitationStatusRows returns sqlmock rows from invitation status entity.
func NewInvitationStatusRows(mock sqlmock.Sqlmock, ents ...*entity.InvitationStatus) *sqlmock.Rows {
	rows := mock.NewRows(invitationStatusColumns)
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
