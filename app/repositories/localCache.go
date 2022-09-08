package repositories

import (
	"coupon/app/config"
	"coupon/app/models"
	"github.com/patrickmn/go-cache"
	"time"
)

type LocaCache struct {
}

var env = config.Env{}
var CacheStorage = cache.New(30*time.Minute, 2*time.Hour)

func GetValue(key string) (models.Item, bool) {
	var item = models.Item{}
	val, found := CacheStorage.Get(key)
	if found {
		item.Id = key
		item.Price = val.(float64)
	}
	return item, found
}

func SetValue(key string, value float64) {
	CacheStorage.Set(key, value, cache.DefaultExpiration)
}

func (l LocaCache) NewCacheStorage(defaultExpiration int, cleanupInterval int) {
	CacheStorage = cache.NewFrom(time.Duration(defaultExpiration)*time.Hour, time.Duration(cleanupInterval)*time.Hour, CacheStorage.Items())
}
