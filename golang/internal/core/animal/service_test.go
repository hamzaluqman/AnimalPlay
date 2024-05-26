package animal

import (
	"errors"
	"strconv"
	"testing"

	"github.com/stretchr/testify/mock"
)

func TestAnimalService_GetAll(t *testing.T) {

	t.Run("GetAll error on repo GetAll", func(t *testing.T) {

		mockAnimalRepository := MockAnimalRepository{}
		errMsg := errors.New("Failed to GetAll from animal repo")
		mockAnimalRepository.On("GetAll", mock.Anything).Return([]*Animal{}, errMsg)

		service := NewService(&mockAnimalRepository)

		expectedAnimalsLen := 0

		animals, err := service.GetAll()

		if err == nil {
			t.Fatalf("Failed to get an error")
		}

		if err.Error() != errMsg.Error() {
			t.Fatalf("Expected %v, got %v", errMsg, err)
		}

		sliceLen := len(animals)

		if sliceLen != expectedAnimalsLen {
			t.Fatalf("Expected slice len %v, got %v", expectedAnimalsLen, sliceLen)
		}

		mockAnimalRepository.AssertExpectations(t)
	})

	t.Run("GetAll success, no animals", func(t *testing.T) {

		mockAnimalRepository := MockAnimalRepository{}
		mockAnimalRepository.On("GetAll", mock.Anything).Return([]*Animal{}, nil)

		expectedAnimalsLen := 0

		service := NewService(&mockAnimalRepository)

		animals, err := service.GetAll()

		if err != nil {
			t.Fatalf("Unexpected error. Expected %v, got %v", nil, err)
		}

		if len(animals) != expectedAnimalsLen {
			t.Fatalf("Unexpected length of animals slice")
		}

		mockAnimalRepository.AssertExpectations(t)
	})

	t.Run("GetAll success, one animal", func(t *testing.T) {

		mockAnimalRepository := MockAnimalRepository{}

		expectedAnimals := getTestAnimals(1)
		expectedAnimalsLen := 1

		mockAnimalRepository.On("GetAll", mock.Anything).Return(expectedAnimals, nil)

		service := NewService(&mockAnimalRepository)

		animals, err := service.GetAll()

		if err != nil {
			t.Fatalf("Unexpected error. Expected %v, got %v", nil, err)
		}

		if len(animals) != expectedAnimalsLen {
			t.Fatalf("Unexpected length of animals slice")
		}

		mockAnimalRepository.AssertExpectations(t)
	})

	t.Run("GetAll success, three animals", func(t *testing.T) {

		mockAnimalRepository := MockAnimalRepository{}

		expectedAnimals := getTestAnimals(3)
		expectedAnimalsLen := 3

		mockAnimalRepository.On("GetAll", mock.Anything).Return(expectedAnimals, nil)

		service := NewService(&mockAnimalRepository)

		animals, err := service.GetAll()

		if err != nil {
			t.Fatalf("Unexpected error. Expected %v, got %v", nil, err)
		}

		if len(animals) != expectedAnimalsLen {
			t.Fatalf("Unexpected length of animals slice")
		}

		mockAnimalRepository.AssertExpectations(t)
	})
}

type MockAnimalRepository struct {
	mock.Mock
}

func (aR *MockAnimalRepository) GetAll() ([]*Animal, error) {
	args := aR.Called()
	return args.Get(0).([]*Animal), args.Error(1)
}

func getTestAnimals(numOfAnimals int) []*Animal {
	var animals []*Animal

	for i := 1; i <= numOfAnimals; i++ {
		animal := NewAnimal()
		animal.Id = int32(i)
		animal.AnimalName = "Name " + strconv.Itoa(i)

		animals = append(animals, animal)
	}

	return animals
}
