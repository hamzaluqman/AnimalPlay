package mysql_repository

import (
	"animal-play/internal/core/animal"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type AnimalRepository struct {
	db        *sqlx.DB
	tableName string
}

func NewAnimalRepository(db *sqlx.DB) *AnimalRepository {
	tableName := animal.TableName

	return &AnimalRepository{
		db:        db,
		tableName: tableName,
	}
}

func (r *AnimalRepository) GetAll() ([]*animal.Animal, error) {
	var animals []*animal.Animal

	query := fmt.Sprintf("SELECT * FROM %v", r.tableName)
	err := r.db.Select(&animals, query)

	if err != nil {
		return animals, err
	}

	return animals, nil
}
