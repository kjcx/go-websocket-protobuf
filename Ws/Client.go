package Ws

import (
	"golang.org/x/net/websocket"
	"fmt"
	"strconv"
	"time"
	"github.com/golang/protobuf/proto/test_proto"
	"WebSocket/Protobuf/Req"
	"WebSocket/Protobuf"
	"WebSocket/Common"
	"WebSocket/Protobuf/Result"
	"encoding/json"
	"WebSocket/Model"
	"log"
)

var arr   [30000]*websocket.Conn
func GetWebSocket() [30000]*websocket.Conn{
	return arr
}

type SendChan struct {
	Uid int32
	Data []byte
}
var chan_req = make(chan SendChan)
var addr = *Common.Addr
var httpurl = *Common.Httpurl
var start = *Common.Start
var end = *Common.End
//初始化ws
func Init() {
	fmt.Println("websocket：", Common.Addr)
	fmt.Println("http：", Common.Httpurl)
	fmt.Println("http：", Common.Start)
	fmt.Println("http：", Common.End)
	var ws *websocket.Conn
	var i int
	for i = start; i < end; i++ {
		arr[i] = ws
		Connect(i)//发起连接
		go ForRead(i)//协程读取内容
	}
	go ChanMsgRead(chan_req)
}
//发起连接
func Connect(i int) *websocket.Conn {

	fmt.Println("ccccc")
	fmt.Println(addr,httpurl)
	url := "ws://" + addr
	origin := fmt.Sprintf("Uid:%d",i)
	fmt.Println(strconv.Itoa(i))
	Token := Common.HttpGet(strconv.Itoa(i))
	if Token != "" {
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

		go HeartbeatSend(arr[i],int32(i))//发送心跳
		return arr[i]
	}else {
		fmt.Println("获取token失败")
		return arr[i]
	}

}


func ForRead(i int) {
	for {
		var msg1 = make([]byte, 5120)
		var n int
		var err error
		if n, err = arr[i].Read(msg1); err != nil {
			fmt.Println("接收信息出错",err,msg1)

			if arr[i].IsClientConn() {
				log.Fatal("连接已断开",err.Error())
			}
		}

		fmt.Printf("Received:", n ,msg1)
		Origin := arr[i].Config().Origin.String()
		Uidstr,_ := strconv.Atoi(Origin[4:])
		Uid := int32(Uidstr)
		fmt.Println("uid:",Uid)
		res := SendRev.Rev(Uid,msg1[:n])
		go SwitchMsg(arr[i],Uid, res)


	}



}

func Send(conn *websocket.Conn, data []byte) {
	websocket.Message.Send(conn, data)
}
func timeWriterSend(conn *websocket.Conn, data []byte) {
	for {
		time.Sleep(time.Second*1)
		//time.Sleep(time.Microsecond * 1000000)
		fmt.Println(conn.Config().Origin,"发送时间:", time.Now())
		websocket.Message.Send(conn, data)
	}
}
//心跳每30秒发一次
func HeartbeatSend(conn *websocket.Conn,Uid int32) {
	for {
		time.Sleep(time.Second*30)
		//time.Sleep(time.Microsecond * 1000000)
		data := Req.HeartbeatReq(Uid)
		go Send(conn,data)
	}
}
//写入信息
func ChanMsgWrite(Data SendChan){
	chan_req <- Data
}

func ChanMsgRead(chan_req chan SendChan){
	for  {
		select {
		case msg :=<-chan_req:
			go Send(arr[msg.Uid],msg.Data)
			//执行websocket发送
			fmt.Println("执行chan接口",msg.Uid)
		}
	}
}
func SwitchMsg(ws *websocket.Conn, Uid int32,res *AutoMsg.MsgBaseSend) {
	if res.Result  == 0 {
		switch res.GetMsgID() {
		case 1057:
			str := Result.ConnectingResult(res.GetData())
			RoleList := str.GetRoleLists()[0]
			fmt.Println("角色id", RoleList.GetRoleId())
			if RoleList.GetRoleId() != "" {
				str_JoinGamReq := Req.JoinGameReq(Uid)
				go Send(ws, str_JoinGamReq)
			} else {
				//创建角色请求
				var str_role = Req.CreateRoleReq(Uid,1007)
				go Send(ws, str_role)
			}
			break
		case 1060:
			str_JoinGamReq := Req.JoinGameReq(Uid)
			go Send(ws, str_JoinGamReq)
		case 1066:
			//Redis.Insert(res.GetData())
			strJoin := Result.JoinGameResult(Uid,res.GetData())
			d,_:= json.MarshalIndent(strJoin,""," ")
			fmt.Println(string(d))
			//测试请求 SignReq
			data := Req.SignReq(Uid)
			go Send(ws, data)
			//str_1106 := Req.ShopAllReq(1106)
			//go timeWriterSend(ws, str_1106)
			//var Name = Common.GetRandomString(3)
			//var str = Req.CreateCompanyReq(Name)
			//go Send(ws, str)
			break
		case 1145:
			str_1145 := Result.ShopAllResult(Uid,res.GetData())
			fmt.Println(str_1145)
		case 1168:
			fmt.Println("MsgId:1168")
			msg := Result.SignResult(Uid,res.GetData())
			_,Day :=  Model.Sign(msg)
			if Day >0  {
				str := Model.DaySign(Uid,Day)
				go Send(ws,str)
			}
			break
		case 1078:
			fmt.Println("使用道具返回",time.Now())
			msg := Result.UseItemResult(Uid,res.GetData())
			fmt.Println(msg)
		case 1059:
			fmt.Println("创建公司返回",time.Now())
			Result.CreateCompanyResult(Uid,res.GetData())
		case 1058:
			Result.CreateBuildResult(Uid,res.GetData())
		case 1109:
			fmt.Println("刷新商城商品返回",time.Now())
			Result.RefDropShopResult(Uid,res.GetData())
		case 1069:
			fmt.Println("出售道具返回",time.Now())
			Result.SellItemResult(Uid,res.GetData())
		case 1055:
			fmt.Println("改变时装返回",time.Now())
			Result.ChangeAvatarResult(Uid,res.GetData())
		case 1142:
			fmt.Println("修改角色名称返回",time.Now())
			Result.UpdateRoleInfoNameResult(Uid,res.GetData())
		case 1023:
			fmt.Println("购买时装返回",time.Now())
			Result.ModelClothesResult(Uid,res.GetData())
		case 1141:
			fmt.Println("更改头像返回",time.Now())
			Result.UpdateRoleInfoIconResult(Uid,res.GetData())
		case 1016:
			fmt.Println("搜索好友返回",time.Now())
			Result.FriendSearchResult(Uid,res.GetData())
		case 1011:
			fmt.Println("申请好友返回",time.Now())
			Result.FriendApplyResult(Uid,res.GetData())
		case 1013:
			fmt.Println("批量申请好友返回",time.Now())
			Result.FriendAddResult(Uid,res.GetData())
		case 1012:
			fmt.Println("删除好友返回",time.Now())
			Result.FriendRemoveResult(Uid,res.GetData())
		case 1133:
			fmt.Println("提升好感度返回",time.Now())
			Result.AddNpcRelationAdvanceResult(Uid,res.GetData())
		case 2026:
			fmt.Println("解锁npc返回",time.Now())
			Result.UnlockNpcResult(Uid,res.GetData())
		case 1037:
			fmt.Println("加载npc好感度返回",time.Now())
			Result.NpcFavorabilityResult(Uid,res.GetData())
		case 1117:
			fmt.Println("招聘返回",time.Now())
			Result.RefStaffResult(Uid,res.GetData())
		case 1126:
			fmt.Println("调入调出员工返回",time.Now())
			Result.ComeOutEmployeeResult(Uid,res.GetData())
		case 1118:
			fmt.Println("加载员工返回",time.Now())
			Result.LoadStaffResult(Uid,res.GetData())
		case 2038:
			fmt.Println("赠送礼物返回",time.Now())
			Result.GiveGiftResult(Uid,res.GetData())
		case 2056:
			fmt.Println("礼物列表",time.Now())
			Result.GiveListResult(Uid,res.GetData())
		case 1176:
			fmt.Println("人才市场列表返回",time.Now())
			Result.GetTalentListResult(Uid,res.GetData())
		case 1004:
			fmt.Println("升级店铺返回",time.Now())
			Result.BuildLvUpResult(Uid,res.GetData())
		case 1063:
			fmt.Println("出售店铺返回",time.Now())
			Result.DestoryBuildResult(Uid,res.GetData())
		case 1178:
			fmt.Println("解雇经理返回",time.Now())
			Result.TalentFireResult(Uid,res.GetData())
		case 1182:
			fmt.Println("刷新人才市场",time.Now())
			Result.TalentRefreshResult(Uid,res.GetData())
		case 1121:
			fmt.Println("返回时间",time.Now())
			fmt.Println("培训员工返回")
			Result.CultivateEmployeeResult(Uid,res.GetData())
		case 1177:
			fmt.Println("雇佣经理",time.Now())
			Result.TalentHireResult(Uid,res.GetData())
		case 1082:
			fmt.Println("庄园请求",time.Now())
			Result.RequestManorResult(Uid,res.GetData())
		case 1201:
			fmt.Println("拜访记录返回",time.Now())
			Result.ManorVisitInfoResult(Uid,res.GetData())
		case 1029:
			fmt.Println("任务返回",time.Now())
			Result.MissionDataResult(Uid,res.GetData())
		case 2006:
			fmt.Println("今日竞拍列表返回",time.Now())
			Result.GetAuctionLandResult(Uid,res.GetData())
		case 2008:
			fmt.Println("竞拍请求返回",time.Now())
			Result.AuctionLandResult(Uid,res.GetData())
		case 2010:
			fmt.Println("自己竞拍列表返回",time.Now())
			Result.MyLandInfoResult(Uid,res.GetData())
		case 1030:
			fmt.Println("水果机列表返回",time.Now())
			Result.FruitsDataResult(Uid,res.GetData())
		case 1035:
			fmt.Println("水果机请求返回",time.Now())
			Result.RaffleFruitsResult(Uid,res.GetData())
		case 1018:
			fmt.Println("拉黑好友返回",time.Now())
			Result.FriendAddBlackResult(Uid,res.GetData())
		case 1114:
			fmt.Println("阅读邮件返回",time.Now())
			Result.ReadMailResult(Uid,res.GetData())
		case 1095:
			fmt.Println("一键获取邮件物品返回",time.Now())
			Result.GetMailItemsResult(Uid,res.GetData())
		case 1113:
			fmt.Println("批量删除邮件返回",time.Now())
			Result.DelMailsResult(Uid,res.GetData())
		case 1019:
			fmt.Println("移除黑名单返回",time.Now())
			Result.FriendRemoveBlackResult(Uid,res.GetData())
		case 1134:
			fmt.Println("居民委托任务返回",time.Now())
			Result.ResidentDelegateResult(Uid,res.GetData())
		case 1091:
			fmt.Println("偷菜返回",time.Now())
			Result.StealSemenResult(Uid,res.GetData())
		case 2014:
			fmt.Println("随机请求庄园返回")
			Result.RandManorResult(Uid,res.GetData())
		case 1110:
			fmt.Println("庄园购买地块返回")
			Result.AddSoilResult(Uid,res.GetData())
		case 1064:
			fmt.Println("大地图",time.Now())
			Result.GetMapResult(Uid,res.GetData())
		case 1192:
			fmt.Println("自己公有区店铺返回",time.Now())
			Result.PublicRoleShopResult(Uid,res.GetData())
		default:
			//go Test(ws)
		}
	}


}