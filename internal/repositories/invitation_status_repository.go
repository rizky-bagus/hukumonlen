package repositories

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
	"github.com/uptrace/bun"

	entity "hukum-onlen-go/entity"
)

// IInvitationStatusRepository is an invitation status repository interface.
type IInvitationStatusRepository interface {
	FindAllInvitationStatuses(ctx context.Context) ([]*entity.InvitationStatus, error)
	FindInvitationStatusByID(ctx context.Context, id string) (*entity.InvitationStatus, error)
	FindPendingInvitationStatus(ctx context.Context) (*entity.InvitationStatus, error)
	CreateInvitationStatus(ctx context.Context, invitationStatus *entity.InvitationStatus) error
}

// InvitationStatusRepository is an invitation status repository.
type InvitationStatusRepository struct {
	db *bun.DB
}

// NewInvitationStatusRepository creates an instance of invitationStatusRepository.
func NewInvitationStatusRepository(db *bun.DB) *InvitationStatusRepository {
	return &InvitationStatusRepository{
		db: db,
	}
}

// FindAllInvitationStatuses returns all invitation status.
func (r *InvitationStatusRepository) FindAllInvitationStatuses(ctx context.Context) ([]*entity.InvitationStatus, error) {
	result := make([]*entity.InvitationStatus, 0)

	if err := r.db.NewSelect().Model(&result).Scan(ctx); err != nil {
		return nil, err
	}

	return result, nil
}

// FindInvitationStatusByID finds an invitation status by ID.
func (r *InvitationStatusRepository) FindInvitationStatusByID(ctx context.Context, id string) (*entity.InvitationStatus, error) {
	result := &entity.InvitationStatus{}

	if err := r.db.NewSelect().Model(result).Where("id = UUID_TO_BIN(?)", id).Scan(ctx); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return result, nil
}

// FindPendingInvitationStatus finds the pending invitation status.
func (r *InvitationStatusRepository) FindPendingInvitationStatus(ctx context.Context) (*entity.InvitationStatus, error) {
	result := &entity.InvitationStatus{}

	if err := r.db.NewSelect().Model(result).Where("name = ?", "Pending").Scan(ctx); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return result, nil
}

// CreateInvitationStatus creates an invitation status.
func (r *InvitationStatusRepository) CreateInvitationStatus(ctx context.Context, invitationStatus *entity.InvitationStatus) error {
	_, err := r.db.NewInsert().Model(invitationStatus).Exec(ctx)
	return err
}
