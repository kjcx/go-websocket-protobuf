package main

import (
	"github.com/go-redis/redis"
	"fmt"
)

func main() {
	var a = "abcd"
	fmt.Println(a[2:])
}

func bbb(client *redis.Client,i string){
	data := client.Get(i)
	//if err != nil {
	//	panic(err)
	//}
	fmt.Println(data)

}
