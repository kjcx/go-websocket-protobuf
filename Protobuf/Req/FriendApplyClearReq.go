package Req

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"WebSocket/Protobuf"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
)

//1025
type FriendApplyClear struct {
	RoleIds []string
}
//拒绝用户申请 1025
func FriendApplyClearReq(Uid int32,FriendApplyClear *FriendApplyClear) []byte {
	FriendApplyClearReq := &AutoMsg.FriendApplyClearReq{
		RoleIds: FriendApplyClear.RoleIds,
	}
	Data,_ := proto.Marshal(FriendApplyClearReq)
	Param,_ := json.Marshal(FriendApplyClearReq)
	//插入日志
	log := &Mgo.Log{
		Uid:Uid,
		MsgId: 1025,
		Name: "拒绝用户申请请求",
		Param: string(Param),
		Date: time.Now(),
		Msg: "",
	}
	log.InsertLog()
	return SendRev.Send(1025,Data)
}
