package cache

/*
import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

type redisCache struct {
	host string
	db   int
	exp  time.Duration
}

var MenuItemCache *redisCache

func NewRedisCache(host string, db int, exp time.Duration) {
	MenuItemCache = &redisCache{
		host: host,
		db:   db,
		exp:  exp,
	}
}

func (cache *redisCache) GetClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cache.host,
		Password: "",
		DB:       cache.db,
	})

}
func (cache *redisCache) Set(key string, value interface{}) {
	client := cache.GetClient()
	jsonData, err := json.Marshal(value)
	fmt.Printf("\nJSON ===== \n %v", string(jsonData))

	err = client.Set(context.Background(), key, jsonData, cache.exp*time.Minute).Err()
	if err != nil {
		fmt.Printf("\nSet Error := %v\n", err)
	}
}

func (cache *redisCache) Get(key string) (string, error) {
	client := cache.GetClient()
	val, err := client.Get(context.Background(), key).Result()
	if err != nil && err == redis.Nil {
		return "", err
	}
	return val, nil
}

func (cache *redisCache) Del(key string) error {
	client := cache.GetClient()
	err := client.Del(context.Background(), key).Err()
	if err != nil {
		return err
	}
	return nil
}
*/
