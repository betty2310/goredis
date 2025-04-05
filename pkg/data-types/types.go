package datatypes

import (
	"time"
)

type IRedisValue interface {
	Get() string
}

type RedisValue struct {
	Value IRedisValue
	Exp   time.Time
}

type RedisString string

func (c RedisString) Get() string {
	return string(c)
}
