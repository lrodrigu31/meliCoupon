package config

import (
	"coupon/app/resources"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Database : automatically establishes the connection to the database
var Database = func() (db *gorm.DB) {
	env := resources.Env{}
	env.Init()

	dsn := "host=ec2-34-200-205-45.compute-1.amazonaws.com user=wgsytmmqubsnja password=5bee3ffc7ab4b4eb069a0a84a8a06bf2eec132806f75664942ed888dd9dd53f7 dbname=dem9knka5gm9ba port=5432 sslmode=require TimeZone=Asia/Shanghai"
	//db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	/*var dsn = fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		env.DBUser(), env.DBPassword(), env.DBHost(), env.DBPort(), env.DBName(), env.DBCharset())
	 "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	*/
	//dsn2 := fmt.Sprintf(		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",		env.DBHost(), env.DBUser(), env.DBPassword(), env.DBName(), env.DBPort())

	fmt.Println(dsn)
	/*if db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "host=ec2-34-200-205-45.compute-1.amazonaws.com user=wgsytmmqubsnja password=5bee3ffc7ab4b4eb069a0a84a8a06bf2eec132806f75664942ed888dd9dd53f7 dbname=dem9knka5gm9ba port=5432 sslmode=verify-full TimeZone=Asia/Shanghai",
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{}); err != nil {
		fmt.Println("Error en la connexion")
		panic(err)
	} else {
		println("Connexion exitosa")
		return db
	}*/

	if db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{}); err != nil {
		fmt.Println("Error en la connexion")
		panic(err)
	} else {
		println("Connexion exitosa")
		return db
	}

}()
