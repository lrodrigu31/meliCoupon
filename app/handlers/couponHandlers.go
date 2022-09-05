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
			if Output, err := services.ResponseItemServices(data); err == true {
				helpers.SenData(rw, Output, http.StatusOK)
			} else {
				helpers.SendError(rw, http.StatusNotFound)
			}
		} else {
			helpers.SendError(rw, http.StatusUnprocessableEntity)
		}
	}
}
