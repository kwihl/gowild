package services

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"gowild.com/internal"
	"gowild.com/pkg/animals"
	"gowild.com/pkg/plants"
)

type ForestService struct {
	animals internal.AnimalRepository
	plants  internal.PlantRepository
}

func NewForestService(animalRepository internal.AnimalRepository, plantRepository internal.PlantRepository) ForestService {
	return ForestService{
		animals: animalRepository,
		plants:  plantRepository,
	}
}

var _ internal.ForestService = (*ForestService)(nil)

func (s *ForestService) ListForestAnimals(ctx context.Context) ([]animals.Animal, error) {
	return nil, fmt.Errorf("unimplemented method")
}

func (s *ForestService) GetForestAnimal(ctx context.Context, id uuid.UUID) (animals.Animal, error) {
	return animals.Animal{}, fmt.Errorf("unimplemented method")
}

func (s *ForestService) ListForestPlants(ctx context.Context) ([]plants.Plant, error) {
	return nil, fmt.Errorf("unimplemented method")
}

func (s *ForestService) GetForestPlant(ctx context.Context, id uuid.UUID) (plants.Plant, error) {
	return plants.Plant{}, fmt.Errorf("unimplemented method")
}
