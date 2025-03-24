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

func (a *animalDTO) toDomain() (domain.Animal, error) {
	return domain.Animal{
		ID:    a.ID,
		Name:  a.Name,
		Noise: a.Noise,
	}, nil
}

type plantDTO struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name,omitempty"`
	Size string    `json:"plantSize,omitempty"`
}

func (a *plantDTO) toDomain() (domain.Plant, error) {

	plantSize, ok := domain.ToPlantSize(a.Size)
	if !ok {
		return domain.Plant{}, fmt.Errorf("invalid plant size")
	}

	return domain.Plant{
		ID:   a.ID,
		Name: a.Name,
		Size: plantSize,
	}, nil
}
