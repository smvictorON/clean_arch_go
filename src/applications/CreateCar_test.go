package applications

import (
	"clean_arch_go/src/domain"
	"clean_arch_go/src/infra/repositories"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateCar_Execute(test *testing.T) {
	repo := repositories.NewCarRepositoryInMemory()

	useCase := NewCreateCar(repo)

	car := domain.Car{
		Year:  2020,
		Model: "Model S",
		Brand: "Tesla",
		Color: "Red",
	}

	id := useCase.Execute(car)

	_, parseErr := uuid.Parse(id)
	assert.NoError(test, parseErr, "ID should be a valid UUID")
}

func TestCreateCar_Execute_MOCK(test *testing.T) {
	mockRepo := repositories.NewCarRepositoryInMemoryMock()

	useCase := NewCreateCar(mockRepo)

	car := domain.Car{
		Id:    uuid.New().String(),
		Year:  2020,
		Model: "Model S",
		Brand: "Tesla",
		Color: "Red",
	}

	mockRepo.On("Create", car).Return(car.Id)

	id := useCase.Execute(car)

	_, parseErr := uuid.Parse(id)
	assert.NoError(test, parseErr, "ID should be a valid UUID")
	mockRepo.AssertCalled(test, "Create", car)
}
