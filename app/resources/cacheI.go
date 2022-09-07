package resources

import (
	"coupon/app/models"
)

type PostCache interface {
	Set(key string, value *models.Item)
	Get(key string) *models.Item
}
