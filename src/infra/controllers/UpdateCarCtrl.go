package controllers

import (
	"clean_arch_go/src/applications"
	"clean_arch_go/src/domain"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type UpdateCarCtrl struct {
	carRepo domain.CarRepository
}

func NewUpdateCarCtrl(carRepo domain.CarRepository) *UpdateCarCtrl {
	return &UpdateCarCtrl{
		carRepo: carRepo,
	}
}

func (ctrl UpdateCarCtrl) Handle(w http.ResponseWriter, r *http.Request) (interface{}, error) {
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
