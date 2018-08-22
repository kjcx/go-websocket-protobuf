package main

import (
	"WebSocket/Mgo"
	"time"
	"fmt"
)

func main() {
	log := &Mgo.Log{
		Uid:   14,
		MsgId: 1004,
		Name:  "使用道具",
		Param: "参数",
		Date:  time.Now(),
		Msg:   "无",
	}
	mongo := Mgo.Mongo()
	Mgo.InsertLog(mongo,log)

	var ch1 = make(chan int)
	var ch2 = make(chan int)
	go f1(ch1)
	go f2(ch2)
	for  {
		select {
		case <-ch1:
			fmt.Println("The first case is selected.")
		case <-ch2:
			fmt.Println("The second case is selected.")
		}
	}

}
