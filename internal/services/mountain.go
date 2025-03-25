package services

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"gowild.com/internal"
	"gowild.com/internal/domain"
)

type MountainService struct {
	animals internal.AnimalRepository
	plants  internal.PlantRepository
}

func NewMountainService(animalRepository internal.AnimalRepository, plantRepository internal.PlantRepository) MountainService {
	return MountainService{
		animals: animalRepository,
		plants:  plantRepository,
	}
}

var _ internal.MountainService = (*MountainService)(nil)

func (s *MountainService) GetMountainAnimals(ctx context.Context) ([]domain.Animal, error) {
	return nil, fmt.Errorf("unimplemented method")
}

func (s *MountainService) GetMountainAnimal(ctx context.Context, id uuid.UUID) (domain.Animal, error) {
	return domain.Animal{}, fmt.Errorf("unimplemented method")
}

func (s *MountainService) GetMountainPlants(ctx context.Context) ([]domain.Plant, error) {
	return nil, fmt.Errorf("unimplemented method")
}

func (s *MountainService) GetMountainPlant(ctx context.Context, id uuid.UUID) (domain.Plant, error) {
	return domain.Plant{}, fmt.Errorf("unimplemented method")
}

func (s *MountainService) GetPeaks(ctx context.Context, minAltitude int) ([]domain.MountainPeak, error) {
	return nil, fmt.Errorf("unimplemented method")
}
