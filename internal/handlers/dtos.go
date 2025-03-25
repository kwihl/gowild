package handlers

import (
	"fmt"

	"github.com/google/uuid"
	"gowild.com/internal/domain"
)

type animalDTO struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name,omitempty"`
	Noise string    `json:"noise,omitempty"`
}

func (a animalDTO) toDomain() (domain.Animal, error) {
	return domain.Animal{
		ID:    a.ID,
		Name:  a.Name,
		Noise: a.Noise,
	}, nil
}

func animalsFromDomain(animals []domain.Animal) []animalDTO {
	dtos := make([]animalDTO, len(animals))
	for i, animal := range animals {
		dtos[i] = animalDTO{
			ID:    animal.ID,
			Name:  animal.Name,
			Noise: animal.Noise,
		}
	}

	return dtos
}

type PlantDTO struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name,omitempty"`
	Size string    `json:"plantSize,omitempty"`
}

func (p PlantDTO) toDomain() (domain.Plant, error) {

	plantSize, ok := domain.ToPlantSize(p.Size)
	if !ok {
		return domain.Plant{}, fmt.Errorf("invalid plant size")
	}

	return domain.Plant{
		ID:   p.ID,
		Name: p.Name,
		Size: plantSize,
	}, nil
}

func plantsFromDomain(plants []domain.Plant) []PlantDTO {
	dtos := make([]PlantDTO, len(plants))
	for i, plant := range plants {
		dtos[i] = PlantDTO{
			ID:   plant.ID,
			Name: plant.Name,
			Size: string(plant.Size),
		}
	}

	return dtos
}
