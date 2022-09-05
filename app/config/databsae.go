package config

import (
	"coupon/app/resources"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Database : automatically establishes the connection to the database
var Database = func() (db *gorm.DB) {
	env := resources.Env{}
	env.Init()

	var dsn = fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		env.DBUser(), env.DBPassword(), env.DBHost(), env.DBPort(), env.DBName(), env.DBCharset())

	fmt.Println(dsn)
	if db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		fmt.Println("Error en la connexion")
		panic(err)
	} else {
		println("Connexion exitosa")
		return db
	}

}()
