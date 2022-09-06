package resources

import (
	"coupon/app/config"
	"coupon/app/models"
	"encoding/json"
	"fmt"
	"net/http"
)

type ItemRepository struct {
}

// getItem : func used to get item  from key specific, implementing getDatabaseItem func
func (item ItemRepository) getItem(itemMeliId string) models.Item {
	itemResponse := getDatabaseItem(itemMeliId)
	return itemResponse
}

//getApiItem : MELI REST API implementation
func getApiItem(itemMeliId string) models.Item {
	env := config.Env{}
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

//getDatabaseItem : this func get the item from database or from getApiItem if not exist in this
func getDatabaseItem(itemMeliId string) models.Item {
	itemResponse := models.Item{}
	localItem := models.LocalItem{}

	// create the item structure the first time implemented
	localItem.MigrateStrut()

	// search an item in database
	if err := config.Database.First(&localItem, "item_id = ?", itemMeliId); err.Error != nil {
		//as the item is not exist in the database, is get from Api
		itemResponse = getApiItem(itemMeliId)
		//validate response from API
		if itemResponse.ValidateStructure() {
			//register in database the item found from API
			localItem.ItemId = itemResponse.Id
			localItem.Price = itemResponse.Price
			config.Database.Save(&localItem)
		} else {
			//as the item is not get from Api, the func return an object empty
			itemResponse = models.Item{}
		}
	} else {
		//assign the results found from database
		itemResponse.Id = localItem.ItemId
		itemResponse.Price = localItem.Price
	}

	//return found item
	return itemResponse
}
