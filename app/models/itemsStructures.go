package models

import (
	"github.com/go-playground/validator"
)

// Item : defines struct for use the items in the app
type Item struct {
	Id    string  `json:"id" validate:"required"`
	Price float64 `json:"price" validate:"required"`
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
