package main

import (
	_ "github.com/brunoliveiradev/GoExpertPostGrad-Challenge/docs"
	"github.com/brunoliveiradev/GoExpertPostGrad-Challenge/internal/api/brasilapi"
	"github.com/brunoliveiradev/GoExpertPostGrad-Challenge/internal/api/viacep"
	"github.com/brunoliveiradev/GoExpertPostGrad-Challenge/internal/handler"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	brasilAPIClient := brasilapi.NewClient(http.DefaultClient)
	viaCepClient := viacep.NewClient(http.DefaultClient)

	addressHandler := handler.NewAddressHandler(brasilAPIClient, viaCepClient)

	router.HandleFunc("/api/cep/{cep}", addressHandler.GetCepHandler).Methods("GET")

	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler) // Swagger UI

	http.ListenAndServe(":8080", router)
}
