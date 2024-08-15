package controllers

import (
	"clean_arch_go/src/applications"
	"clean_arch_go/src/domain"
	"encoding/json"
	"fmt"
	"net/http"
)

type Controller struct {
	carRepo domain.CarRepository
}

func NewController(carRepo domain.CarRepository) *Controller {
	return &Controller{
		carRepo: carRepo,
	}
}

func (ctrl Controller) CreateCar(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	var car domain.Car

	if err := json.NewDecoder(r.Body).Decode(&car); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return nil, err
	}

	res, err := applications.NewCreateCar(ctrl.carRepo).Execute(car)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	return res, nil
}

func (ctrl Controller) GetCars(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	res, err := applications.NewGetAllCars(ctrl.carRepo).Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	return res, nil
}

func (ctrl Controller) UpdateCar(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	fmt.Fprintln(w, "UpdateCar")
	return nil, nil
}

func (ctrl Controller) DeleteCar(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	fmt.Fprintln(w, "DeleteCar")
	return nil, nil
}

func (ctrl Controller) GetCarsByModel(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	fmt.Fprintln(w, "GetCarsByModel")
	return nil, nil
}

func (ctrl Controller) GetCarsById(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	fmt.Fprintln(w, "GetCarsById")
	return nil, nil
}
