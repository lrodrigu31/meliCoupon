package repository

import (
	"coupon/app/config"
	"coupon/app/models"
	"coupon/app/resources"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetItem(itemMeliId string) models.Item {
	itemResponse := getDatabaseItem(itemMeliId)
	return itemResponse
}

//getApiItem : MELI REST API implementation
func getApiItem(itemMeliId string) models.Item {
	env := resources.Env{}
	env.Init()

	itemResponse := models.Item{}
	if response, err := http.Get(env.MeliAPIRest() + itemMeliId); err != nil {
		fmt.Println(err.Error())
	} else {
		decoder := json.NewDecoder(response.Body)
		if err := decoder.Decode(&itemResponse); err != nil {
			fmt.Println(err)
		}
	}
	return itemResponse
}

func getDatabaseItem(itemMeliId string) models.Item {
	itemResponse := models.Item{}
	localItem := models.LocalItem{}
	localItem.MigrateStrut()
	if err := config.Database.First(&localItem, "item_id = ?", itemMeliId); err.Error != nil {
		itemResponse = getApiItem(itemMeliId)
		if itemResponse.ValidateStructure() {
			localItem.ItemId = itemResponse.Id
			localItem.Price = itemResponse.Price
			config.Database.Save(&localItem)
		} else {
			itemResponse = models.Item{}
		}
	} else {
		itemResponse.Id = localItem.ItemId
		itemResponse.Price = localItem.Price
	}
	return itemResponse
}
