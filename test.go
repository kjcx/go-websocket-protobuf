package main

import (
	"github.com/go-redis/redis"
	"fmt"
	"strconv"
	"WebSocket/Mgo"
)

func main() {
	mongo := Mgo.Mongo()
	str := Mgo.GetMsg(mongo,1059)
	fmt.Println(str)
}

func aaa(client *redis.Client,i int){
	err := client.Set(strconv.Itoa(i),"123",0).Err()
	//time.Sleep(time.Microsecond * 10000)
	if err != nil {
		panic(err)
	}
}

func bbb(client *redis.Client,i string){
	data := client.Get(i)
	//if err != nil {
	//	panic(err)
	//}
	fmt.Println(data)

}
