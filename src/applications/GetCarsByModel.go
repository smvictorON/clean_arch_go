package applications

import (
	"clean_arch_go/src/domain"
)

type GetCarsByModel struct {
	carRepo domain.CarRepository
}

func NewGetCarsByModel(carRepo domain.CarRepository) *GetCarsByModel {
	return &GetCarsByModel{
		carRepo: carRepo,
	}
}

func (app *GetCarsByModel) Execute(carModel string) []domain.Car {
	cars := app.carRepo.ReadByModel(carModel)
	if cars == nil {
		return []domain.Car{}
	}
	return cars
}
