package main

import (
	"coupon/app/resources"
	"coupon/app/routes"
	"log"
	"net/http"
	"time"
)

func main() {
	env := resources.Env{}
	env.Init()
	mux := routes.LoadRoutes()

	srv := &http.Server{
		Handler:      mux,
		Addr:         ":" + env.HostPort(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
