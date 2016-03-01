package main

import (
	"log"

	"github.com/garyburd/redigo/redis"
)

func newRedisConn(URL string) *redis.PubSubConn {
	log.Printf("Connecting to %s", URL)

	c, err := redis.DialURL(URL)
	if err != nil {
		log.Panic(err)
	}

	return &redis.PubSubConn{c}
}

func watchChannel(subConn *redis.PubSubConn, channel string) {
	subConn.Subscribe(channel)

	for {
		switch v := subConn.Receive().(type) {
		case redis.Message:
			log.Printf("%s: message: %s\n", v.Channel, v.Data)
			err := parseCommand(v.Data)
			if err != nil {
				log.Printf("Command Error: %s", err)
			}
		case redis.Subscription:
			log.Printf("%s: %s %d\n", v.Channel, v.Kind, v.Count)
		case error:
			log.Printf("Error: %s", v)
		}
	}
}
