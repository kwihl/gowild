package internal

import (
	"context"

	"github.com/google/uuid"
	"gowild.com/pkg/animals"
	"gowild.com/pkg/biomes"
	"gowild.com/pkg/plants"
)

type ForestService interface {
	ListForestAnimals(ctx context.Context) ([]animals.Animal, error)
	GetForestAnimal(ctx context.Context, id uuid.UUID) (animals.Animal, error)
	ListForestPlants(ctx context.Context) ([]plants.Plant, error)
	GetForestPlant(ctx context.Context, id uuid.UUID) (plants.Plant, error)
}

type MountainService interface {
	GetMountainAnimals(ctx context.Context) ([]animals.Animal, error)
	GetMountainAnimal(ctx context.Context, id uuid.UUID) (animals.Animal, error)
	GetMountainPlants(ctx context.Context) ([]plants.Plant, error)
	GetMountainPlant(ctx context.Context, id uuid.UUID) (plants.Plant, error)
	GetPeaks(ctx context.Context, minAltitude int) ([]biomes.MountainPeak, error)
}
