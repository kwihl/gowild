package internal

import (
	"context"

	"github.com/google/uuid"
	fauna "gowild.com/pkg/animals"
	"gowild.com/pkg/biomes"
	flora "gowild.com/pkg/plants"
)

type AnimalRepository interface {
	CreateAnimal(ctx context.Context, animals []fauna.Animal, biome biomes.Biome) error
	ReadAllAnimals(ctx context.Context) ([]fauna.Animal, error)
	ReadAnimal(ctx context.Context, id uuid.UUID) (fauna.Animal, error)
	ReadAnimalByNameAndBiome(ctx context.Context, name string, biome biomes.Biome) (fauna.Animal, error)
	UpdateAnimal(ctx context.Context, updated fauna.Animal) error
	DeleteAnimal(ctx context.Context, id uuid.UUID) error
}

type PlantRepository interface {
	CreatePlants(ctx context.Context, animals []flora.Plant, biome biomes.Biome) error
	ReadAllPlants(ctx context.Context) ([]flora.Plant, error)
	ReadPlant(ctx context.Context, id uuid.UUID) (flora.Plant, error)
	ReadPlantByNameAndBiome(ctx context.Context, name string, biome biomes.Biome) (flora.Plant, error)
	UpdatePlant(ctx context.Context, updated flora.Plant) error
	DeletePlant(ctx context.Context, id uuid.UUID) error
}
