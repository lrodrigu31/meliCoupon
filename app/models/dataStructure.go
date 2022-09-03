package models

import "github.com/go-playground/validator"

// InputData : structure to model the arguments received from the API
type InputData struct {
	ItemIds []string `json:"item_ids" validate:"required"`
	Amount  float64  `json:"amount" validate:"required"`
}

// OutputData : structure to answer the service
type OutputData struct {
	ItemIds []string `json:"item_ids"`
	Total   float64  `json:"total"`
}

// ValidateStructure : validates that the InputData contains the required fields
func (data InputData) ValidateStructure() bool {
	validate := validator.New()
	if err := validate.Struct(data); err != nil {
		return false
	} else {
		return true
	}
}
