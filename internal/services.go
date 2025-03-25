package internal

import (
	"context"

	"github.com/google/uuid"
	"gowild.com/internal/domain"
)

type ForestService interface {
	ListForestAnimals(ctx context.Context) ([]domain.Animal, error)
	GetForestAnimal(ctx context.Context, id uuid.UUID) (domain.Animal, error)
	ListForestPlants(ctx context.Context) ([]domain.Plant, error)
	GetForestPlant(ctx context.Context, id uuid.UUID) (domain.Plant, error)
}

type MountainService interface {
	GetMountainAnimals(ctx context.Context) ([]domain.Animal, error)
	GetMountainAnimal(ctx context.Context, id uuid.UUID) (domain.Animal, error)
	GetMountainPlants(ctx context.Context) ([]domain.Plant, error)
	GetMountainPlant(ctx context.Context, id uuid.UUID) (domain.Plant, error)
	GetPeaks(ctx context.Context, minAltitude int) ([]domain.MountainPeak, error)
}
