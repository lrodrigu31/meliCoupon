package main

import (
	"coupon/app/resources"
	"coupon/app/routes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	env := resources.Env{}
	env.Init()

	fmt.Println("::::::::::", env.HostPort(), "::::::::::")
	mux := routes.LoadRoutes()
	log.Fatal(http.ListenAndServe(env.HostName()+":"+env.HostPort(), mux))
}
