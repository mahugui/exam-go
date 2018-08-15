package cmodels

import (
	"time"

	"github.com/gomodule/redigo/redis"

	"github.com/goexam/conf"
)

var RedisConn = make(map[string]*redis.Pool)

func Setup() error {
	for key, value := range conf.RedisSetting{
		RedisConn[key] = getNewPool(value)
	}
	return nil
}

func getNewPool(value *conf.Redis) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     value.MaxIdle,
		MaxActive:   value.MaxActive,
		IdleTimeout: value.IdleTimeout,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", value.Host)
			if err != nil {
				return nil, err
			}
			if value.Password != "" {
				if _, err := c.Do("AUTH", value.Password); err != nil {
					c.Close()
					return nil, err
				}
			}
			if value.DB != 0 {
				if _, err := c.Do("SELECT",  value.DB); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}