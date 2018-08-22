package main

import (
	"WebSocket/HttpController"
	"fmt"
	"strconv"
	"WebSocket/Redis"
)

func main() {
	var i int
	for i=1;i<11538;i++ {
		key := strconv.Itoa(i)
		Token := HttpController.GetToken("192.168.31.232:9501",key)
		fmt.Println(Token)
		Redis.InsertToken(key,Token,key)
	}

}
