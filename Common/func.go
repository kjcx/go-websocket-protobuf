package Common

import (
	"math/rand"
	"time"
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"log"
)

func main() {
	
}
func GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

type Token struct {
	Token string
}

type Result struct {
	Data *Token
	Msg  string
}
func HttpGet(httpurl string,key string) string {
	//time.Sleep(time.Second * 1)
	fmt.Println("key=>", key)
	fmt.Println("httpurl=>", httpurl)
	fmt.Println(time.Now())
	url := fmt.Sprintf("http://%s%s%s", httpurl, "/Account/CKlogin?key=", key)
	//url := "http://" + httpurl + "/Account/CKlogin?key=" + key
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		// handle error
		log.Fatal("handle error http")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		// handle error
		log.Fatal("handle error Body")
		return ""
	}else {
		//fmt.Println(string(body))

		L := Result{}
		err := json.Unmarshal([]byte(body), &L)
		if err!=nil{
			fmt.Println("err",err.Error())
			return ""
		}else{
			return L.Data.Token
		}

		//fmt.Println(L.Data.Token)
		//fmt.Println(L.Msg)
	}





}
//通道发送消息
func ReqChan(ch chan string,Str string) {
	ch <- Str
}

//func SendChan(ws *websocket.Conn,Str string) {
//	fmt.Println(111)
//	var ch1 = make(chan int)
//	var ch2 = make(chan string)
//	var chan_req = make(chan string)
//	go ReqChan(chan_req,Str)
//	for  {
//		select {
//		case str := <-chan_req:
//			fmt.Println(str)
//		case <-ch1:
//			//执行websocket发送
//			fmt.Println("执行使用道具接口")
//			str := Req.UseItemReq(&Req.UseItem{10110,1})
//			fmt.Println(str)
//			//go timeWriter(arr[14],str)
//			fmt.Println("The first case is selected.")
//		case <-ch2:
//			for n:= range ch2{
//				fmt.Println(n)
//			}
//			fmt.Println("The second case is selected.")
//		}
//	}
//
//}
