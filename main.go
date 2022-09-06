package main

import (
	"context"
	"coupon/app/config"
	"coupon/app/routes"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"net/http"
	"time"
)

var ctx = context.Background()

func main() {
	env := config.Env{}
	env.Init()
	mux := routes.Routes{}.GetRoutes()
	srv := &http.Server{
		Handler:      mux,
		Addr:         ":" + env.HostPort(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		//Password: "", // no password set
		//DB:       0,  // use default DB
	})

	pong, err := rdb.Ping(ctx).Result()
	fmt.Println(pong, err)

	log.Fatal(srv.ListenAndServe())
}
