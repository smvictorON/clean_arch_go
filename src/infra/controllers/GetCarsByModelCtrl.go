package controllers

import (
	"clean_arch_go/src/applications"
	"clean_arch_go/src/domain"
	"clean_arch_go/src/infra/presenters"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type GetCarsByModelCtrl struct {
	carRepo   domain.CarRepository
	presenter presenters.Presenter
}

func NewGetCarsByModelCtrl(carRepo domain.CarRepository, presenter presenters.Presenter) *GetCarsByModelCtrl {
	return &GetCarsByModelCtrl{
		carRepo:   carRepo,
		presenter: presenter,
	}
}

func (ctrl GetCarsByModelCtrl) Handle(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	carModel := vars["model"]

	res := applications.NewGetCarsByModel(ctrl.carRepo).Execute(carModel)

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(ctrl.presenter.Format(res)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil, err
	}

	return res, nil
}
