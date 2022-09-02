package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dsn = "root:root@tcp(localhost:3306)/goweb_db?charset=utf8mb4&parseTime=True&loc=Local"
var Database = func() (db *gorm.DB) {

	if db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		fmt.Println("Error en la connexion")
		panic(err)
	} else {
		println("Connexion exitosa")
		return db
	}
}()

/*
func getCities() ([]City, error) {
	r, err := http.Get("http://" + addr + "/cities")
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	return decodeCities(r.Body)
}

func decodeCity(r io.Reader) (City, error) {
	city := City{}
	dec := json.NewDecoder(r)
	err := dec.Decode(&city)
	return city, err
}*/
