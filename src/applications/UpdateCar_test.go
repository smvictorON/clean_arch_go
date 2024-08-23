package applications

import (
	"clean_arch_go/src/domain"
	"clean_arch_go/src/infra/repositories"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUpdateCar_Execute(test *testing.T) {
	repo := repositories.NewCarRepositoryInMemory()

	car := domain.Car{
		Year:  2020,
		Model: "Model S",
		Brand: "Tesla",
		Color: "Red",
	}

	id := NewCreateCar(repo).Execute(car)

	car.Brand = "BYD"
	car.Model = "Seal"
	car.Color = "Gray"
	car.Year = 2024

	res := NewUpdateCar(repo).Execute(id, car)

	assert.Equal(test, true, res)

	resCar, err := NewGetCarById(repo).Execute(id)

	assert.NoError(test, err)
	_, parseErr := uuid.Parse(id)
	assert.NoError(test, parseErr, "ID should be a valid UUID")
	assert.Equal(test, resCar.Year, car.Year)
	assert.Equal(test, resCar.Model, car.Model)
	assert.Equal(test, resCar.Brand, car.Brand)
	assert.Equal(test, resCar.Color, car.Color)
}

func TestUpdateCar_Execute_Err(test *testing.T) {
	repo := repositories.NewCarRepositoryInMemory()

	car := domain.Car{
		Year:  2020,
		Model: "Model S",
		Brand: "Tesla",
		Color: "Red",
	}

	id := NewCreateCar(repo).Execute(car)

	car.Brand = "BYD"
	car.Model = "Seal"
	car.Color = "Gray"
	car.Year = 2024

	res := NewUpdateCar(repo).Execute("", car)

	assert.Equal(test, false, res)

	resCar, err := NewGetCarById(repo).Execute(id)

	assert.NoError(test, err)
	_, parseErr := uuid.Parse(id)
	assert.NoError(test, parseErr, "ID should be a valid UUID")
	assert.Equal(test, resCar.Year, 2020)
	assert.Equal(test, resCar.Model, "Model S")
	assert.Equal(test, resCar.Brand, "Tesla")
	assert.Equal(test, resCar.Color, "Red")
}

func TestUpdateCar_Execute_MOCK(test *testing.T) {
	mockRepo := repositories.NewCarRepositoryInMemoryMock()

	useCase := NewUpdateCar(mockRepo)

	car := domain.Car{
		Id:    uuid.New().String(),
		Year:  2020,
		Model: "Model S",
		Brand: "Tesla",
		Color: "Red",
	}

	mockRepo.On("Update").Return(true)

	res := useCase.Execute(car.Id, domain.Car{
		Id:    car.Id,
		Year:  2024,
		Model: "Seal",
		Brand: "BYD",
		Color: "Gray",
	})

	assert.Equal(test, true, res)
}

func TestUpdateCar_Execute_MOCK_Err(test *testing.T) {
	mockRepo := repositories.NewCarRepositoryInMemoryMock()

	useCase := NewUpdateCar(mockRepo)

	car := domain.Car{
		Id:    uuid.New().String(),
		Year:  2020,
		Model: "Model S",
		Brand: "Tesla",
		Color: "Red",
	}

	mockRepo.On("Update").Return(false)

	res := useCase.Execute("", domain.Car{
		Id:    car.Id,
		Year:  2024,
		Model: "Seal",
		Brand: "BYD",
		Color: "Gray",
	})

	assert.Equal(test, false, res)
}
