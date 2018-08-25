package Ws

import (
	"golang.org/x/net/websocket"
	"fmt"
	"strconv"
	"time"
	"github.com/golang/protobuf/proto/test_proto"
	"log"
	"WebSocket/Protobuf/Req"
	"WebSocket/Protobuf"
	"WebSocket/Protobuf/Result"
	"WebSocket/Common"
	"WebSocket/Model"
)

var arr   [30000]*websocket.Conn
func GetWebSocket() [30000]*websocket.Conn{
	return arr
}
//初始化ws
func Init(addr string,httpurl string,start int,end int) {
	fmt.Println("websocket：", addr)
	fmt.Println("http：", httpurl)
	fmt.Println("http：", start)
	fmt.Println("http：", end)
	var ws *websocket.Conn
	var i int
	for i = start; i < end; i++ {
		arr[i] = ws
		Connect(i,addr,httpurl)//发起连接
		go ForRead(i)//协程读取内容
	}
}
//发起连接
func Connect(i int,addr string,httpurl string) *websocket.Conn {

	fmt.Println("ccccc")
	fmt.Println(addr,httpurl)
	url := "ws://" + addr
	origin := fmt.Sprintf("Uid:%d",i)
	fmt.Println(strconv.Itoa(i))
	Token := Common.HttpGet(httpurl,strconv.Itoa(i))
	var str = Req.ConnectingReq(1004, Token)
	fmt.Println(str)
	var err error
	arr[i], err = websocket.Dial(url, "", origin)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("arr=》")
	fmt.Println(arr[i].Config().Origin)

	go Send(arr[i], str)
	return arr[i]
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
		Origin := arr[i].Config().Origin.String()
		Uidstr,_ := strconv.Atoi(Origin[4:])
		Uid := int32(Uidstr)
		res := SendRev.Rev(Uid,msg1[:n])
		go SwitchMsg(arr[i],Uid, res)

	}
}

func Send(conn *websocket.Conn, data []byte) {
	websocket.Message.Send(conn, data)
}
func timeWriterSend(conn *websocket.Conn, data []byte) {
	for {
		time.Sleep(time.Second*5)
		//time.Sleep(time.Microsecond * 1000000)
		fmt.Println("发送时间:", time.Now())
		websocket.Message.Send(conn, data)
	}
}

func SwitchMsg(ws *websocket.Conn, Uid int32,res *AutoMsg.MsgBaseSend) {
	switch res.GetMsgID() {
	case 1057:
		str := Result.ConnectingResult(res.GetData())
		RoleList := str.GetRoleLists()[0]
		fmt.Println("角色id", RoleList.GetRoleId())
		if RoleList.GetRoleId() != "" {
			str_JoinGamReq := Req.JoinGameReq(1012)
			go Send(ws, str_JoinGamReq)
		} else {
			//创建角色请求
			var str_role = Req.CreateRoleReq(1007)
			go Send(ws, str_role)
		}
		break
	case 1066:
		//Redis.Insert(res.GetData())
		strJoin := Result.JoinGameResult(res.GetData())
		fmt.Println(strJoin.GetMission())
		//测试请求 SignReq
		data := Req.SignReq()
		go Send(ws, data)
		str_1106 := Req.ShopAllReq(1106)
		go timeWriterSend(ws, str_1106)
		var Name = Common.GetRandomString(3)
		var str = Req.CreateCompanyReq(Name)
		go Send(ws, str)
		break
	case 1145:
		str_1145 := Result.ShopAllResult(res.GetData())
		fmt.Println(str_1145)
	case 1168:
		fmt.Println("MsgId:1168")
		msg := Result.SignResult(res.GetData())
		_,Day :=  Model.Sign(msg)
		if Day >0  {
			str := Model.DaySign(Day)
			go Send(ws,str)
		}
		break
	case 1078:
		fmt.Println("使用道具返回")
		msg := Result.UseItemResult(Uid,res.GetData())
		fmt.Println(msg)
	case 1059:
		fmt.Println("创建公司返回")
		Result.CreateCompanyResult(res.GetData())
	case 1058:
		Result.CreateBuildResult(res.GetData())
	case 1109:
		fmt.Println("刷新商城商品返回")
		Result.RefDropShopResult(res.GetData())
	case 1069:
		fmt.Println("出售道具返回")
		Result.SellItemResult(Uid,res.GetData())
	case 1055:
		fmt.Println("改变时装返回")
		Result.ChangeAvatarResult(Uid,res.GetData())
	case 1023:
		fmt.Println("购买时装返回")
		Result.ModelClothesResult(Uid,res.GetData())
	case 1141:
		fmt.Println("更改头像返回")
		Result.UpdateRoleInfoIconResult(Uid,res.GetData())
	case 1016:
		fmt.Println("搜索好友返回")
		Result.FriendSearchResult(Uid,res.GetData())
	case 1011:
		fmt.Println("申请好友返回")
		Result.FriendApplyResult(Uid,res.GetData())
	case 1013:
		fmt.Println("批量申请好友返回")
		Result.FriendAddResult(Uid,res.GetData())
	case 1012:
		fmt.Println("删除好友返回")
		Result.FriendRemoveResult(Uid,res.GetData())
	default:
		//go Test(ws)
	}
}