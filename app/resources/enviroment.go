package resources

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Env struct {
}

func (e Env) Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func (e Env) HostPort() string {
	return os.Getenv("SERVER_PORT")
}
func (e Env) HostName() string {
	return os.Getenv("SERVER_HOST")
}

func (e Env) DBPort() string {
	return os.Getenv("DB_PORT")
}
func (e Env) DBHost() string {
	return os.Getenv("DB_HOST")
}
func (e Env) DBUser() string {
	return os.Getenv("MYSQL_USER")
}
func (e Env) DBPassword() string {
	return os.Getenv("MYSQL_PASSWORD")
}
func (e Env) DBName() string {
	return os.Getenv("MYSQL_DATABASE")
}
func (e Env) DBCharset() string {
	return os.Getenv("MYSQL_CHARSET")
}

func (e Env) MeliAPIRest() string {
	return os.Getenv("MELI_API_URL")
}
