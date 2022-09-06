package routes

import (
	mux2 "github.com/gorilla/mux"
)

type Routes struct {
}

//GetRoutes : is the public method for access to route creates in getRoutes
func (routes Routes) GetRoutes() *mux2.Router {
	return getRoutes()
}

//api : is the interface from routes package
type api interface {
	routes(mux *mux2.Router) *mux2.Router
}

//getRoutes : is the private function to implement every API routes
func getRoutes() *mux2.Router {
	mux := mux2.NewRouter()
	mux = couponRoutes{}.routes(mux)
	mux = rootRoutes{}.routes(mux)
	return mux
}
