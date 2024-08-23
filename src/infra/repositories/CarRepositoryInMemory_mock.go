package repositories

import (
	"clean_arch_go/src/domain"

	"github.com/stretchr/testify/mock"
)

type CarRepositoryInMemoryMock struct {
	mock.Mock
}

func NewCarRepositoryInMemoryMock() *CarRepositoryInMemoryMock {
	return &CarRepositoryInMemoryMock{}
}

func (repo *CarRepositoryInMemoryMock) Create(car domain.Car) string {
	args := repo.Called(car)
	return args.String(0)
}

func (repo *CarRepositoryInMemoryMock) ReadAll() []domain.Car {
	args := repo.Called()
	return args.Get(0).([]domain.Car)
}

func (repo *CarRepositoryInMemoryMock) ReadByModel(carModel string) []domain.Car {
	args := repo.Called()
	return args.Get(0).([]domain.Car)
}

func (repo *CarRepositoryInMemoryMock) ReadOne(carId string) (*domain.Car, error) {
	args := repo.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Car), args.Error(1)
}

func (repo *CarRepositoryInMemoryMock) Update(carId string, car domain.Car) bool {
	args := repo.Called()
	return args.Bool(0)
}

func (repo *CarRepositoryInMemoryMock) Delete(carId string) bool {
	args := repo.Called()
	return args.Bool(0)
}
