package handlers

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/google/uuid"
	fauna "gowild.com/pkg/animals"
	flora "gowild.com/pkg/plants"
)

func toDTO[T any](body io.ReadCloser) (*T, error) {
	decoder := json.NewDecoder(body)
	var dto T
	if err := decoder.Decode(&dto); err != nil {
		return nil, err
	}

	fmt.Println(dto)
	return &dto, nil
}

type animalDTO struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Noise string    `json:"noise"`
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

func (a *animalDTO) toDomain() (fauna.Animal, error) {
	return fauna.Animal{
		ID:    a.ID,
		Name:  a.Name,
		Noise: a.Noise,
	}, nil
}
