package repositories

import (
	"fmt"

	"github.com/google/uuid"
	"gowild.com/internal/domain"
)

type animalDAO struct {
	ID    uuid.UUID
	Name  string
	Noise *string
}

func (a animalDAO) toDomain() domain.Animal {
	return domain.Animal{
		ID:    a.ID,
		Name:  a.Name,
		Noise: a.Noise,
	}
}

type plantDAO struct {
	ID   uuid.UUID
	Name string
	Size string
}

func (p plantDAO) toDomain() (domain.Plant, error) {

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
