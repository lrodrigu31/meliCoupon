package routes

import (
	"coupon/app/handlers"
	mux2 "github.com/gorilla/mux"
)

func CouponRoutes(mux *mux2.Router) *mux2.Router {
	// endPoints
	mux.HandleFunc("/coupon/", handlers.GetCouponHandlers).Methods("POST")
	mux.HandleFunc("/coupon/", handlers.GetWellCome).Methods("GET")

	return mux
}
