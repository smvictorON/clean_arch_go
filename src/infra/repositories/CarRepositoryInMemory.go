package repositories

import (
	"clean_arch_go/src/domain"
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

func (repo *CarRepositoryInMemory) ReadByModel(carId string) []domain.Car {

	return []domain.Car{}
}

func (repo *CarRepositoryInMemory) ReadOne(carId string) domain.Car {

	return domain.Car{}
}

func (repo *CarRepositoryInMemory) Update(carId string, car domain.Car) bool {

	return true
}

func (repo *CarRepositoryInMemory) Delete(carId string) bool {

	return true
}
