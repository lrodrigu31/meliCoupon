package models

type Coupon struct {
	Id      int64    `json:"id"`
	ItemIds []string `json:"item_ids" validate:"required"`
	Amount  float64  `json:"amount" validate:"required"`
}

type CouponResponse struct {
	ItemIds []string `json:"item_ids"`
	Total   float64  `json:"total"`
}

/*
func MigrarUser() {
	db.Database.AutoMigrate(Coupon{})
}*/
