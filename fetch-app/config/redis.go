package config

import (
	"os"

	"github.com/gomodule/redigo/redis"
)

var Redis redis.Conn

func InitRedis() (*redis.Pool, error) {
	pool := redis.NewPool(
		func() (redis.Conn, error) {
			return redis.Dial("tcp", os.Getenv("REDIS_HOST")+":6379")
		},
		20,
	)
	pool.MaxActive = 20

	return pool, nil
}
