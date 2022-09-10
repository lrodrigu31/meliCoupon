package main

import (
	"coupon/app/config"
	"coupon/app/repositories"
	"coupon/app/routes"
	_ "coupon/docs"
	"log"
	"net/http"
	"strconv"
	"time"
)

// @title           Mercadolibre coupon API
// @version         1.0
// @description     Returns items that client can buy with a coupon.

// @contact.name   Luis Fernando Rodriguez Llanos
// @contact.url    https://www.linkedin.com/in/lrodriguez031/
// @contact.email  luisrodriguez031@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:3000
// @BasePath  /coupon/

func main() {
	env := config.Env{}
	env.Init()

	//main server
	mux := routes.Routes{}.GetRoutes().StrictSlash(false)
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
