package controllers

import (
	"clean_arch_go/src/applications"
	"clean_arch_go/src/domain"
	"clean_arch_go/src/infra/presenters"
	"encoding/json"
	"net/http"
)

type GetAllCarsCtrl struct {
	carRepo   domain.CarRepository
	presenter presenters.Presenter
}

func NewGetAllCarsCtrl(carRepo domain.CarRepository, presenter presenters.Presenter) *GetAllCarsCtrl {
	return &GetAllCarsCtrl{
		carRepo:   carRepo,
		presenter: presenter,
	}
}

func (ctrl GetAllCarsCtrl) Handle(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	res := applications.NewGetAllCars(ctrl.carRepo).Execute()

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(ctrl.presenter.Format(res)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	return res, nil
}
