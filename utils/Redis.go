package utils

import (
	"crawler/conf"
	"encoding/json"
	"errors"
	"sync"
	"time"

	"github.com/gomodule/redigo/redis"
)

var redisPool *redis.Pool
var mu sync.Mutex

// NewRedisPool ...
func NewRedisPool(redisInfo conf.RedisInfo) *CacherClient {
	if redisPool == nil {
		// 双重锁
		mu.Lock()
		defer mu.Unlock()
		if redisPool == nil {
			redisPool = &redis.Pool{
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
	}
	return &CacherClient{
		RedisPool: redisPool,
		prefix:    "",
		marshal:   json.Marshal,
		unmashal:  json.Unmarshal,
	}
}

// CacherClient ...
type CacherClient struct {
	RedisPool *redis.Pool
	prefix    string
	marshal   func(v interface{}) ([]byte, error)
	unmashal  func(data []byte, v interface{}) error
}

// SetPrefix ...
func (rc *CacherClient) SetPrefix(prefix string) {
	rc.prefix = prefix
}

// SetCodingFunc ...
func (rc *CacherClient) SetCodingFunc(
	mar func(v interface{}) ([]byte, error),
	unmar func(data []byte, v interface{}) error) {
	rc.marshal = mar
	rc.unmashal = unmar
}

// Do 直接调用conn Do方法
func (rc *CacherClient) Do(commandName string, args ...interface{}) (result interface{}, err error) {
	conn := rc.RedisPool.Get()
	defer conn.Close()
	return conn.Do(commandName, args...)
}

// Get 直接使用时，需要自己手动反序列化（一般不直接使用）
func (rc *CacherClient) Get(key string) (interface{}, error) {
	return rc.Do("GET", rc.getKey(key))
}

// GetInt func
func (rc *CacherClient) GetInt(key string) (int, error) {
	return redis.Int(rc.Get(rc.getKey(key)))
}

// GetInt64 func
func (rc *CacherClient) GetInt64(key string) (int64, error) {
	return redis.Int64(rc.Get(rc.getKey(key)))
}

// GetString func
func (rc *CacherClient) GetString(key string) (string, error) {
	return redis.String(rc.Get(rc.getKey(key)))
}

// GetBool func
func (rc *CacherClient) GetBool(key string) (bool, error) {
	return redis.Bool(rc.Get(rc.getKey(key)))
}

// GetObject ...
func (rc *CacherClient) GetObject(key string, val interface{}) error {
	reply, err := rc.Get(rc.getKey(key))
	return rc.decode(reply, err, val)
}

// Set string
// val[0]: value
// val[1]: nil | seconds | PX | NX | XX
// val[2]: nil | milliseconds
// val[3]: nil | NX | XX
func (rc *CacherClient) Set(key string, val ...interface{}) error {
	if len(val) <= 0 {
		return errors.New("CacherClient.Set func val param at least have one piece")
	}
	var err error
	// 序列化value值
	val[0], err = rc.encode(val[0])
	if err != nil {
		return nil
	}
	param := rc.paramMerge(rc.getKey(key), val...)
	_, err = rc.Do("SET", param...)
	return err
}

// Del func
func (rc *CacherClient) Del(key string) error {
	_, err := rc.Do("DEL", rc.getKey(key))
	return err
}

// TTL func
func (rc *CacherClient) TTL(key string) (int64, error) {
	return redis.Int64(rc.Do("TTL", rc.getKey(key)))
}

// Expire func
func (rc *CacherClient) Expire(key string, expire int64) error {
	_, err := rc.Do("EXPIRE", rc.getKey(key), expire)
	return err
}

// HSet ...
func (rc *CacherClient) HSet(key string, val ...interface{}) error {
	if len(val) <= 1 {
		return errors.New("CacherClient.Set func val param at least have two pieces")
	}
	var err error
	// 序列化value值
	val[1], err = rc.encode(val[1])
	if err != nil {
		return nil
	}
	value := rc.paramMerge(rc.getKey(key), val...)
	_, err = rc.Do("HSET", value...)
	return err
}

// HGet 直接使用时，需要自己手动反序列化（一般不直接使用）
func (rc *CacherClient) HGet(key string, field string) (result interface{}, err error) {
	key = rc.getKey(key)
	return rc.Do("HGET", key, field)
}

// HGetObject ...
func (rc *CacherClient) HGetObject(key string, field string, val interface{}) error {
	reply, err := rc.HGet(rc.getKey(key), field)
	return rc.decode(reply, err, val)
}

// HGetString ...
func (rc *CacherClient) HGetString(key string, field string) (string, error) {
	return redis.String(rc.HGet(key, field))
}

// HGetInt ...
func (rc *CacherClient) HGetInt(key string, field string) (int, error) {
	return redis.Int(rc.HGet(key, field))
}

// HGetInt64 ...
func (rc *CacherClient) HGetInt64(key string, field string) (int64, error) {
	return redis.Int64(rc.HGet(key, field))
}

// HGetBool ...
func (rc *CacherClient) HGetBool(key string, field string) (bool, error) {
	return redis.Bool(rc.HGet(key, field))
}

// HMSet 必须全部参数正确才能插入
func (rc *CacherClient) HMSet(key string, val ...interface{}) error {
	lens := len(val)
	if lens < 2 {
		return errors.New("CacherClient HMSet func val param at least have two pieces")
	}
	for i := 1; i < lens; i += 2 {
		var err error
		val[i], err = rc.encode(val[i])
		if err != nil {
			return err
		}
	}
	value := rc.paramMerge(rc.getKey(key), val...)
	_, err := rc.Do("HMSET", value...)
	return err
}

// HMGet TODO: 处理返回interface{} 结果 及 反序列化
// func (rc *CacherClient) HMGet(key string, fields ...interface{}) (result interface{}, err error) {
// 	value := rc.paramMerge(rc.getKey(key), fields...)
// 	data, err := rc.Do("HMGET", value...)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return data, err
// }

// LSet ...
func (rc *CacherClient) LSet(key string, val ...interface{}) error {
	var err error
	// 序列化value值
	if len(val) < 2 {
		return errors.New("CacherClient LSet func val param at least have two pieces")
	}
	val[1], err = rc.encode(val[1])
	if err != nil {
		return nil
	}
	val = rc.paramMerge(key, val)
	_, err = rc.Do("LSET", val...)
	return err
}

// LPush ...
func (rc *CacherClient) LPush(key string, val ...interface{}) error {
	return rc.push("LPUSH", key, val...)
}

// LPushX ...
func (rc *CacherClient) LPushX(key string, val ...interface{}) error {
	return rc.push("LPUSHX", key, val...)
}

// RPush ...
func (rc *CacherClient) RPush(key string, val ...interface{}) error {
	return rc.push("RPUSH", key, val...)
}

// RPushX ...
func (rc *CacherClient) RPushX(key string, val ...interface{}) error {
	return rc.push("RPUSHX", key, val...)
}

// push 全部成功或全部失败
func (rc *CacherClient) push(method string, key string, val ...interface{}) error {
	lens := len(val)
	var err error
	for i := 0; i < lens; i++ {
		val[i], err = rc.encode(val[i])
		if err != nil {
			return err
		}
	}
	value := rc.paramMerge(rc.getKey(key), val...)
	_, err = rc.Do(method, value...)
	return err
}

// LPop ...
func (rc *CacherClient) LPop(key string, val interface{}) error {
	return rc.pop("LPOP", key, val)
}

// RPop ...
func (rc *CacherClient) RPop(key string, val interface{}) error {
	return rc.pop("RPOP", key, val)
}

func (rc *CacherClient) pop(method string, key string, val interface{}) error {
	data, err := rc.Do(method, rc.getKey(key))
	return rc.decode(data, err, &val)
}

func (rc *CacherClient) getKey(key string) string {
	return rc.prefix + key
}

// 把key 合并到interface{} 中
func (rc *CacherClient) paramMerge(key string, val ...interface{}) []interface{} {
	value := []interface{}{key}
	value = append(value, val...)
	return value
}

// 编码
func (rc *CacherClient) encode(val interface{}) (interface{}, error) {
	switch v := val.(type) {
	case string, int, uint, int8, int16, int32, int64, float32, float64, bool:
		return v, nil
	default:
		b, err := rc.marshal(val)
		if err != nil {
			return nil, err
		}
		return string(b), nil
	}
}

// 解码
func (rc *CacherClient) decode(reply interface{}, err error, val interface{}) error {
	str, err := redis.String(reply, err)
	if err != nil {
		return nil
	}
	return rc.unmashal([]byte(str), val)
}
