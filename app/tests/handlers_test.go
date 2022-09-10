package test

import (
	"bytes"
	"coupon/app/handlers"
	"coupon/app/models"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetRootHandles(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	handlers.Handlers{}.GetWelcome(w, r)
	res := w.Result()
	defer res.Body.Close()

	_, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	if w.Code != 200 {
		t.Errorf("expected error to be 200 got %v", err)
	}
}
func TestGetCouponHandles(t *testing.T) {
	//assert.Equal(t, http.StatusOK, w.Code)
	//assert.Equal(t, []byte("abcd"), w.Body.Bytes())
	itemIds := []string{"MCO516264975", "MCO468600670", "MCO451957965", "MCO612637712", "MCO869692087"}
	amount := 500000.00
	data := models.InputData{ItemIds: itemIds, Amount: amount}
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(data)

	r := httptest.NewRequest(http.MethodGet, "/coupon/", b)
	w := httptest.NewRecorder()
	handlers.Handlers{}.GetCouponHandlers(w, r)
	res := w.Result()
	defer res.Body.Close()

	_, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	if w.Code != 200 {
		t.Errorf("expected error to be 200 got %v", err)
	}
	/*if string(data) != "Bienvenido a MeliCoupon" { //data es _
		t.Errorf("expected Bienvenido a MeliCoupon got %v", string(data))
	}*/
}
func TestGetCouponHandlesTotalZero(t *testing.T) {
	itemIds := []string{"MCO1", "MCO2", "MCO3"}
	amount := 500000.00
	data := models.InputData{ItemIds: itemIds, Amount: amount}
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(data)

	r := httptest.NewRequest(http.MethodGet, "/coupon/", b)
	w := httptest.NewRecorder()
	handlers.Handlers{}.GetCouponHandlers(w, r)
	res := w.Result()
	defer res.Body.Close()

	_, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	if w.Code != 200 {
		t.Errorf("expected error to be 200 got %v", err)
	}
}
func TestGetError404CouponHandles(t *testing.T) {
	//assert.Equal(t, http.StatusOK, w.Code)
	//assert.Equal(t, []byte("abcd"), w.Body.Bytes())
	itemIds := []string{"MCO516264975", "MCO468600670", "MCO451957965", "MCO612637712", "MCO869692087"}
	amount := 500.00
	data := models.InputData{ItemIds: itemIds, Amount: amount}
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(data)

	r := httptest.NewRequest(http.MethodGet, "/coupon/", b)
	w := httptest.NewRecorder()
	handlers.Handlers{}.GetCouponHandlers(w, r)
	res := w.Result()
	defer res.Body.Close()

	_, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	if w.Code != 404 {
		t.Errorf("expected error to be 200 got %v", err)
	}
	/*if string(data) != "Bienvenido a MeliCoupon" { //data es _
		t.Errorf("expected Bienvenido a MeliCoupon got %v", string(data))
	}*/
}
func TestGetError422CouponHandles(t *testing.T) {
	//assert.Equal(t, http.StatusOK, w.Code)
	//assert.Equal(t, []byte("abcd"), w.Body.Bytes())
	itemIds := "MCO516264975"
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(itemIds)

	r := httptest.NewRequest(http.MethodGet, "/coupon/", b)
	w := httptest.NewRecorder()
	handlers.Handlers{}.GetCouponHandlers(w, r)
	res := w.Result()
	defer res.Body.Close()

	_, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	if w.Code != 422 {
		t.Errorf("expected error to be 200 got %v", err)
	}
	/*if string(data) != "Bienvenido a MeliCoupon" { //data es _
		t.Errorf("expected Bienvenido a MeliCoupon got %v", string(data))
	}*/
}
