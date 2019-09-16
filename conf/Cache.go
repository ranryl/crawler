package conf

import "time"

// Cache struct
type Cache struct {
	Master RedisInfo
	Slave  RedisInfo
}

// RedisInfo struct
type RedisInfo struct {
	Network         string
	Address         string
	Password        string
	Db              int
	MaxIdle         int
	MaxActive       int
	IdleTimeout     time.Duration
	MaxConnLifetime time.Duration
	TimeOut         time.Duration
}
