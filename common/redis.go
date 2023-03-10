package common

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"

	"github.com/go-redis/cache/v8"
)

// type Object struct {
// 	Str string
// 	Num int
// }

type CacheService struct {
	cache     *cache.Cache
	context   context.Context
	redisConn *redis.Ring
}

var CacheInstance CacheService

func InitCacheService() {
	config := AppConfig()
	CacheInstance.redisConn = redis.NewRing(&redis.RingOptions{
		Addrs: map[string]string{
			config.RedisHost: ":" + config.RedisPort,
		},
	})

	CacheInstance.cache = cache.New(&cache.Options{
		Redis:      CacheInstance.redisConn,
		LocalCache: cache.NewTinyLFU(1000, time.Minute),
	})

	CacheInstance.context = context.TODO()
}

func (service *CacheService) Set(key string, value string, ttl time.Duration) (err error) {
	err = CacheInstance.cache.Set(&cache.Item{
		Ctx:   CacheInstance.context,
		Key:   key,
		Value: value,
		TTL:   ttl,
	})
	return
}

func (service *CacheService) Get(key string) (wanted string, err error) {
	if err := CacheInstance.cache.Get(service.context, key, &wanted); err != nil {
		return "", err
	}
	return
}
