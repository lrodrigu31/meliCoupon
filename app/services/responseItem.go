package services

import (
	"coupon/app/models"
)

func ResponseItemServices(data models.InputData) (models.OutputData, bool) {
	priceList, minValue := GetPriceList(data.ItemIds)
	if minValue > data.Amount {
		return models.OutputData{}, false
	}
	Output := models.OutputData{}
	Output.ItemIds, Output.Total = Calculate(priceList.Items, data.Amount)

	return Output, true

}
