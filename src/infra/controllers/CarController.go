package controllers

import (
	"clean_arch_go/src/applications"
	"clean_arch_go/src/domain"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
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

	res := applications.NewCreateCar(ctrl.carRepo).Execute(car)

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	return res, nil
}

func (ctrl Controller) GetCars(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	res := applications.NewGetAllCars(ctrl.carRepo).Execute()

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	return res, nil
}

func (ctrl Controller) UpdateCar(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	carId := vars["id"]

	var car domain.Car
	if err := json.NewDecoder(r.Body).Decode(&car); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return nil, err
	}

	res := applications.NewUpdateCar(ctrl.carRepo).Execute(carId, car)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil, err
	}

	return res, nil
}

func (ctrl Controller) DeleteCar(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	carId := vars["id"]

	res := applications.NewDeleteCar(ctrl.carRepo).Execute(carId)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil, err
	}

	return res, nil
}

func (ctrl Controller) GetCarsByModel(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	carModel := vars["model"]

	res := applications.NewGetCarsByModel(ctrl.carRepo).Execute(carModel)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil, err
	}

	return res, nil
}

func (ctrl Controller) GetCarsById(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	carId := vars["id"]

	res, err := applications.NewGetCarById(ctrl.carRepo).Execute(carId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil, err
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil, err
	}

	return res, nil
}
