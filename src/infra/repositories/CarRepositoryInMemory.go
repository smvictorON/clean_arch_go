package repositories

import (
	"clean_arch_go/src/domain"
	"errors"
	"strings"
)

type CarRepositoryInMemory struct {
	cars []domain.Car
}

func NewCarRepositoryInMemory() *CarRepositoryInMemory {
	return &CarRepositoryInMemory{
		cars: []domain.Car{},
	}
}

func (repo *CarRepositoryInMemory) Create(car domain.Car) string {
	repo.cars = append(repo.cars, car)
	return car.Id
}

func (repo *CarRepositoryInMemory) ReadAll() []domain.Car {
	return repo.cars
}

func (repo *CarRepositoryInMemory) ReadByModel(carModel string) []domain.Car {
	var carsWithModel []domain.Car

	for _, car := range repo.cars {
		if strings.EqualFold(car.Model, carModel) {
			carsWithModel = append(carsWithModel, car)
		}
	}

	return carsWithModel
}

func (repo *CarRepositoryInMemory) ReadOne(carId string) (domain.Car, error) {
	for _, car := range repo.cars {
		if car.Id == carId {
			return car, nil
		}
	}
	return domain.Car{}, errors.New("car not found")
}

func (repo *CarRepositoryInMemory) Update(carId string, car domain.Car) bool {
	var carFound *domain.Car
	for i, car := range repo.cars {
		if car.Id == carId {
			carFound = &repo.cars[i]
			break
		}
	}

	if carFound == nil {
		return false
	}

	if car.Year != 0 {
		carFound.Year = car.Year
	}
	if car.Model != "" {
		carFound.Model = car.Model
	}
	if car.Brand != "" {
		carFound.Brand = car.Brand
	}
	if car.Color != "" {
		carFound.Color = car.Color
	}

	return true
}

func (repo *CarRepositoryInMemory) Delete(carId string) bool {
	var newCars []domain.Car

	for _, car := range repo.cars {
		if car.Id != carId {
			newCars = append(newCars, car)
		}
	}

	var res = len(newCars) != len(repo.cars)

	repo.cars = newCars

	return res
}
