package routes

import (
	"coupon/app/handlers"
	mux2 "github.com/gorilla/mux"
)

type couponRoutes struct {
}

//routes : is the func where create the app access coupon routes
func (c couponRoutes) routes(mux *mux2.Router) *mux2.Router {
	// endPoints
	mux.HandleFunc("/coupon/", handlers.GetCouponHandlers).Methods("POST")
	mux.HandleFunc("/coupon", handlers.GetCouponHandlers).Methods("POST")

	return mux
}
