package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

// Env : Struct used to access the environment variables
type Env struct {
}

// Init : func used to init and load the environment variables
func (e Env) Init() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	} else {
		return nil
	}
	return err
}

//Methods to load app server environment variables

func (e Env) HostPort() string {
	return os.Getenv("PORT")
}

// Methods to load cache environment variables

func (e Env) CACHEExpiration() string {
	return os.Getenv("DEFAULT_EXPIRATION")
}
func (e Env) CACHEClean() string {
	return os.Getenv("CLEANUP_INTERVAL")
}

//Methods to load another environment variables

func (e Env) MeliAPIRest() string {
	return os.Getenv("MELI_API_URL")
}
