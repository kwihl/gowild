package repositories

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"gowild.com/internal"
	"gowild.com/internal/domain"
)

type animalRepository struct {
	dbpool *pgxpool.Pool
}

func NewAnimalRepository(dbpool *pgxpool.Pool) animalRepository {
	return animalRepository{
		dbpool: dbpool,
	}
}

var _ internal.AnimalRepository = (*animalRepository)(nil)

func (r *animalRepository) CreateAnimal(ctx context.Context, animals []domain.Animal, biome domain.Biome) error {
	return fmt.Errorf("unimplemented method")
}

func (r *animalRepository) ReadAllAnimals(ctx context.Context) ([]domain.Animal, error) {
	rows, err := r.dbpool.Query(ctx, `
		SELECT
			id,
			name,
			noise
		FROM animal
		WHERE
			biome = 'forest';
	`)

	if err != nil {
		return nil, err
	}

	daos, err := pgx.CollectRows(rows, pgx.RowToStructByName[animalDAO])
	if err != nil {
		return nil, err
	}

	animals := make([]domain.Animal, len(daos))
	for i, dao := range daos {
		animals[i] = dao.toDomain()
	}

	return animals, nil
}

func (r *animalRepository) ReadAnimal(ctx context.Context, id uuid.UUID) (domain.Animal, error) {
	return domain.Animal{}, fmt.Errorf("unimplemented method")
}

func (r *animalRepository) ReadAnimalByNameAndBiome(ctx context.Context, name string, biome domain.Biome) (domain.Animal, error) {
	return domain.Animal{}, fmt.Errorf("unimplemented method")
}

func (r *animalRepository) UpdateAnimal(ctx context.Context, updated domain.Animal) error {
	return fmt.Errorf("unimplemented method")
}

func (r *animalRepository) DeleteAnimal(ctx context.Context, id uuid.UUID) error {
	return fmt.Errorf("unimplemented method")
}
