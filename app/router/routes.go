package router

import (
	"coupon/handlers"
	mux2 "github.com/gorilla/mux"
)

func LoadRoutes() *mux2.Router {

	// rutas
	mux := mux2.NewRouter()

	// endPoints
	mux.HandleFunc("/coupon/", handlers.CreateCoupon).Methods("POST")

	return mux
}
