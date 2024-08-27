package controllers

import (
	"clean_arch_go/src/applications"
	"clean_arch_go/src/domain"
	"encoding/json"
	"net/http"
)

type CreateCarCtrl struct {
	carRepo domain.CarRepository
}

func NewCreateCarCtrl(carRepo domain.CarRepository) *CreateCarCtrl {
	return &CreateCarCtrl{
		carRepo: carRepo,
	}
}

func (ctrl CreateCarCtrl) Handle(w http.ResponseWriter, r *http.Request) (interface{}, error) {
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
