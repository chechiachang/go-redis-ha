package main

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis-redis-ha.default.svc.cluster.local:6379",
		Password: "12345678", // no password set
		DB:       0,          // use default DB
	})

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
