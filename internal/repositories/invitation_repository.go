package repositories

import (
	"context"

	"github.com/uptrace/bun"

	entity "hukum-onlen-go/entity"
)

// IInvitationRepository is an invitation  repository interface.
type IInvitationRepository interface {
	CreateInvitations(ctx context.Context, invitations []*entity.Invitation) error
}

// InvitationRepository is an invitation  repository.
type InvitationRepository struct {
	db *bun.DB
}

// NewInvitationRepository creates an instance of invitationRepository.
func NewInvitationRepository(db *bun.DB) *InvitationRepository {
	return &InvitationRepository{
		db: db,
	}
}

// CreateInvitations creates multiple invitations.
func (r *InvitationRepository) CreateInvitations(ctx context.Context, invitations []*entity.Invitation) error {
	_, err := r.db.NewInsert().Model(&invitations).Exec(ctx)
	return err
}
