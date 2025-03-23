package plants

import "github.com/google/uuid"

type PlantSize string

const (
	Small  PlantSize = "small"
	Medium PlantSize = "medium"
	Large  PlantSize = "large"
)

func ToPlantSize(s string) (PlantSize, bool) {

	switch PlantSize(s) {
	case Small:
		return Small, true
	case Medium:
		return Medium, true
	case Large:
		return Large, true
	default:
		return PlantSize(""), false
	}
}

type Plant struct {
	ID   uuid.UUID
	Name string
	Size PlantSize
}
