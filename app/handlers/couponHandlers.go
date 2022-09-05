package handlers

import (
	"coupon/app/helpers"
	"coupon/app/models"
	"coupon/app/services"
	"encoding/json"
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

			if Output, err := services.ResponseItemServices(data); err == true {
				helpers.SenData(rw, Output, http.StatusOK)
			} else {
				helpers.SendError(rw, http.StatusNotFound)
			}

			//priceList, _ := services.GetPriceList(data.ItemIds)
			//response := services.Calculate(priceList.Items, 500)
			//fmt.Println(priceList)
			//services.CrearMatriz()
			//helpers.SenData(rw, response, http.StatusOK)
		} else {
			helpers.SendError(rw, http.StatusUnprocessableEntity)
		}
	}
}
