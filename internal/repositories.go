package internal

import (
	"context"

	"github.com/google/uuid"
	"gowild.com/internal/domain"
)

type AnimalRepository interface {
	CreateAnimal(ctx context.Context, animals []domain.Animal, biome domain.Biome) error
	ReadAllAnimals(ctx context.Context) ([]domain.Animal, error)
	ReadAnimal(ctx context.Context, id uuid.UUID) (domain.Animal, error)
	ReadAnimalByNameAndBiome(ctx context.Context, name string, biome domain.Biome) (domain.Animal, error)
	UpdateAnimal(ctx context.Context, updated domain.Animal) error
	DeleteAnimal(ctx context.Context, id uuid.UUID) error
}

type PlantRepository interface {
	CreatePlants(ctx context.Context, animals []domain.Plant, biome domain.Biome) error
	ReadAllPlants(ctx context.Context) ([]domain.Plant, error)
	ReadPlant(ctx context.Context, id uuid.UUID) (domain.Plant, error)
	ReadPlantByNameAndBiome(ctx context.Context, name string, biome domain.Biome) (domain.Plant, error)
	UpdatePlant(ctx context.Context, updated domain.Plant) error
	DeletePlant(ctx context.Context, id uuid.UUID) error
}
