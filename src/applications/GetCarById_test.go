package applications

import (
	"clean_arch_go/src/domain"
	"clean_arch_go/src/infra/repositories"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetCarById_Execute(test *testing.T) {
	repo := repositories.NewCarRepositoryInMemory()

	useCase := NewGetCarById(repo)

	car1 := domain.Car{
		Year:  2020,
		Model: "Model S",
		Brand: "Tesla",
		Color: "Red",
	}

	car2 := domain.Car{
		Year:  2022,
		Model: "Model X",
		Brand: "Tesla",
		Color: "Blue",
	}

	id := NewCreateCar(repo).Execute(car1)
	NewCreateCar(repo).Execute(car2)

	car, err := useCase.Execute(id)

	assert.NoError(test, err)
	_, parseErr := uuid.Parse(id)
	assert.NoError(test, parseErr, "ID should be a valid UUID")
	assert.Equal(test, car.Year, car1.Year)
	assert.Equal(test, car.Model, car1.Model)
	assert.Equal(test, car.Brand, car1.Brand)
	assert.Equal(test, car.Color, car1.Color)
}

func TestGetCarById_Execute_Err(test *testing.T) {
	repo := repositories.NewCarRepositoryInMemory()

	useCase := NewGetCarById(repo)

	car1 := domain.Car{
		Year:  2020,
		Model: "Model S",
		Brand: "Tesla",
		Color: "Red",
	}

	car2 := domain.Car{
		Year:  2022,
		Model: "Model X",
		Brand: "Tesla",
		Color: "Blue",
	}

	NewCreateCar(repo).Execute(car1)
	NewCreateCar(repo).Execute(car2)

	res, err := useCase.Execute("")

	assert.Error(test, err)
	assert.Equal(test, err.Error(), "car not found")
	assert.Nil(test, res)
}

func TestGetCarById_Execute_MOCK(test *testing.T) {
	mockRepo := repositories.NewCarRepositoryInMemoryMock()

	useCase := NewGetCarById(mockRepo)

	cars := []domain.Car{
		{
			Id:    uuid.New().String(),
			Year:  2020,
			Model: "Model S",
			Brand: "Tesla",
			Color: "Red",
		},
		{
			Id:    uuid.New().String(),
			Year:  2022,
			Model: "Model X",
			Brand: "Tesla",
			Color: "Blue",
		},
	}

	mockRepo.On("ReadOne").Return(&cars[0], nil)

	res, err := useCase.Execute(cars[0].Id)

	assert.NoError(test, err)
	_, parseErr := uuid.Parse(cars[0].Id)
	assert.NoError(test, parseErr, "ID should be a valid UUID")
	assert.Equal(test, res, &cars[0])
}

func TestGetCarById_Execute_MOCK_Err(test *testing.T) {
	mockRepo := repositories.NewCarRepositoryInMemoryMock()

	useCase := NewGetCarById(mockRepo)

	mockRepo.On("ReadOne").Return(nil, errors.New("car not found"))

	res, err := useCase.Execute("")

	assert.Error(test, err)
	assert.Equal(test, err.Error(), "car not found")
	assert.Nil(test, res)
}
