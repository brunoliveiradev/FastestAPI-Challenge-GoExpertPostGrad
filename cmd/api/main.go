package main

import (
	"github.com/brunoliveiradev/GoExpertPostGrad-Challenge/internal/api/brasilapi"
	"github.com/brunoliveiradev/GoExpertPostGrad-Challenge/internal/api/viacep"
	"github.com/brunoliveiradev/GoExpertPostGrad-Challenge/internal/handler"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	brasilAPIClient := brasilapi.NewClient(http.DefaultClient)
	viaCepClient := viacep.NewClient(http.DefaultClient)

	addressHandler := handler.NewAddressHandler(brasilAPIClient, viaCepClient)

	router.HandleFunc("/api/cep/{cep}", addressHandler.GetCepHandler).Methods("GET")
	http.ListenAndServe(":8080", router)
}
