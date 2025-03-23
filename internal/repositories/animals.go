package repositories

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"gowild.com/internal"
	fauna "gowild.com/pkg/animals"
	"gowild.com/pkg/biomes"
)

type animalRepository struct {
	dbpool *pgxpool.Pool
}

func NewAnimalRepository(dbpool *pgxpool.Pool) animalRepository {
	return animalRepository{
		dbpool: dbpool,
	}
}

var _ internal.AnimalRepository = (*animalRepository)(nil)

func (r *animalRepository) CreateAnimal(ctx context.Context, animals []fauna.Animal, biome biomes.Biome) error {
	return fmt.Errorf("unimplemented method")
}

func (r *animalRepository) ReadAllAnimals(ctx context.Context) ([]fauna.Animal, error) {
	return nil, fmt.Errorf("unimplemented method")
}

func (r *animalRepository) ReadAnimal(ctx context.Context, id uuid.UUID) (fauna.Animal, error) {
	return fauna.Animal{}, fmt.Errorf("unimplemented method")
}

func (r *animalRepository) ReadAnimalByNameAndBiome(ctx context.Context, name string, biome biomes.Biome) (fauna.Animal, error) {
	return fauna.Animal{}, fmt.Errorf("unimplemented method")
}

func (r *animalRepository) UpdateAnimal(ctx context.Context, updated fauna.Animal) error {
	return fmt.Errorf("unimplemented method")
}

func (r *animalRepository) DeleteAnimal(ctx context.Context, id uuid.UUID) error {
	return fmt.Errorf("unimplemented method")
}
