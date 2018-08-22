package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/proto/test_proto"
	"golang.org/x/net/websocket"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
	"WebSocket/Protobuf/Req"
	"WebSocket/Protobuf/Result"

	"WebSocket/Model"
	"WebSocket/Common"
	"WebSocket/Redis"
)

var addr = flag.String("ws", "slb.ckp520.cn:9501", "http service address")
var httpurl = flag.String("http", "slb.ckp520.cn:9501", "http service address")
var start = flag.Int("start", 3000, "key start")
var end = flag.Int("end", 3001, "key end")
var arr [30000]*websocket.Conn

//var ws1 *websocket.Conn
//var ws2 *websocket.Conn

func main() {
	flag.Parse()

	fmt.Println("websocket：", *addr)
	fmt.Println("http：", *httpurl)
	//arr[14] = ws1
	//arr[15] = ws2
	var ws1 *websocket.Conn
	var i int
	for i = *start; i < *end; i++ {
		arr[i] = ws1
		connect(i)

		//var msg1 [3072]byte
		go ForRead(i)
	}
	for {
		time.Sleep(time.Second * 1)
	}

}
func ForRead(i int) {
	for {
		var msg1 = make([]byte, 5120)
		var n int
		var err error
		if n, err = arr[i].Read(msg1); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Received:", msg1, n)
		res := Rev(msg1[:n])
		go SwitchMsg(arr[i], res)

	}
}

type Channel struct {
	conn *websocket.Conn
	// WebSocket connection.
}

func Print(str interface{}) {
	var c1, c2, c3 chan int
	var i1, i2 int
	select {
	case i1 = <-c1:
		fmt.Printf("received ", i1, " from c1\n")
	case c2 <- i2:
		fmt.Printf("sent ", i2, " to c2\n")
	case i3, ok := (<-c3): // same as: i3, ok := <-c3
		if ok {
			fmt.Printf("received ", i3, " from c3\n")
		} else {
			fmt.Printf("c3 is closed\n")
		}
	default:
		fmt.Printf("no communication\n")
	}

}
func SwitchMsg(ws *websocket.Conn, res *AutoMsg.MsgBaseSend) {
	switch res.GetMsgID() {
	case 1057:
		str := ConnectingResult(res.GetData())
		RoleList := str.GetRoleLists()[0]
		fmt.Println("角色id", RoleList.GetRoleId())
		if RoleList.GetRoleId() != "" {
			str_JoinGamReq := JoinGamReq(1012)
			go timeWriter(ws, str_JoinGamReq)
		} else {
			//创建角色请求
			var str_role = CreateRoleReq(1007)
			go timeWriter(ws, str_role)
		}
		break
	case 1066:
		Redis.Insert(res.GetData())
		strJoin := JoinGameResult(res.GetData())
		fmt.Println(strJoin.GetMission())
		//测试请求 SignReq
		data := Req.SignReq()
		go timeWriter(ws, data)
		//str_1106 := ShopAllReq(1106)
		//go timeWriterSend(ws, str_1106)
		var Name = Common.GetRandomString(3)
		var str = Req.CreateCompanyReq(Name)
		go timeWriter(ws, str)
		break
	case 1145:
		str_1145 := ShopAllResult(res.GetData())
		fmt.Println(str_1145)
	case 1168:
		fmt.Println("MsgId:1168")
		msg := Result.SignResult(res.GetData())
		_,Day :=  Model.Sign(msg)
		if Day >0  {
			str := Model.DaySign(Day)
			go timeWriter(ws,str)
		}
		break

	}
}

//发送消息
func Send(MsgID int32, Data []byte) []byte {

	MsgSend := &AutoMsg.MsgBaseRev{
		MsgId: MsgID,
		Data:  Data,
	}
	data, _ := proto.Marshal(MsgSend)
	return data
}
func Rev(Data []byte) *AutoMsg.MsgBaseSend {
	fmt.Println("data:")
	MsgRev := &AutoMsg.MsgBaseSend{}
	err := proto.Unmarshal(Data, MsgRev)
	if err != nil {
		fmt.Println("marshaling error1111: ", err)
	}
	MsgID := MsgRev.GetMsgID()
	if MsgRev.GetResult() >0{
		fmt.Println("有错误消息")
	}

	fmt.Println(MsgID)
	//MsgRev.GetData()
	return MsgRev
}

//发起连接

func ConnectingReq(MsgID int32, Token string) []byte {
	test := &AutoMsg.ConnectingReq{
		//Token: "token:5405ef6818c54e0ec07430fb1ad968ba",
		Token: Token,
	}
	data, err := proto.Marshal(test)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	fmt.Println(data)
	return Send(MsgID, data)

}

//解析登录返回
func ConnectingResult(Data []byte) *AutoMsg.ConnectingResult {
	Connecting := &AutoMsg.ConnectingResult{}
	err := proto.Unmarshal(Data, Connecting)
	if err != nil {
		log.Fatal("marshaling error: ")
	}
	return Connecting
}

func HttpGet(key string) string {
	//time.Sleep(time.Second * 1)
	fmt.Println("key=>", key)
	fmt.Println(time.Now())
	url := fmt.Sprintf("http://%s%s%s", *httpurl, "/Account/CKlogin?key=", key)
	//url := "http://" + httpurl + "/Account/CKlogin?key=" + key
	resp, err := http.Get(url)
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		// handle error
	}

	type Token struct {
		Token string
	}

	type Result struct {
		Data *Token
		Msg  string
	}

	//fmt.Println(string(body))

	L := Result{}
	json.Unmarshal([]byte(body), &L)
	fmt.Println(L.Data.Token)

	return L.Data.Token
	//fmt.Println(L.Data.Token)
	//fmt.Println(L.Msg)

}

//创建角色
func CreateRoleReq(MsgId int32) []byte {
	Name := "kjcx" + Common.GetRandomString(3)
	CreateRoleReq := &AutoMsg.CreateRoleReq{
		Name: Name,
		Sex:  1,
	}
	data, _ := proto.Marshal(CreateRoleReq)
	return Send(MsgId, data)
}

//解析角色
func CreateRoleResult(Data []byte) *AutoMsg.CreateRoleResult {
	fmt.Println("解析角色")
	CreateRole := &AutoMsg.CreateRoleResult{}
	err := proto.Unmarshal(Data, CreateRole)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	fmt.Println(CreateRole.GetRoleId())
	return CreateRole

}

//加入游戏
func JoinGamReq(MsgId int32) []byte {
	JoinGameReq := &AutoMsg.JoinGameReq{}
	data, _ := proto.Marshal(JoinGameReq)
	return Send(MsgId, data)
}

//加入游戏返回
func JoinGameResult(Data []byte) *AutoMsg.JoinGameResult {
	JoinGameResult := &AutoMsg.JoinGameResult{}
	err := proto.Unmarshal(Data, JoinGameResult)
	if err != nil {
		fmt.Println("marshaling error:1012 ", err)
	}
	return JoinGameResult
}

//商店请求
func ShopAllReq(MsgId int32) []byte {
	Shop := &AutoMsg.ShopAllReq{}
	data, err := proto.Marshal(Shop)
	if err != nil {
		fmt.Println("1")
	}
	return Send(MsgId, data)

}

//解析商店
func ShopAllResult(Data []byte) *AutoMsg.ShopAllResult {
	ShopAllResult := &AutoMsg.ShopAllResult{}
	err := proto.Unmarshal(Data, ShopAllResult)
	if err != nil {
		fmt.Println("解析ShopAllResult出错")
	}
	return ShopAllResult
}

func connect(i int) *websocket.Conn {
	flag.Parse()

	url := "ws://" + *addr + "/ws"
	origin := "test://1111111/"

	fmt.Println(strconv.Itoa(i))
	Token := HttpGet(strconv.Itoa(i))
	var str = ConnectingReq(1004, Token)
	fmt.Println(str)

	var err error
	arr[i], err = websocket.Dial(url, "", origin)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(arr[i])

	go timeWriter(arr[i], str)
	return arr[i]

	//return ws
	//for {
	//	//var msg []byte
	//
	//	var msg = make([]byte, 512)
	//	var n int
	//	if n, err = ws.Read(msg); err != nil {
	//		log.Fatal(err)
	//	}
	//	fmt.Printf("Received: ", msg[:n])
	//	//fmt.Printf("received: %s\n", msg)
	//}
}

func timeWriter(conn *websocket.Conn, data []byte) {
	websocket.Message.Send(conn, data)
	//for {
	//	time.Sleep(time.Second * 1)
	//	websocket.Message.Send(conn, data)
	//}
}
func timeWriterSend(conn *websocket.Conn, data []byte) {
	for {
		time.Sleep(time.Microsecond * 1000000)
		fmt.Println("发送时间:", time.Now())
		websocket.Message.Send(conn, data)
	}
}

