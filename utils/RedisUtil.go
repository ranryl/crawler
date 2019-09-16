package utils

import (
	"crawler/conf"
	"time"

	"github.com/gomodule/redigo/redis"
)

// RedisPool ...
var RedisPool *redis.Pool

// NewPool redis pool
func NewPool(redisInfo conf.RedisInfo) {
	RedisPool = &redis.Pool{
		MaxIdle:     redisInfo.MaxIdle,
		MaxActive:   redisInfo.MaxActive,
		IdleTimeout: redisInfo.IdleTimeout,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial(redisInfo.Network, redisInfo.Address,
				redis.DialPassword(redisInfo.Password),
				redis.DialDatabase(redisInfo.Db),
				redis.DialConnectTimeout(redisInfo.TimeOut*time.Second),
				redis.DialReadTimeout(redisInfo.TimeOut*time.Second),
				redis.DialWriteTimeout(redisInfo.TimeOut*time.Second),
			)
			if err != nil {
				return nil, err
			}
			return conn, nil
		},
	}
}
