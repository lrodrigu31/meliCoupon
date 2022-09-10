package handlers

import "net/http"

type IHandler interface {
	getWelcome(rw http.ResponseWriter, r *http.Request)
	getCouponHandlers(rw http.ResponseWriter, r *http.Request)
}

func (h Handlers) GetWelcome(rw http.ResponseWriter, r *http.Request) {
	h.getWelcome(rw, r)
}
func (h Handlers) GetCouponHandlers(rw http.ResponseWriter, r *http.Request) {
	h.getCouponHandlers(rw, r)
}
