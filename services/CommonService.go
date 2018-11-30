package services

import (
	"github.com/gomodule/redigo/redis"
	"os"
)

type CommonService struct {
	RedisPool *redis.Pool
}

func NewCommonService() *CommonService {
	return &CommonService{
		RedisPool: NewRedisPool(),
	}
}

func NewRedisPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:   20,
		MaxActive: 200,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", "localhost:6379")
			if err != nil {
				return nil, err
			}
			if os.Getenv("REDIS_PASSWD") != "" {
				if _, err := c.Do("AUTH", os.Getenv("REDIS_PASSWD")); err != nil {
					c.Close()
					return nil, err
				}
			}
			if _, err := c.Do("SELECT", "0"); err != nil {
				c.Close()
				return nil, err
			}
			return c, err
		},
	}
}

