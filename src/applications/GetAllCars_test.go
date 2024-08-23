package applications

import (
	"clean_arch_go/src/domain"
	"clean_arch_go/src/infra/repositories"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllCars_Execute(test *testing.T) {
	repo := repositories.NewCarRepositoryInMemory()

	useCase := NewGetAllCars(repo)

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

	cars := useCase.Execute()

	assert.Equal(test, len(cars), 2)
	assert.Equal(test, cars[0].Year, car1.Year)
	assert.Equal(test, cars[0].Model, car1.Model)
	assert.Equal(test, cars[0].Brand, car1.Brand)
	assert.Equal(test, cars[0].Color, car1.Color)
	assert.Equal(test, cars[1].Year, car2.Year)
	assert.Equal(test, cars[1].Model, car2.Model)
	assert.Equal(test, cars[1].Brand, car2.Brand)
	assert.Equal(test, cars[1].Color, car2.Color)
}

func TestGetAllCars_Execute_MOCK(test *testing.T) {
	mockRepo := repositories.NewCarRepositoryInMemoryMock()

	useCase := NewGetAllCars(mockRepo)

	cars := []domain.Car{
		{
			Year:  2020,
			Model: "Model S",
			Brand: "Tesla",
			Color: "Red",
		},
		{
			Year:  2022,
			Model: "Model X",
			Brand: "Tesla",
			Color: "Blue",
		},
	}

	mockRepo.On("ReadAll").Return(cars)

	res := useCase.Execute()

	assert.Equal(test, len(res), 2)
	assert.Equal(test, res[0].Year, cars[0].Year)
	assert.Equal(test, res[0].Model, cars[0].Model)
	assert.Equal(test, res[0].Brand, cars[0].Brand)
	assert.Equal(test, res[0].Color, cars[0].Color)
	assert.Equal(test, res[1].Year, cars[1].Year)
	assert.Equal(test, res[1].Model, cars[1].Model)
	assert.Equal(test, res[1].Brand, cars[1].Brand)
	assert.Equal(test, res[1].Color, cars[1].Color)
}
