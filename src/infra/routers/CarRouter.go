package routers

import (
	"clean_arch_go/src/domain"
	"clean_arch_go/src/infra/controllers"
	"clean_arch_go/src/infra/middlewares"
	"clean_arch_go/src/infra/presenters"
	"net/http"

	"github.com/gorilla/mux"
)

func CarRouter(
	r *mux.Router,
	carRepo domain.CarRepository,
	wipeIdPresenter presenters.Presenter,
) {
	r.Use(middlewares.AuthMiddleware)

	r.HandleFunc("/cars", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			controllers.NewGetAllCarsCtrl(carRepo, wipeIdPresenter).Handle(w, r)
		case http.MethodPost:
			controllers.NewCreateCarCtrl(carRepo).Handle(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}).Methods(http.MethodGet, http.MethodPost)

	r.HandleFunc("/cars/{id:[a-zA-Z0-9_-]+}", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			controllers.NewGetCarByIdCtrl(carRepo).Handle(w, r)
		case http.MethodPatch:
			controllers.NewUpdateCarCtrl(carRepo).Handle(w, r)
		case http.MethodDelete:
			controllers.NewDeleteCarCtrl(carRepo).Handle(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}).Methods(http.MethodGet, http.MethodPatch, http.MethodDelete)

	r.HandleFunc("/cars/model/{model:[a-zA-Z0-9_-]+}", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			controllers.NewGetCarsByModelCtrl(carRepo, wipeIdPresenter).Handle(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}).Methods(http.MethodGet)
}
