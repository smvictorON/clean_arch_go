package applications

import (
	"clean_arch_go/src/domain"
)

type CreateCar struct {
	carRepo domain.CarRepository
}

func NewCreateCar(carRepo domain.CarRepository) *CreateCar {
	return &CreateCar{
		carRepo: carRepo,
	}
}

func (app *CreateCar) Execute(car domain.Car) string {
	return app.carRepo.Create(car)
}
