package main

import (
	"coupon/app/config"
	"coupon/app/resources"
	"coupon/app/routes"
	"log"
	"net/http"
	"strconv"
	"time"
)

func main() {
	env := config.Env{}
	env.Init()
	//TODO :: servidor principal
	mux := routes.Routes{}.GetRoutes()
	srv := &http.Server{
		Handler:      mux,
		Addr:         ":" + env.HostPort(),
		ReadTimeout:  120 * time.Second,
		WriteTimeout: 120 * time.Second,
	}
	//TODO :: Caché
	if defaultExpiration, err := strconv.Atoi(env.CACHEExpiration()); err == nil {
		if cleanupInterval, err := strconv.Atoi(env.CACHEClean()); err == nil {
			resources.LocaCache{}.NewCacheStorage(defaultExpiration, cleanupInterval)
		}
	}

	log.Fatal(srv.ListenAndServe())
}
