package routes

import (
	mux2 "github.com/gorilla/mux"
)

func LoadRoutes() *mux2.Router {

	mux := mux2.NewRouter()

	// implementar rutas
	mux = CouponRoutes(mux)

	return mux
}
