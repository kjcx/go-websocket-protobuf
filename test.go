package main

import (
	"github.com/go-redis/redis"
	"fmt"
	"strconv"
	"time"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       4,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(err,pong)
	var i int
	fmt.Println(time.Now())
	for i=0;i<100000;i++ {
		go insert(client,i)
	}
	fmt.Println(time.Now())
	time.Sleep(10*time.Second)
}
func insert(client *redis.Client,i int){
	client.Set(strconv.Itoa(i),"123",0);
}

