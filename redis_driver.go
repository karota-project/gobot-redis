package redis

import (
	"github.com/hybridgroup/gobot"
)

type RedisDriver struct {
	gobot.Driver
}

type RedisInterface interface {
}

func NewRedisDriver(a *RedisAdaptor, name string) *RedisDriver {
	return &RedisDriver{
		Driver: *gobot.NewDriver(
			name,
			"redis.RedisDriver",
			a,
		),
	}
}

func (r *RedisDriver) adaptor() *RedisAdaptor {
	return r.Driver.Adaptor().(*RedisAdaptor)
}

func (r *RedisDriver) Start() bool { return true }
func (r *RedisDriver) Halt() bool  { return true }
