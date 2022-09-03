package handlers

import (
	"coupon/helpers"
	"coupon/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func ListItemsLisHandlers(rw http.ResponseWriter, r *http.Request) {
	listItems := models.ListItems{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&listItems); err != nil {
		helpers.SendError(rw, http.StatusUnprocessableEntity)
	} else {
		if listItems.ValidateStructure() {
			//TODO: procesar y responder
			helpers.SenData(rw, listItems, http.StatusOK)
		} else {
			helpers.SendError(rw, http.StatusUnprocessableEntity)
		}

		//resources.Database.Save(&listItems)
		fmt.Println(models.GetItem("MCO907414010"))
	}
}
