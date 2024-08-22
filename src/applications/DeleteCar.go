package applications

import (
	"clean_arch_go/src/domain"
)

type DeleteCar struct {
	carRepo domain.CarRepository
}

func NewDeleteCar(carRepo domain.CarRepository) *DeleteCar {
	return &DeleteCar{
		carRepo: carRepo,
	}
}

func (app *DeleteCar) Execute(carId string) bool {
	return app.carRepo.Delete(carId)
}
