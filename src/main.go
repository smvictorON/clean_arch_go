package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"clean_arch_go/src/infra/repositories"
	"clean_arch_go/src/infra/routers"

	"github.com/gorilla/mux"
)

func main() {
	mux := mux.NewRouter()
	carRepo := repositories.NewCarRepositoryInMemory()

	routers.CarRouter(
		mux,
		carRepo,
	)

	fmt.Println("Running on", os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), mux))
}
