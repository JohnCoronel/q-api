package storage

import "github.com/go-redis/redis"

type cache struct {
	cache *redis.Client
}
