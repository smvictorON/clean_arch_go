package domain

type CarRepository interface {
	Create(Car) string
	ReadAll() []Car
	ReadByModel(string) []Car
	ReadOne(string) (*Car, error)
	Update(string, Car) bool
	Delete(string) bool
}
