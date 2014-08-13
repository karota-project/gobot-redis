package redis

import (
	"github.com/hybridgroup/gobot"
	"testing"
)

func initTestRedisDriver() *RedisDriver {
	return NewRedisDriver(NewRedisAdaptor("myAdaptor"), "myDriver")
}

func TestRedisDriverStart(t *testing.T) {
	d := initTestRedisDriver()
	gobot.Expect(t, d.Start(), true)
}

func TestRedisDriverHalt(t *testing.T) {
	d := initTestRedisDriver()
	gobot.Expect(t, d.Halt(), true)
}
