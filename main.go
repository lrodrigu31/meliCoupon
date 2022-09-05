package main

import (
	"coupon/app/routes"
	"log"
	"net/http"
)

//TODO : implementar variable de entorno
func main() {
	/*val := []float64{100, 100, 100, 100, 100}
	wt := []float64{100, 210, 260, 80, 90}
	W := 500.0
	n := len(val)

	println(services.KnapSack(W, wt, val, n))*/

	mux := routes.LoadRoutes()
	log.Fatal(http.ListenAndServe(":3000", mux))
}
