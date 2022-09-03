package services

import (
	"coupon/models"
	"coupon/repository"
	"fmt"
)

func Calculate(items map[string]float64, amount float64) []string {

	for i, item := range items {
		fmt.Println("i:", i)
		fmt.Println("items:", item)
	}

	list := []string{"MLA1", "MLA2", "MLA4", "MLA5"}
	return list
}

func GetPriceList(itemIds []string) (models.PriceList, float64) {
	priceList := make(map[string]float64)
	curPrice := 0.0
	for _, itemId := range itemIds {
		item := repository.GetItem(itemId)
		if _, present := priceList[item.Id]; !present {
			if item.ValidateStructure() {
				minPrice(&curPrice, &item.Price)
				priceList[item.Id] = item.Price
			}
		}
	}
	listMapped := models.PriceList{}
	listMapped.Items = priceList
	return listMapped, curPrice
}

func minPrice(price *float64, priceToCompare *float64) {
	if *price == 0.0 {
		*price = *priceToCompare
	} else if *price > *priceToCompare {
		*price = *priceToCompare
	}
}
