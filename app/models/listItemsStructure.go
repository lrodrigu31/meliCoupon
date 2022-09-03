package models

import "github.com/go-playground/validator"

// ListItems : interface to model the arguments received from the API
type ListItems struct {
	Id      int64    `json:"id"`
	ItemIds []string `json:"item_ids" validate:"required"`
	Amount  float64  `json:"amount" validate:"required"`
}

// ListItemsResponse : interface to answer the service
type ListItemsResponse struct {
	ItemIds []string `json:"item_ids"`
	Total   float64  `json:"total"`
}

// ValidateStructure : validates that the listItem contains the required fields
func (list ListItems) ValidateStructure() bool {
	validate := validator.New()
	if err := validate.Struct(list); err != nil {
		return false
	} else {
		return true
	}
}
