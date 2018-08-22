package main

import (
	"time"
	"fmt"
)

func main() {
	t := time.NewTicker(time.Second)
	for v := range t.C {
		fmt.Println("hello, ", v)
	}
}
