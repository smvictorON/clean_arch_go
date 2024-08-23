package applications

import (
	"clean_arch_go/src/domain"
	"clean_arch_go/src/infra/repositories"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCarsByModel_Execute(test *testing.T) {
	repo := repositories.NewCarRepositoryInMemory()

	useCase := NewGetCarsByModel(repo)

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

	car3 := domain.Car{
		Year:  2022,
		Model: "model x",
		Brand: "Tesla",
		Color: "Yellow",
	}

	car4 := domain.Car{
		Year:  2022,
		Model: "MODEL X",
		Brand: "Tesla",
		Color: "White",
	}

	NewCreateCar(repo).Execute(car1)
	NewCreateCar(repo).Execute(car2)
	NewCreateCar(repo).Execute(car3)
	NewCreateCar(repo).Execute(car4)

	cars := useCase.Execute("MoDeL x")

	assert.Equal(test, len(cars), 3)
	assert.Equal(test, cars[0].Year, car2.Year)
	assert.Equal(test, cars[0].Model, car2.Model)
	assert.Equal(test, cars[0].Brand, car2.Brand)
	assert.Equal(test, cars[0].Color, car2.Color)
	assert.Equal(test, cars[1].Year, car3.Year)
	assert.Equal(test, cars[1].Model, car3.Model)
	assert.Equal(test, cars[1].Brand, car3.Brand)
	assert.Equal(test, cars[1].Color, car3.Color)
	assert.Equal(test, cars[2].Year, car4.Year)
	assert.Equal(test, cars[2].Model, car4.Model)
	assert.Equal(test, cars[2].Brand, car4.Brand)
	assert.Equal(test, cars[2].Color, car4.Color)
}

func TestGetCarsByModel_Execute_MOCK(test *testing.T) {
	mockRepo := repositories.NewCarRepositoryInMemoryMock()

	useCase := NewGetCarsByModel(mockRepo)

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
		{
			Year:  2022,
			Model: "model x",
			Brand: "Tesla",
			Color: "Yellow",
		},
		{
			Year:  2022,
			Model: "MODEL X",
			Brand: "Tesla",
			Color: "White",
		},
	}

	mockRepo.On("ReadByModel").Return([]domain.Car{cars[1], cars[2], cars[3]})

	res := useCase.Execute("MoDeL x")

	assert.Equal(test, 3, len(res))
	assert.Equal(test, res[0].Year, cars[1].Year)
	assert.Equal(test, res[0].Model, cars[1].Model)
	assert.Equal(test, res[0].Brand, cars[1].Brand)
	assert.Equal(test, res[0].Color, cars[1].Color)
	assert.Equal(test, res[1].Year, cars[2].Year)
	assert.Equal(test, res[1].Model, cars[2].Model)
	assert.Equal(test, res[1].Brand, cars[2].Brand)
	assert.Equal(test, res[1].Color, cars[2].Color)
	assert.Equal(test, res[2].Year, cars[3].Year)
	assert.Equal(test, res[2].Model, cars[3].Model)
	assert.Equal(test, res[2].Brand, cars[3].Brand)
	assert.Equal(test, res[2].Color, cars[3].Color)
}
