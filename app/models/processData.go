package models

import (
	"coupon/resources"
	"encoding/json"
	"fmt"
	"net/http"
)

//getRemoteItem : MELI REST API implementation TODO: implementar variable de entorno
func getRemoteItem(itemMeliId string) Item {
	itemResponse := Item{}
	if response, err := http.Get("https://api.mercadolibre.com/items/" + itemMeliId); err != nil {
		fmt.Println(err.Error())
	} else {
		decoder := json.NewDecoder(response.Body)
		if err := decoder.Decode(&itemResponse); err != nil {
			fmt.Println(err)
		}
	}
	return itemResponse
}

func GetItem(itemMeliId string) Item {
	itemResponse := getDatabaseItem(itemMeliId)
	return itemResponse
}

func getDatabaseItem(itemMeliId string) Item {
	itemResponse := Item{}
	localItem := LocalItem{}
	if err := resources.Database.First(&localItem, "item_id = ?", itemMeliId); err.Error != nil {
		itemResponse = getRemoteItem(itemMeliId)
		localItem.ItemId = itemResponse.Id
		localItem.Price = itemResponse.Price
		resources.Database.Save(&localItem)
	} else {
		itemResponse.Id = localItem.ItemId
		itemResponse.Price = localItem.Price
	}
	return itemResponse
}
