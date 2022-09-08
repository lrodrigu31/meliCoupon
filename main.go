package main

import (
	"coupon/app/config"
	"coupon/app/repositories"
	"coupon/app/routes"
	"log"
	"net/http"
	"strconv"
	"time"
)

func main() {
	env := config.Env{}
	env.Init()

	//main server
	mux := routes.Routes{}.GetRoutes()
	srv := &http.Server{
		Handler:      mux,
		Addr:         ":" + env.HostPort(),
		ReadTimeout:  120 * time.Second,
		WriteTimeout: 120 * time.Second,
	}
	// initialization Caché
	if defaultExpiration, err := strconv.Atoi(env.CACHEExpiration()); err == nil {
		if cleanupInterval, err := strconv.Atoi(env.CACHEClean()); err == nil {
			repositories.LocaCache{}.NewCacheStorage(defaultExpiration, cleanupInterval)
		}
	}

	log.Fatal(srv.ListenAndServe())
}
