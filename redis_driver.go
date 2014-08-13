package redis

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/hybridgroup/gobot"
)

type RedisDriver struct {
	gobot.Driver
}

type RedisInterface interface {
}

func NewRedisDriver(a *RedisAdaptor, name string) *RedisDriver {
	r := &RedisDriver{
		Driver: *gobot.NewDriver(
			name,
			"redis.RedisDriver",
			a,
		),
	}

	r.AddCommand("Get", func(params map[string]interface{}) interface{} {
		key := params["key"].(string)

		value, err := redis.String(r.adaptor().redisConn.Do("GET", key))
		if err == nil {
			fmt.Println(value)
		}

		return true
	})

	r.AddCommand("Set", func(params map[string]interface{}) interface{} {
		key := params["key"].(string)
		value := params["value"].(string)

		r.adaptor().redisConn.Do("SET", key, value)

		return true
	})

	return r
}

func (r *RedisDriver) adaptor() *RedisAdaptor {
	return r.Driver.Adaptor().(*RedisAdaptor)
}

func (r *RedisDriver) Start() bool {
	return true
}

func (r *RedisDriver) Halt() bool {
	return true
}

func (r *RedisDriver) RedisConn() redis.Conn {
	return r.adaptor().redisConn
}
