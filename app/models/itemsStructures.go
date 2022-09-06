package models

import (
	"coupon/app/config"
	"github.com/go-playground/validator"
)

// Item : defines struct for use the items in the app
type Item struct {
	Id    string  `json:"id" validate:"required"`
	Price float64 `json:"price" validate:"required"`
}

// LocalItem : defines struct for registers from database
type LocalItem struct {
	Id     int64   `json:"id"`
	ItemId string  `json:"item_id" validate:"required"`
	Price  float64 `json:"price" validate:"required"`
}

//MigrateStrut this is implement func for to migrate struct LocalItem
func (localItem LocalItem) MigrateStrut() {
	config.MigrateStrut(localItem)
}

// ValidateStructure : validates that the InputData contains the required fields
func (data Item) ValidateStructure() bool {
	validate := validator.New()
	if err := validate.Struct(data); err != nil {
		return false
	} else {
		return true
	}
}
