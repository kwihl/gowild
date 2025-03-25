package handlers

import (
	"fmt"

	"github.com/google/uuid"
	"gowild.com/internal/domain"
)

type animalDTO struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name,omitempty"`
	Noise *string   `json:"noise,omitempty"`
}

func (a animalDTO) toDomain() domain.Animal {
	return domain.Animal{
		ID:    a.ID,
		Name:  a.Name,
		Noise: a.Noise,
	}
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

type plantDTO struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name,omitempty"`
	Size string    `json:"plantSize,omitempty"`
}

func (p plantDTO) toDomain() (domain.Plant, error) {

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

func plantsFromDomain(plants []domain.Plant) []plantDTO {
	dtos := make([]plantDTO, len(plants))
	for i, plant := range plants {
		dtos[i] = plantDTO{
			ID:   plant.ID,
			Name: plant.Name,
			Size: string(plant.Size),
		}
	}

	return dtos
}
