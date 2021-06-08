package redis_orm

import (
	"github.com/go-redis/redis/v8"
)

func NewDatabase(redisAddress string) (*Database, error){
	if len(redisAddress) == 0{
		return nil, AddressLengthError
	}
	opts, err := redis.ParseURL(redisAddress)
	if err != nil{
		return nil, err
	}

	return &Database{Rdb: redis.NewClient(opts)}, nil
}