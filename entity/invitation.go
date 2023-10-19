package entity

import "github.com/uptrace/bun"

// Invitation is an entity that represents an invitation.
type Invitation struct {
	bun.BaseModel `bun:"table:invitations"`

	ID          UUID `json:"id" bun:",pk"`
	MemberID    UUID `json:"member_id"`
	GatheringID UUID `json:"gathering_id"`
	StatusID    UUID `json:"status_id"`
	Auditable

	Member    *Member           `json:"member" bun:"rel:belongs-to,join:member_id=id"`
	Gathering *Gathering        `json:"gathering" bun:"rel:belongs-to,join:gathering_id=id"`
	Status    *InvitationStatus `json:"status" bun:"rel:belongs-to,join:status_id=id"`
}
