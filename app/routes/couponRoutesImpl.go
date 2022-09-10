package routes

import (
	"coupon/app/handlers"
	mux2 "github.com/gorilla/mux"
	"net/http"
)

type couponRoutes struct {
}

//routes : is the func where create the app access coupon routes
func (c couponRoutes) routes(mux *mux2.Router) *mux2.Router {

	var h = handlers.Handlers{}
	// endPoints
	mux.HandleFunc("/coupon", h.GetCouponHandlers).Methods(http.MethodPost)
	mux.HandleFunc("/coupon/", h.GetCouponHandlers).Methods(http.MethodPost)

	return mux
}
