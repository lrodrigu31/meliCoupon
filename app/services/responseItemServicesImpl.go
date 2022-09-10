package services

import (
	"coupon/app/models"
)

//ResponseServices : is the struct for services response
type ResponseServices struct {
}

// responseItemServices : is the public method for services response
func (response ResponseServices) responseItemServices(data models.InputData) (models.OutputData, bool) {
	itemService := ItemService{}
	priceList, minValue := itemService.GetPriceList(data.ItemIds)
	if minValue > data.Amount {
		return models.OutputData{}, false
	}
	Output := models.OutputData{}
	Output.ItemIds, Output.Total = itemService.ProcessResponse(priceList.Items, data.Amount)

	return Output, true
}
