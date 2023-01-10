package RedisDB

import (
	"github.com/gomodule/redigo/redis"
)

var pool1 *redis.Pool

func InitRdb() {
	pool1 = &redis.Pool{
		MaxIdle: 50,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "127.0.0.1:6379")
		},
	}
}

func GetRedisConn() redis.Conn {

	return pool1.Get()
}
