package fixture

import (
	"time"

	"github.com/DATA-DOG/go-sqlmock"

	entity "hukum-onlen-go/entity"
)

var (
	memberColumns = []string{
		"id",
		"email",
		"first_name",
		"last_name",
		"created_at",
		"updated_at",
		"deleted_at",
	}
)

// NewMember creates a member entity.
func NewMember(email, firstName, lastName string) *entity.Member {
	now := time.Now()
	return &entity.Member{
		ID:        entity.NewUUID(),
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
		Auditable: entity.Auditable{
			CreatedAt: now,
			UpdatedAt: now,
		},
	}
}

// NewMemberRows returns sqlmock rows from member entity.
func NewMemberRows(mock sqlmock.Sqlmock, ents ...*entity.Member) *sqlmock.Rows {
	rows := mock.NewRows(memberColumns)
	for _, ent := range ents {
		rows.AddRow(
			ent.ID,
			ent.Email,
			ent.FirstName,
			ent.LastName,
			ent.CreatedAt,
			ent.UpdatedAt,
			ent.DeletedAt,
		)
	}
	return rows
}
