package applications

import (
	"clean_arch_go/src/domain"

	"github.com/google/uuid"
)

type CreateCar struct {
	carRepo domain.CarRepository
}

func NewCreateCar(carRepo domain.CarRepository) *CreateCar {
	return &CreateCar{
		carRepo: carRepo,
	}
}

func (app *CreateCar) Execute(car domain.Car) (string, error) {
	car.Id = uuid.New().String()
	return app.carRepo.Create(car), nil
}
