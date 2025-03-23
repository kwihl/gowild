package animals

import "github.com/google/uuid"

type Animal struct {
	ID    uuid.UUID
	Name  string
	Noise string
}
