package handlers

import (
	"fmt"

	"github.com/google/uuid"
	fauna "gowild.com/pkg/animals"
	flora "gowild.com/pkg/plants"
)

type animalDTO struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Noise string    `json:"noise"`
}

func (a *animalDTO) toDomain() (fauna.Animal, error) {
	return fauna.Animal{
		ID:    a.ID,
		Name:  a.Name,
		Noise: a.Noise,
	}, nil
}

type plantDTO struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Size string    `json:"plantSize"`
}

func (a *plantDTO) toDomain() (flora.Plant, error) {

	plantSize, ok := flora.ToPlantSize(a.Size)
	if !ok {
		return flora.Plant{}, fmt.Errorf("invalid plant size")
	}

	return flora.Plant{
		ID:   a.ID,
		Name: a.Name,
		Size: plantSize,
	}, nil
}
