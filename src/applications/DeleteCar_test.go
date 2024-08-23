package applications

import (
	"clean_arch_go/src/domain"
	"clean_arch_go/src/infra/repositories"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestDeleteCar_Execute(test *testing.T) {
	repo := repositories.NewCarRepositoryInMemory()

	car1 := domain.Car{
		Year:  2020,
		Model: "Model S",
		Brand: "Tesla",
		Color: "Red",
	}

	car2 := domain.Car{
		Year:  2021,
		Model: "Model E",
		Brand: "Tesla",
		Color: "Blue",
	}

	car3 := domain.Car{
		Year:  2022,
		Model: "Model S",
		Brand: "Tesla",
		Color: "Black",
	}

	NewCreateCar(repo).Execute(car1)
	NewCreateCar(repo).Execute(car2)
	NewCreateCar(repo).Execute(car3)

	cars := NewGetAllCars(repo).Execute()

	assert.Equal(test, 3, len(cars))

	res := NewDeleteCar(repo).Execute(cars[1].Id)

	assert.Equal(test, true, res)

	cars = NewGetAllCars(repo).Execute()

	assert.Equal(test, 2, len(cars))
	assert.Equal(test, car1.Year, cars[0].Year)
	assert.Equal(test, car1.Model, cars[0].Model)
	assert.Equal(test, car1.Brand, cars[0].Brand)
	assert.Equal(test, car1.Color, cars[0].Color)
	assert.Equal(test, car3.Year, cars[1].Year)
	assert.Equal(test, car3.Model, cars[1].Model)
	assert.Equal(test, car3.Brand, cars[1].Brand)
	assert.Equal(test, car3.Color, cars[1].Color)
}

func TestDeleteCar_Execute_Err(test *testing.T) {
	repo := repositories.NewCarRepositoryInMemory()

	car1 := domain.Car{
		Year:  2020,
		Model: "Model S",
		Brand: "Tesla",
		Color: "Red",
	}

	car2 := domain.Car{
		Year:  2021,
		Model: "Model E",
		Brand: "Tesla",
		Color: "Blue",
	}

	car3 := domain.Car{
		Year:  2022,
		Model: "Model S",
		Brand: "Tesla",
		Color: "Black",
	}

	NewCreateCar(repo).Execute(car1)
	NewCreateCar(repo).Execute(car2)
	NewCreateCar(repo).Execute(car3)

	cars := NewGetAllCars(repo).Execute()

	assert.Equal(test, 3, len(cars))

	res := NewDeleteCar(repo).Execute("")

	assert.Equal(test, false, res)

	cars = NewGetAllCars(repo).Execute()

	assert.Equal(test, 3, len(cars))
}

func TestDeleteCar_Execute_MOCK(test *testing.T) {
	mockRepo := repositories.NewCarRepositoryInMemoryMock()

	mockRepo.On("Delete").Return(true)

	res := NewDeleteCar(mockRepo).Execute(uuid.New().String())

	assert.Equal(test, true, res)
}

func TestDeleteCar_Execute_MOCK_Err(test *testing.T) {
	mockRepo := repositories.NewCarRepositoryInMemoryMock()

	mockRepo.On("Delete").Return(false)

	res := NewDeleteCar(mockRepo).Execute(uuid.New().String())

	assert.Equal(test, false, res)
}
