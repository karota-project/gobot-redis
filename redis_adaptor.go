package redis

import (
	"github.com/hybridgroup/gobot"
)

type RedisAdaptor struct {
	gobot.Adaptor
}

func NewRedisAdaptor(name string) *RedisAdaptor {
	return &RedisAdaptor{
		Adaptor: *gobot.NewAdaptor(
			name,
			"redis.RedisAdaptor",
		),
	}
}

func (r *RedisAdaptor) Connect() bool {
	return true
}

func (r *RedisAdaptor) Finalize() bool {
	return true
}
