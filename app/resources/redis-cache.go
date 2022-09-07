package resources

import (
	"coupon/app/models"
	json2 "encoding/json"
	"fmt"
	"gopkg.in/redis.v3"
	"time"
)

type redisCache struct {
	host    string
	db      int
	expires time.Duration
}

func NewRedisCache(host string, db int, exp time.Duration) PostCache {

	return &redisCache{
		host:    host,
		db:      db,
		expires: exp,
	}
}
func (cache *redisCache) getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",              // no password set
		DB:       int64(cache.db), // use default DB
	})
}
func (cache *redisCache) Set(key string, value *models.Item) {
	client := cache.getClient()

	json, err := json2.Marshal(value)
	if err != nil {
		panic(err)
	}
	client.Set(key, json, cache.expires*time.Second)
}

func (cache *redisCache) Get(key string) *models.Item {
	client := cache.getClient()
	fmt.Println(":::::::::::::::", client, ":::::::::::::::")
	//val, err := client.Get(key).Result()
	//fmt.Println(":::::::::::::::", err, ":::::::::::::::")
	//if err != nil {
	//	return nil
	//}
	post := models.Item{}
	//err = json2.Unmarshal([]byte(val), &post)
	fmt.Println(":::::::::::::::", post, ":::::::::::::::")
	//if err != nil {
	//	return nil
	//}
	fmt.Println(":::::::::::::::", &post, ":::::::::::::::")
	return &post
}
