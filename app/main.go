package main

import (
	"coupon/router"
	"log"
	"net/http"
)

func main() {
	// cargar rutas
	mux := router.LoadRoutes()
	//iniciar servidor
	log.Fatal(http.ListenAndServe(":3000", mux))
}
