package Common

import (
	"math/rand"
	"time"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"log"
	"io"
	//"github.com/julienschmidt/httprouter"
	//"strconv"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"strconv"
	"flag"
)


var Addr = flag.String("ws", "192.168.31.232:9501", "http service address")
var Httpurl = flag.String("http", "192.168.31.232:9501", "http service address")

var Start = flag.Int("start", 14, "key start")
var End = flag.Int("end", 15, "key end")

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
func HttpGet(key string) string {
	//time.Sleep(time.Second * 1)
	fmt.Println("key=>", key)
	fmt.Println("httpurl=>", *Httpurl)
	fmt.Println(time.Now())
	url := fmt.Sprintf("http://%s%s%s", *Httpurl, "/Account/CKlogin?key=", key)
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
//http参数
type Param struct {
	Uid int32
	Body []byte
}
type HttpParam struct {
	R *http.Request
	Ps httprouter.Params
}
//解析请求的参数
func (HttpParam *HttpParam) GetParam() *Param{
	Name := HttpParam.Ps.ByName("name")
	Nameint,_ := strconv.Atoi(Name)
	Uid := int32(Nameint)
	Param := &Param{}
	Param.Uid = Uid
	Body := Decode(HttpParam.R.Body)
	Param.Body = Body
	return Param
}
//解析body
func Decode(Body io.Reader) []byte {
	body, err := ioutil.ReadAll(Body)
	if err != nil {
		fmt.Printf("read body err, %v\n", err)
	}
	return body
}
//格式化json打印
func JsonFmt(st interface{}) string{
	str,_ :=json.MarshalIndent(st,""," ")
	return string(str)
}


//添加道具
func AddItemSend(Uid int32,ItemId int32,Count int64) bool{
	url := fmt.Sprintf("http://%s/Gm/addItem?uid=%s&itemId=%s&num=%s", *Httpurl,strconv.Itoa(int(Uid)),strconv.Itoa(int(ItemId)),strconv.Itoa(int(Count)) )
	//url := "http://" + httpurl + "/Account/CKlogin?key=" + key
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		// handle error
		log.Fatal("handle error http")
	}
	defer resp.Body.Close()
	ioutil.ReadAll(resp.Body)

	if err != nil {
		// handle error
		log.Fatal("handle error Body")
		return false
	}else {
		//fmt.Println(string(body))
		return true

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
