package testunitario_test

import (
	"coupon/app/config"
	"fmt"
	"github.com/joho/godotenv"
	"net/url"
	"strconv"
	"testing"
)

func TestEnv_Init(t *testing.T) {

	err := godotenv.Load()
	fmt.Println(err)
	env := config.Env{}
	if err := env.Init(); err != nil {
		t.Errorf("Error al cargar las variables de entorno")
	} else {
		t.Log("Variables de entorno cargadas correctamente")
	}
}

func TestEnv_HostPort(t *testing.T) {
	env := config.Env{}
	port := env.HostPort()
	if len(port) > 0 {
		t.Log("el host port esta cargando un valor")
	} else {
		t.Errorf("el host port esta cargando sin valor")
	}
}

func TestEnv_CACHEExpiration(t *testing.T) {
	env := config.Env{}
	minutes := env.CACHEExpiration()
	if _, err := strconv.Atoi(minutes); err == nil {
		t.Log("el cache expiration esta cargando un valor")
	} else {
		t.Errorf("el cache expiration no esta cargando un valor")
	}
}
func TestEnv_CACHEClean(t *testing.T) {
	env := config.Env{}
	hours := env.CACHEClean()
	if _, err := strconv.Atoi(hours); err == nil {
		t.Log("el cache clean esta cargando un valor")
	} else {
		t.Errorf("el cache clean no esta cargando un valor")
	}

}

func TestEnv_MeliAPIRest(t *testing.T) {
	env := config.Env{}
	apiUrl := env.MeliAPIRest()
	if _, err := url.ParseRequestURI(apiUrl); err == nil {
		t.Log("la variable Api Rest esta cargando un valor con el formato correcto")
	} else {
		t.Errorf("la variable Api Rest esta cargando un valor con el formato incorrecto")
	}

}

func TestGetCache(t *testing.T) {

}
