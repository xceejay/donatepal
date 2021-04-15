package cache

import (
	"github.com/gomodule/redigo/redis"
)

type RedisCache struct{}

func (rediscache RedisCache) InitCache() (redis.Conn, error) {

	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		return nil, err
	}

	return conn, nil

}
