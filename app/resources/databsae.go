package resources

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//TODO: implementar variables de entorno
var dsn = "root:root@tcp(localhost:3306)/items_db?charset=utf8mb4&parseTime=True&loc=Local"

// Database : automatically establishes the connection to the database
var Database = func() (db *gorm.DB) {
	if db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		fmt.Println("Error en la connexion")
		panic(err)
	} else {
		println("Connexion exitosa")
		return db
	}
}()
