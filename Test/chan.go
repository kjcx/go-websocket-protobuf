package Test

//func Test(ws *websocket.Conn) {
//	fmt.Println(111)
//	var ch1 = make(chan int)
//	var ch2 = make(chan int)
//	go f1(ch1)
//	for  {
//		select {
//		case <-ch1:
//			//执行websocket发送
//			fmt.Println("执行使用道具接口")
//			Req.UseItemReq(&Req.UseItem{10110,1})
//			fmt.Println("The first case is selected.")
//		case <-ch2:
//			fmt.Println("The second case is selected.")
//		}
//	}
//
//}
//
//func f1(ch chan int) {
//	time.Sleep(time.Second * 5)
//	ch <- 1
//}
//func f2(ch chan int) {
//	time.Sleep(time.Second * 10)
//	ch <- 1
//}
