package applications

import (
	"clean_arch_go/src/domain"
)

type UpdateCar struct {
	carRepo domain.CarRepository
}

func NewUpdateCar(carRepo domain.CarRepository) *UpdateCar {
	return &UpdateCar{
		carRepo: carRepo,
	}
}

func (app *UpdateCar) Execute(carId string, car domain.Car) bool {
	return app.carRepo.Update(carId, car)
}
