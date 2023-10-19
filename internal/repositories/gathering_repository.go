package repositories

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
	"github.com/uptrace/bun"

	entity "hukum-onlen-go/entity"
)

// IGatheringRepository is a gathering type repository interface.
type IGatheringRepository interface {
	FindAllGatherings(ctx context.Context) ([]*entity.Gathering, error)
	FindGatheringByID(ctx context.Context, id string) (*entity.Gathering, error)
	CreateGathering(ctx context.Context, gathering *entity.Gathering) error
}

// GatheringRepository is a gathering type repository.
type GatheringRepository struct {
	db *bun.DB
}

// NewGatheringRepository creates an instance of GatheringRepository.
func NewGatheringRepository(db *bun.DB) *GatheringRepository {
	return &GatheringRepository{
		db: db,
	}
}

// FindAllGatherings returns all gathering types.
func (r *GatheringRepository) FindAllGatherings(ctx context.Context) ([]*entity.Gathering, error) {
	result := make([]*entity.Gathering, 0)

	if err := r.db.NewSelect().Model(&result).Scan(ctx); err != nil {
		return nil, err
	}

	return result, nil
}

// FindGatheringByID finds a gathering type by ID.
func (r *GatheringRepository) FindGatheringByID(ctx context.Context, id string) (*entity.Gathering, error) {
	result := &entity.Gathering{}

	if err := r.db.NewSelect().Model(result).Where("id = UUID_TO_BIN(?)", id).Scan(ctx); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return result, nil
}

// CreateGathering creates a gathering type.
func (r *GatheringRepository) CreateGathering(ctx context.Context, gathering *entity.Gathering) error {
	_, err := r.db.NewInsert().Model(gathering).Exec(ctx)
	return err
}
