package controllers

import (
	"fmt"
	"net/http"
)

func GetCars(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "GetCars")
}

func CreateCar(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "CreateCar")
}

func UpdateCar(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "UpdateCar")
}

func DeleteCar(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "DeleteCar")
}

func GetCarsByModel(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "GetCarsByModel")
}

func GetCarsById(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "GetCarsById")
}
