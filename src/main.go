package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"clean_arch_go/src/infra/presenters"
	"clean_arch_go/src/infra/repositories"
	"clean_arch_go/src/infra/routers"

	"github.com/gorilla/mux"
)

func main() {
	mux := mux.NewRouter()
	carRepo := repositories.NewCarRepositoryInMemory()
	wipeIdPresenter := presenters.NewWipeIdPresenter()

	routers.CarRouter(
		mux,
		carRepo,
		wipeIdPresenter,
	)

	fmt.Println("Running on", os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), mux))
}
