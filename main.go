package main

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis-redis-ha.default.svc.cluster.local:26379", // connect to sentinel, not directly to redis
		Password: "",                                               // no password set
		DB:       0,                                                // use default DB
	})

	// Recieve
	pubsub := client.PSubscribe("pticker")
	defer pubsub.Close()
	go func() {
		for {
			msg, err := pubsub.ReceiveMessage()
			if err != nil {
				panic(err)
			}
			fmt.Printf("recieved %s", msg)
		}
	}()

	// Publish
	go func() {
		for {
			if err := client.Publish("pticker", time.Now().Second()).Err(); err != nil {
				panic(err)
			}
			time.Sleep(1 * time.Second)
		}
	}()

	// Set & Get
	for {

		err := client.Set("ticker", time.Now().Second(), 0).Err()
		if err != nil {
			fmt.Println(err.Error())
		}

		val, err := client.Get("ticker").Result()
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println("Got: ", val)
		time.Sleep(1 * time.Second)

	}
}
