package main

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

const (
	Addr     = "redis-redis-ha.default.svc.cluster.local:26379"
	Password = ""
)

func main() {
	// Set & Get
	setClient := newClient()
	getClient := newClient()
	for {

		if err := setClient.Set("ticker", time.Now().Second(), 0).Err(); err != nil {
			fmt.Println(err.Error())
		}

		val, err := getClient.Get("ticker").Result()
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println("Got: ", val)
		time.Sleep(1 * time.Second)

	}

	// Recieve
	subClient := newClient()
	pubClient := newClient()
	topic := "pticker"

	sub := subClient.PSubscribe(topic)
	defer sub.Close()
	go func() {
		for {
			msg, err := sub.ReceiveMessage()
			if err != nil {
				fmt.Println(err.Error())
				//panic(err)
			}
			fmt.Printf("recieved %s", msg)
		}
	}()

	// Publish
	go func() {
		for {
			if err := pubClient.Publish(topic, time.Now().Second()).Err(); err != nil {
				fmt.Println(err.Error())
				//panic(err)
			}
			time.Sleep(1 * time.Second)
		}
	}()
}

func newClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     Addr,     // connect to sentinel, not directly to redis
		Password: Password, // no password set
		DB:       0,        // use default DB
	})
}
