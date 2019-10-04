package conf

import "time"

// Cache struct
type Cache struct {
	Master RedisInfo `yaml:"master"`
	Slave  RedisInfo `yaml:"slave"`
}

// RedisInfo struct
type RedisInfo struct {
	Network         string        `yaml:"network"`
	Address         string        `yaml:"address"`
	Password        string        `yaml:"password"`
	Db              int           `yaml:"db"`
	MaxIdle         int           `yaml:"maxidle"`
	MaxActive       int           `yaml:"maxactive"`
	IdleTimeout     time.Duration `yaml:"idletimeout"`
	MaxConnLifetime time.Duration `yaml:"maxconnlifetime"`
	TimeOut         time.Duration `yaml:"timeout"`
}
