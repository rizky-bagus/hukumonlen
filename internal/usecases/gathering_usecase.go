package usecases

import (
	"context"
	"time"

	"hukum-onlen-go/entity"
	"hukum-onlen-go/internal/repositories"
)

// IGatheringUsecase is a gathering use case interface.
type IGatheringUsecase interface {
	FindAllInvitationStatuses(ctx context.Context) ([]*entity.InvitationStatus, error)
	FindInvitationStatusByID(ctx context.Context, id string) (*entity.InvitationStatus, error)
	CreateInvitationStatus(ctx context.Context, name string) error

	FindAllGatheringTypes(ctx context.Context) ([]*entity.GatheringType, error)
	FindGatheringTypeByID(ctx context.Context, id string) (*entity.GatheringType, error)
	CreateGatheringType(ctx context.Context, name string) error

	FindAllGatherings(ctx context.Context) ([]*entity.Gathering, error)
	FindGatheringByID(ctx context.Context, id string) (*entity.Gathering, error)
	CreateGathering(
		ctx context.Context,
		name,
		location string,
		creatorID,
		typeID entity.UUID,
		scheduledAt time.Time,
		attendeeIDs []entity.UUID,
	) error
}

// GatheringUsecase is a gathering use case.
type GatheringUsecase struct {
	memberRepository           repositories.IMemberRepository
	invitationStatusRepository repositories.IInvitationStatusRepository
	gatheringTypeRepository    repositories.IGatheringTypeRepository
	invitationRepository       repositories.IInvitationRepository
	gatheringRepository        repositories.IGatheringRepository
}

// NewGatheringUsecase creates an instance of GatheringUsecase.
func NewGatheringUsecase(
	memberRepository repositories.IMemberRepository,
	invitationStatusRepository repositories.IInvitationStatusRepository,
	gatheringTypeRepository repositories.IGatheringTypeRepository,
	invitationRepository repositories.IInvitationRepository,
	gatheringRepository repositories.IGatheringRepository,
) *GatheringUsecase {
	return &GatheringUsecase{
		memberRepository:           memberRepository,
		invitationStatusRepository: invitationStatusRepository,
		gatheringTypeRepository:    gatheringTypeRepository,
		invitationRepository:       invitationRepository,
		gatheringRepository:        gatheringRepository,
	}
}

// FindAllInvitationStatuses returns all invitation statuses.
func (g *GatheringUsecase) FindAllInvitationStatuses(ctx context.Context) ([]*entity.InvitationStatus, error) {
	return g.invitationStatusRepository.FindAllInvitationStatuses(ctx)
}

// FindInvitationStatusByID finds an invitation status by ID.
func (g *GatheringUsecase) FindInvitationStatusByID(ctx context.Context, id string) (*entity.InvitationStatus, error) {
	return g.invitationStatusRepository.FindInvitationStatusByID(ctx, id)
}

// CreateInvitationStatus creates a new invitation status.
func (g *GatheringUsecase) CreateInvitationStatus(ctx context.Context, name string) error {
	toBeCreated := &entity.InvitationStatus{
		ID:   entity.NewUUID(),
		Name: name,
	}

	return g.invitationStatusRepository.CreateInvitationStatus(ctx, toBeCreated)
}

// FindAllGatheringTypes returns all gathering types.
func (g *GatheringUsecase) FindAllGatheringTypes(ctx context.Context) ([]*entity.GatheringType, error) {
	return g.gatheringTypeRepository.FindAllGatheringTypes(ctx)
}

// FindGatheringTypeByID finds a gathering type by ID.
func (g *GatheringUsecase) FindGatheringTypeByID(ctx context.Context, id string) (*entity.GatheringType, error) {
	return g.gatheringTypeRepository.FindGatheringTypeByID(ctx, id)
}

// CreateGatheringType creates a new gathering type.
func (g *GatheringUsecase) CreateGatheringType(ctx context.Context, name string) error {
	toBeCreated := &entity.GatheringType{
		ID:   entity.NewUUID(),
		Name: name,
	}

	return g.gatheringTypeRepository.CreateGatheringType(ctx, toBeCreated)
}

// FindAllGatherings returns all gatherings.
func (g *GatheringUsecase) FindAllGatherings(ctx context.Context) ([]*entity.Gathering, error) {
	return g.gatheringRepository.FindAllGatherings(ctx)
}

// FindGatheringByID finds a gathering by ID.
func (g *GatheringUsecase) FindGatheringByID(ctx context.Context, id string) (*entity.Gathering, error) {
	return g.gatheringRepository.FindGatheringByID(ctx, id)
}

// CreateGathering creates a new gathering.
func (g *GatheringUsecase) CreateGathering(
	ctx context.Context,
	name,
	location string,
	creatorID,
	typeID entity.UUID,
	scheduledAt time.Time,
	attendeeIDs []entity.UUID,
) error {
	gathering := &entity.Gathering{
		ID:          entity.NewUUID(),
		Name:        name,
		Location:    location,
		CreatorID:   creatorID,
		TypeID:      typeID,
		ScheduledAt: scheduledAt,
	}

	if err := g.gatheringRepository.CreateGathering(ctx, gathering); err != nil {
		return err
	}

	if len(attendeeIDs) == 0 {
		return nil
	}

	pendingStatus, err := g.invitationStatusRepository.FindPendingInvitationStatus(ctx)
	if err != nil {
		return err
	}

	invitations := make([]*entity.Invitation, 0, len(attendeeIDs))
	for _, attendeeID := range attendeeIDs {
		invitations = append(invitations, &entity.Invitation{
			ID:          entity.NewUUID(),
			MemberID:    attendeeID,
			GatheringID: gathering.ID,
			StatusID:    pendingStatus.ID,
		})
	}

	return g.invitationRepository.CreateInvitations(ctx, invitations)
}
