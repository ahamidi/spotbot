package main

import (
	"flag"
	"log"
)

var redisURL = flag.String("redis", "", "Redis Server URL")
var channel = flag.String("channel", "", "Redis Channel")

func main() {
	log.Println("Spotbot Remote")
	flag.Parse()

	c := newRedisConn(*redisURL)
	defer c.Close()

	watchChannel(c, *channel)
}
