package main

import (
	"fmt"
	"time"
)
func main(){
go	Test()
	fmt.Println("1")
	time.Sleep(time.Second*7)
}
func Test() {
	fmt.Println(111)
	var ch1 = make(chan int)
	var ch2 = make(chan string)
	go f2(ch2)
	for  {
		select {
		case <-ch1:
			//执行websocket发送
			fmt.Println("执行使用道具接口")
			fmt.Println("The first case is selected.")
		case v:= <-ch2:
			fmt.Println(v)
			fmt.Println("The second case is selected.")
		}
	}

}

func f1(ch chan int) {
	time.Sleep(time.Second * 5)
	ch <- 1
}

func f2(ch chan string) {
	//time.Sleep(time.Second * 5)
	ch <- `{"a":1,"b":"2}`
}
