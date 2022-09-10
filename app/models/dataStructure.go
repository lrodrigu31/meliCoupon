package models

// InputData : structure to model the arguments received from the API godoc
//@Description User's request with the favorite list of items and the coupon total.
type InputData struct {
	ItemIds []string `json:"item_ids" validate:"required" example:"MCO516264975,MCO468600670,MCO451957965,MCO612637712,MCO869692087"`
	Amount  float64  `json:"amount" validate:"required" example:"500000"`
}

// OutputData : structure to answer the service godoc
//@Description Response with the list of items that the user can buy and the total expended on those items.
type OutputData struct {
	ItemIds []string `json:"item_ids" example:"MCO451957965,MCO468600670"`
	Total   float64  `json:"total" example:"470000"`
}
