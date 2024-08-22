package applications

import (
	"clean_arch_go/src/domain"
)

type GetCarById struct {
	carRepo domain.CarRepository
}

func NewGetCarById(carRepo domain.CarRepository) *GetCarById {
	return &GetCarById{
		carRepo: carRepo,
	}
}

func (app *GetCarById) Execute(carId string) (domain.Car, error) {
	return app.carRepo.ReadOne(carId)
}
