package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

// Cache struct
type Cache struct {
	client redis.UniversalClient
	ttl    time.Duration
}

const apqPrefix = "apq:"

// NewCache to create new object Cache
func NewCache(ctx context.Context, redisAddress string, password string, ttl time.Duration) (*Cache, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     redisAddress,
		Password: password, // no password set
		DB:       0,        // use default DB
	})

	err := client.Ping(ctx).Err()
	if err != nil {
		return nil, fmt.Errorf("could not create cache: %w", err)
	}

	return &Cache{client: client, ttl: ttl}, nil
}

// Add cache
func (c *Cache) Add(ctx context.Context, key string, value interface{}) {
	c.client.Set(ctx, apqPrefix+key, value, c.ttl)
}

// Get Cache
func (c *Cache) Get(ctx context.Context, key string) (interface{}, bool) {
	s, err := c.client.Get(ctx, apqPrefix+key).Result()
	if err != nil {
		return struct{}{}, false
	}
	return s, true
}
