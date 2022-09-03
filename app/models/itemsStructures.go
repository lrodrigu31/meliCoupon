package models

import "coupon/resources"

type Item struct {
	Id    string `json:"id" validate:"required"`
	Price int    `json:"price" validate:"required"`
}

type LocalItem struct {
	Id     int64  `json:"id"`
	ItemId string `json:"id" validate:"required"`
	Price  int    `json:"price" validate:"required"`
}

var MigrateLocalItemsStruct = func() error {
	err := resources.Database.AutoMigrate(LocalItem{})
	return err
}()
