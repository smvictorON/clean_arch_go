package applications

import (
	"clean_arch_go/src/domain"
)

type GetAllCars struct {
	carRepo domain.CarRepository
}

func NewGetAllCars(carRepo domain.CarRepository) *GetAllCars {
	return &GetAllCars{
		carRepo: carRepo,
	}
}

func (app *GetAllCars) Execute() ([]domain.Car, error) {
	return app.carRepo.ReadAll(), nil
}
