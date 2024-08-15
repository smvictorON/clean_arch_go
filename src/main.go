package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"clean_arch_go/src/infra/routers"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Err while loading .env: %v", err)
	}

	mux := mux.NewRouter()
	routers.CarRouter(mux)

	fmt.Println("Running on", os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), mux))
}
