package redis

import (
  "github.com/hybridgroup/gobot"
  "testing"
)

func initTestRedisAdaptor() *RedisAdaptor {
  return NewRedisAdaptor("myAdaptor")
}

func TestRedisAdaptorConnect(t *testing.T) {
  a := initTestRedisAdaptor()
  gobot.Expect(t, a.Connect(), true)
}

func TestRedisAdaptorFinalize(t *testing.T) {
  a := initTestRedisAdaptor()
  gobot.Expect(t, a.Finalize(), true)
}
