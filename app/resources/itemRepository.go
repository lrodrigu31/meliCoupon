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

//getItem : this func is used to get item from key specific either from cache or from getApiItem if not exist in this
func (item ItemRepository) getItem(itemMeliId string) models.Item {
	var itemResponse = models.Item{}

	if item, sw := GetValue(itemMeliId); !sw {
		itemResponse = getApiItem(itemMeliId)
		//validate response from API
		if itemResponse.ValidateStructure() {

			//register in cache the item found from API
			SetValue(itemMeliId, itemResponse.Price)
		} else {
			//as the item is not get from Api, the func return an object empty
			itemResponse = models.Item{}
		}
	} else {
		itemResponse = item
	}

	//return found item
	return itemResponse
}
