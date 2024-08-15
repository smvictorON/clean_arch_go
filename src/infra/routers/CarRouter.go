package routers

import (
	"clean_arch_go/src/infra/controllers"
	"net/http"
)

func CarRouter(
	mux *http.ServeMux,
) {
	mux.HandleFunc("/cars", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			controllers.GetCars(w, r)
		case http.MethodPost:
			controllers.CreateCar(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/cars/{id:[a-zA-Z0-9_-]+}", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			controllers.GetCarsById(w, r)
		case http.MethodPatch:
			controllers.CreateCar(w, r)
		case http.MethodDelete:
			controllers.CreateCar(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/cars/model/:model", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			controllers.GetCarsById(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
}
