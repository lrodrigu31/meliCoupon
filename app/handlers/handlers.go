package handlers

import (
	"coupon/app/helpers"
	"coupon/app/models"
	"coupon/app/services"
	"encoding/json"
	"net/http"
)

type Handlers struct {
}

// GetCouponHandlers godoc
//@Consume json
//@Produce json
//@Param request body models.InputData true "Request body example"
//@Summary Returns a list of items from user's fav item list that can buy with a coupon
//@Description This API returns a list of items that user can buy given the user's fav item list and a coupon.
//@Description The total of item list cannot exceed the amount of the coupon
//@Success 200 {object} models.OutputData
//failure 404 {string} "Resource not found"
//@Router / [post]
func (h Handlers) getCouponHandlers(rw http.ResponseWriter, r *http.Request) {
	data := models.InputData{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&data); err != nil {
		helpers.SendError(rw, http.StatusUnprocessableEntity, "the body struct invalid")
	} else {
		response := services.ResponseServices{}
		if Output, err := response.ResponseItemServices(data); err == true {
			helpers.SenData(rw, Output, http.StatusOK)
		} else {
			helpers.SendError(rw, http.StatusNotFound, "Resource not found")
		}
	}
}

// GetWelcome is a handler that return the message "Bienvenido a MeliCoupon"
func (h Handlers) getWelcome(rw http.ResponseWriter, r *http.Request) {
	helpers.SenData(rw, "Bienvenido a MeliCoupon", http.StatusOK)
}
