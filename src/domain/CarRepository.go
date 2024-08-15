package domain

type CarRepository interface {
	Create(Car) string
	ReadAll() []Car
	ReadByModel(string) []Car
	ReadOne(string) Car
	Update(string, Car) bool
	Delete(string) bool
}
