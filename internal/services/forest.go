package services

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"gowild.com/internal"
	"gowild.com/internal/domain"
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

func (s *ForestService) ListForestAnimals(ctx context.Context) ([]domain.Animal, error) {
	return s.animals.ReadAllAnimals(ctx)
}

func (s *ForestService) GetForestAnimal(ctx context.Context, id uuid.UUID) (domain.Animal, error) {
	return domain.Animal{}, fmt.Errorf("unimplemented method")
}

func (s *ForestService) ListForestPlants(ctx context.Context) ([]domain.Plant, error) {
	return nil, fmt.Errorf("unimplemented method")
}

func (s *ForestService) GetForestPlant(ctx context.Context, id uuid.UUID) (domain.Plant, error) {
	return domain.Plant{}, fmt.Errorf("unimplemented method")
}
