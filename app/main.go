package main

import (
	"coupon/routes"
	"log"
	"net/http"
)

//TODO : implementar variable de entorno
func main() {
	mux := routes.LoadRoutes()
	log.Fatal(http.ListenAndServe(":3000", mux))
}
