package main

import (
	"fmt"

	redigo "github.com/garyburd/redigo/redis"
	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/api"
	"github.com/karota-project/gobot-redis"
)

func main() {
	master := gobot.NewGobot()
	api.NewAPI(master).Start()

	redisAdaptor := redis.NewRedisAdaptor("redis-a01", ":6379")
	redisDriver := redis.NewRedisDriver(redisAdaptor, "redis-d01")

	master.AddRobot(
		gobot.NewRobot(
			"redis",
			[]gobot.Connection{redisAdaptor},
			[]gobot.Device{redisDriver},
			func() {
				fmt.Println("work")
				redisDriver.RedisConn().Do("SET", "message1", "hello world")
				value, err := redigo.String(redisDriver.RedisConn().Do("GET", "message1"))

				fmt.Println(value)
				fmt.Println(err)
			}))

	master.Start()
}
