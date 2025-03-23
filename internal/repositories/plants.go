package repositories

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"gowild.com/internal"
	"gowild.com/pkg/biomes"
	flora "gowild.com/pkg/plants"
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

func (r *plantRepository) CreatePlants(ctx context.Context, Plants []flora.Plant, biome biomes.Biome) error {
	return fmt.Errorf("unimplemented method")
}

func (r *plantRepository) ReadAllPlants(ctx context.Context) ([]flora.Plant, error) {
	return nil, fmt.Errorf("unimplemented method")
}

func (r *plantRepository) ReadPlant(ctx context.Context, id uuid.UUID) (flora.Plant, error) {
	return flora.Plant{}, fmt.Errorf("unimplemented method")
}

func (r *plantRepository) ReadPlantByNameAndBiome(ctx context.Context, name string, biome biomes.Biome) (flora.Plant, error) {
	return flora.Plant{}, fmt.Errorf("unimplemented method")
}

func (r *plantRepository) UpdatePlant(ctx context.Context, updated flora.Plant) error {
	return fmt.Errorf("unimplemented method")
}

func (r *plantRepository) DeletePlant(ctx context.Context, id uuid.UUID) error {
	return fmt.Errorf("unimplemented method")
}
