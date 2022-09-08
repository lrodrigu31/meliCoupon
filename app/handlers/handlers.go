package handlers

import (
	"coupon/app/helpers"
	"coupon/app/models"
	"coupon/app/services"
	"encoding/json"
	"net/http"
)

// GetCouponHandlers : coupon api request handler, this get the list of items that maximizes the total to spend without exceeding it.
func GetCouponHandlers(rw http.ResponseWriter, r *http.Request) {
	data := models.InputData{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&data); err != nil {
		helpers.SendError(rw, http.StatusUnprocessableEntity, "the body struct invalid")
	} else {
		if data.ValidateStructure() {
			response := services.ResponseServices{}
			if Output, err := response.ResponseItemServices(data); err == true {
				helpers.SenData(rw, Output, http.StatusOK)
			} else {
				helpers.SendError(rw, http.StatusNotFound, "Resource not found")
			}
		} else {
			helpers.SendError(rw, http.StatusUnprocessableEntity, "the body struct invalid")
		}
	}
}

// GetWelcome is a handler that return the message "Bienvenido  a MeliCoupon"
func GetWelcome(rw http.ResponseWriter, r *http.Request) {
	helpers.SenData(rw, "Bienvenido  a MeliCoupon", http.StatusOK)
}
