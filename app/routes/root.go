package routes

import (
	"coupon/app/handlers"
	mux2 "github.com/gorilla/mux"
)

type rootRoutes struct {
}

//routes : is the func where create the app access root routes
func (m rootRoutes) routes(mux *mux2.Router) *mux2.Router {
	// endPoints
	mux.HandleFunc("/", handlers.GetWellCome).Methods("GET")
	mux.HandleFunc("", handlers.GetWellCome).Methods("GET")

	return mux
}
