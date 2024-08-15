package routers

import (
	"clean_arch_go/src/infra/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func CarRouter(
	r *mux.Router,
) {
	r.HandleFunc("/cars", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			controllers.GetCars(w, r)
		case http.MethodPost:
			controllers.CreateCar(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}).Methods(http.MethodGet, http.MethodPost)

	r.HandleFunc("/cars/{id:[a-zA-Z0-9_-]+}", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			controllers.GetCarsById(w, r)
		case http.MethodPatch:
			controllers.UpdateCar(w, r)
		case http.MethodDelete:
			controllers.DeleteCar(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}).Methods(http.MethodGet, http.MethodPatch, http.MethodDelete)

	r.HandleFunc("/cars/model/{model:[a-zA-Z0-9_-]+}", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			controllers.GetCarsByModel(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}).Methods(http.MethodGet)
}
