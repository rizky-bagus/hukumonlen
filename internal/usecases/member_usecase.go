package usecases

import (
	"context"

	entity "hukum-onlen-go/entity"
	"hukum-onlen-go/internal/repositories"
)

// IMemberUsecase is a member usecase interface.
type IMemberUsecase interface {
	FindAllMembers(ctx context.Context, email string) ([]*entity.Member, error)
	FindMemberByID(ctx context.Context, id string) (*entity.Member, error)
	CreateMember(ctx context.Context, email, firstName, lastName string) error
}

// MemberUsecase is a member usecase.
type MemberUsecase struct {
	memberRepository repositories.IMemberRepository
}

// NewMemberUsecase creates an instance of MemberUsecase.
func NewMemberUsecase(memberRepository repositories.IMemberRepository) *MemberUsecase {
	return &MemberUsecase{
		memberRepository: memberRepository,
	}
}

// FindAllMembers returns all members.
func (g *MemberUsecase) FindAllMembers(ctx context.Context, email string) ([]*entity.Member, error) {
	return g.memberRepository.FindAllMembers(ctx, email)
}

// FindMemberByID finds an invitation status by ID.
func (g *MemberUsecase) FindMemberByID(ctx context.Context, id string) (*entity.Member, error) {
	return g.memberRepository.FindMemberByID(ctx, id)
}

// CreateMember creates a new member.
func (g *MemberUsecase) CreateMember(ctx context.Context, email, firstName, lastName string) error {
	toBeCreated := &entity.Member{
		ID:        entity.NewUUID(),
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
	}

	return g.memberRepository.CreateMember(ctx, toBeCreated)
}
