package main

import (
	"coupon/app/config"
	"coupon/app/routes"
	"fmt"
	"gopkg.in/redis.v3"
	"log"
	"net/http"
	"time"
)

//var ctx = context.Background()

func main() {
	env := config.Env{}
	env.Init()
	//TODO :: servidor principal
	mux := routes.Routes{}.GetRoutes()
	srv := &http.Server{
		Handler:      mux,
		Addr:         ":" + env.HostPort(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	//TODO :: servidor redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		//DB:       0,  // use default DB
	})

	pong, err := rdb.Ping().Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(pong, err)

	log.Fatal(srv.ListenAndServe())
}
