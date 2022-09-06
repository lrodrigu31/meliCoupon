package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Database : automatically establishes the connection to the postgres database
var Database = func() (db *gorm.DB) {
	env := Env{}
	env.Init()
	var dsn = fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Shanghai",
		env.DBHost(), env.DBUser(), env.DBPassword(), env.DBName(), env.DBPort(), env.DBSSLmode())

	if db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{}); err != nil {
		fmt.Println("Error en la connexion")
		panic(err)
	} else {
		println("Connexion exitosa")
		return db
	}

}()
