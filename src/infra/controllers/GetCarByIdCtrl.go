package controllers

import (
	"clean_arch_go/src/applications"
	"clean_arch_go/src/domain"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type GetCarByIdCtrl struct {
	carRepo domain.CarRepository
}

func NewGetCarByIdCtrl(carRepo domain.CarRepository) *GetCarByIdCtrl {
	return &GetCarByIdCtrl{
		carRepo: carRepo,
	}
}

func (ctrl GetCarByIdCtrl) Handle(w http.ResponseWriter, r *http.Request) (interface{}, error) {
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
