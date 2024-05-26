package animal

import "time"

const TableName = "animal"

type Animal struct {
	Id          int32  `db:"id"`
	DateCreated string `db:"dateCreated"`
	DateUpdated string `db:"dateUpdated"`
	AnimalName  string `db:"animalName"`
	AnimalType  string `db:"animalType"`
}

func NewAnimal() *Animal {
	return &Animal{
		Id:          0,
		DateCreated: time.Now().UTC().String(),
		DateUpdated: time.Now().UTC().String(),
		AnimalName:  "",
		AnimalType:  "",
	}
}
