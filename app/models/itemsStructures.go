package models

import (
	"coupon/config"
	"github.com/go-playground/validator"
)

type Item struct {
	Id    string  `json:"id" validate:"required"`
	Price float64 `json:"price" validate:"required"`
}

type LocalItem struct {
	Id     int64   `json:"id"`
	ItemId string  `json:"item_id" validate:"required"`
	Price  float64 `json:"price" validate:"required"`
}

type PriceList struct {
	Items map[string]float64
}

type Prices struct {
}

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
