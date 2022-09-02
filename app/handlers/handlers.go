package handlers

import (
	"coupon/helpers"
	"coupon/models"
	"encoding/json"
	"github.com/go-playground/validator"
	"net/http"
)

func CreateCoupon(rw http.ResponseWriter, r *http.Request) {
	coupon := models.Coupon{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&coupon); err != nil {
		helpers.SendError(rw, http.StatusUnprocessableEntity)
	} else {
		validate := validator.New()
		if err := validate.Struct(coupon); err != nil {
			helpers.SendError(rw, http.StatusUnprocessableEntity)
		}
		//db.Database.Save(&coupon)
		helpers.SenData(rw, coupon, http.StatusCreated)
	}
}
