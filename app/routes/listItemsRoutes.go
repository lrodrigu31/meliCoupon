package routes

import (
	"coupon/handlers"
	mux2 "github.com/gorilla/mux"
)

func ListItemsRoutes(mux *mux2.Router) *mux2.Router {
	// endPoints
	mux.HandleFunc("/coupon/", handlers.ListItemsLisHandlers).Methods("POST")

	return mux
}
