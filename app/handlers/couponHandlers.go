package handlers

import (
	"coupon/helpers"
	"coupon/models"
	"coupon/services"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetCouponHandlers(rw http.ResponseWriter, r *http.Request) {
	data := models.InputData{}
	//	daraResponse := models.OutputData{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&data); err != nil {
		helpers.SendError(rw, http.StatusUnprocessableEntity)
	} else {
		if data.ValidateStructure() {
			//TODO: procesar y responder
			//helpers.SenData(rw, data, http.StatusOK)
			priceList, minPrice := services.GetPriceList(data.ItemIds)
			fmt.Println(priceList)
			services.CrearMatriz()
			helpers.SenData(rw, minPrice, http.StatusOK)
		} else {
			helpers.SendError(rw, http.StatusUnprocessableEntity)
		}
	}
}
