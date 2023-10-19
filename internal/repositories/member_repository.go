package repositories

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
	"github.com/uptrace/bun"

	entity "hukum-onlen-go/entity"
)

// IMemberRepository is a member repository interface.
type IMemberRepository interface {
	FindAllMembers(ctx context.Context, email string) ([]*entity.Member, error)
	FindMemberByID(ctx context.Context, id string) (*entity.Member, error)
	CreateMember(ctx context.Context, member *entity.Member) error
}

// MemberRepository is a member repository.
type MemberRepository struct {
	db *bun.DB
}

// NewMemberRepository creates an instance of MemberRepository.
func NewMemberRepository(db *bun.DB) *MemberRepository {
	return &MemberRepository{
		db: db,
	}
}

// FindAllMembers returns all members.
func (r *MemberRepository) FindAllMembers(ctx context.Context, email string) ([]*entity.Member, error) {
	result := make([]*entity.Member, 0)

	query := r.db.NewSelect().Model(&result)
	if email != "" {
		query = query.Where("email = ?", email)
	}

	if err := query.Scan(ctx); err != nil {
		return nil, err
	}

	return result, nil
}

// FindMemberByID finds a member by ID.
func (r *MemberRepository) FindMemberByID(ctx context.Context, id string) (*entity.Member, error) {
	result := &entity.Member{}

	if err := r.db.NewSelect().Model(result).Where("id = UUID_TO_BIN(?)", id).Scan(ctx); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return result, nil
}

// CreateMember creates a member.
func (r *MemberRepository) CreateMember(ctx context.Context, member *entity.Member) error {
	_, err := r.db.NewInsert().Model(member).Exec(ctx)
	return err
}
