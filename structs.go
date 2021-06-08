package redis_orm

import "github.com/go-redis/redis/v8"

type Database struct {
	Rdb *redis.Client
}