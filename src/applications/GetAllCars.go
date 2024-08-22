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

func (app *GetAllCars) Execute() []domain.Car {
	cars := app.carRepo.ReadAll()
	if cars == nil {
		return []domain.Car{}
	}
	return cars
}
