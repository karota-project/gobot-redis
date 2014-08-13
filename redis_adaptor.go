package redis

import (
	"github.com/garyburd/redigo/redis"
	"github.com/hybridgroup/gobot"
)

type RedisAdaptor struct {
	gobot.Adaptor
	redisConn redis.Conn
}

func NewRedisAdaptor(name string, port string) *RedisAdaptor {
	return &RedisAdaptor{
		Adaptor: *gobot.NewAdaptor(
			name,
			"redis.RedisAdaptor",
			port,
		),
	}
}

func (r *RedisAdaptor) Connect() bool {
	if r.Connected() == true {
		disconnect(r)
	}
	connect(r)
	return true
}

func (r *RedisAdaptor) Finalize() bool {
	disconnect(r)
	return true
}

func connect(r *RedisAdaptor) {
	conn, err := redis.Dial("tcp", r.Port())
	if err != nil {
		panic(err)
	}

	r.redisConn = conn
	r.SetConnected(true)
}

func disconnect(r *RedisAdaptor) {
	if r.redisConn != nil {
		r.redisConn.Close()
	}
	r.SetConnected(false)
}
