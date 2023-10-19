package repositories

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
	"github.com/uptrace/bun"

	entity "hukum-onlen-go/entity"
)

// IGatheringTypeRepository is a gathering type repository interface.
type IGatheringTypeRepository interface {
	FindAllGatheringTypes(ctx context.Context) ([]*entity.GatheringType, error)
	FindGatheringTypeByID(ctx context.Context, id string) (*entity.GatheringType, error)
	CreateGatheringType(ctx context.Context, gatheringType *entity.GatheringType) error
}

// GatheringTypeRepository is a gathering type repository.
type GatheringTypeRepository struct {
	db *bun.DB
}

// NewGatheringTypeRepository creates an instance of GatheringTypeRepository.
func NewGatheringTypeRepository(db *bun.DB) *GatheringTypeRepository {
	return &GatheringTypeRepository{
		db: db,
	}
}

// FindAllGatheringTypes returns all gathering types.
func (r *GatheringTypeRepository) FindAllGatheringTypes(ctx context.Context) ([]*entity.GatheringType, error) {
	result := make([]*entity.GatheringType, 0)

	if err := r.db.NewSelect().Model(&result).Scan(ctx); err != nil {
		return nil, err
	}

	return result, nil
}

// FindGatheringTypeByID finds a gathering type by ID.
func (r *GatheringTypeRepository) FindGatheringTypeByID(ctx context.Context, id string) (*entity.GatheringType, error) {
	result := &entity.GatheringType{}

	if err := r.db.NewSelect().Model(result).Where("id = UUID_TO_BIN(?)", id).Scan(ctx); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return result, nil
}

// CreateGatheringType creates a gathering type.
func (r *GatheringTypeRepository) CreateGatheringType(ctx context.Context, gatheringType *entity.GatheringType) error {
	_, err := r.db.NewInsert().Model(gatheringType).Exec(ctx)
	return err
}
