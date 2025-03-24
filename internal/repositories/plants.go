package repositories

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"gowild.com/internal"
	"gowild.com/internal/domain"
)

type plantRepository struct {
	dbpool *pgxpool.Pool
}

func NewPlantRepository(dbpool *pgxpool.Pool) plantRepository {
	return plantRepository{
		dbpool: dbpool,
	}
}

var _ internal.PlantRepository = (*plantRepository)(nil)

func (r *plantRepository) CreatePlants(ctx context.Context, Plants []domain.Plant, biome domain.Biome) error {
	return fmt.Errorf("unimplemented method")
}

func (r *plantRepository) ReadAllPlants(ctx context.Context) ([]domain.Plant, error) {
	return nil, fmt.Errorf("unimplemented method")
}

func (r *plantRepository) ReadPlant(ctx context.Context, id uuid.UUID) (domain.Plant, error) {
	return domain.Plant{}, fmt.Errorf("unimplemented method")
}

func (r *plantRepository) ReadPlantByNameAndBiome(ctx context.Context, name string, biome domain.Biome) (domain.Plant, error) {
	return domain.Plant{}, fmt.Errorf("unimplemented method")
}

func (r *plantRepository) UpdatePlant(ctx context.Context, updated domain.Plant) error {
	return fmt.Errorf("unimplemented method")
}

func (r *plantRepository) DeletePlant(ctx context.Context, id uuid.UUID) error {
	return fmt.Errorf("unimplemented method")
}
