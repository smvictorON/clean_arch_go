package controllers

import (
	"clean_arch_go/src/applications"
	"clean_arch_go/src/domain"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type DeleteCarCtrl struct {
	carRepo domain.CarRepository
}

func NewDeleteCarCtrl(carRepo domain.CarRepository) *DeleteCarCtrl {
	return &DeleteCarCtrl{
		carRepo: carRepo,
	}
}

func (ctrl DeleteCarCtrl) Handle(w http.ResponseWriter, r *http.Request) (interface{}, error) {
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
