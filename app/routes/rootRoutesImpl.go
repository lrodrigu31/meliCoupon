package routes

import (
	"coupon/app/handlers"
	_ "coupon/docs"
	mux2 "github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
)

type rootRoutes struct {
}

//routes : is the func where create the app access root routes
func (m rootRoutes) routes(mux *mux2.Router) *mux2.Router {
	var h = handlers.Handlers{}
	// endPoints
	mux.HandleFunc("/", h.GetWelcome).Methods(http.MethodGet)
	mux.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	return mux
}
