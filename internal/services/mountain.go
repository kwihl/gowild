package services

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"gowild.com/internal"
	"gowild.com/pkg/animals"
	"gowild.com/pkg/biomes"
	"gowild.com/pkg/plants"
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

func (s *MountainService) GetMountainAnimals(ctx context.Context) ([]animals.Animal, error) {
	return nil, fmt.Errorf("unimplemented method")
}

func (s *MountainService) GetMountainAnimal(ctx context.Context, id uuid.UUID) (animals.Animal, error) {
	return animals.Animal{}, fmt.Errorf("unimplemented method")
}

func (s *MountainService) GetMountainPlants(ctx context.Context) ([]plants.Plant, error) {
	return nil, fmt.Errorf("unimplemented method")
}

func (s *MountainService) GetMountainPlant(ctx context.Context, id uuid.UUID) (plants.Plant, error) {
	return plants.Plant{}, fmt.Errorf("unimplemented method")
}

func (s *MountainService) GetPeaks(ctx context.Context, minAltitude int) ([]biomes.MountainPeak, error) {
	return nil, fmt.Errorf("unimplemented method")
}
