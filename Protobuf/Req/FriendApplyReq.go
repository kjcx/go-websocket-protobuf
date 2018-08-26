package Req

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"WebSocket/Protobuf"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
)
type FriendApply struct {
	RoleId string
}
//1021 申请加好友
func FriendApplyReq(Uid int32,FriendApply *FriendApply) []byte{
	FriendApplyReq := &AutoMsg.FriendApplyReq{
		RoleId: FriendApply.RoleId,
	}
	Data,_ := proto.Marshal(FriendApplyReq)
	Param,_ := json.Marshal(FriendApplyReq)
	//插入日志
	log := &Mgo.Log{
		Uid:Uid,
		MsgId: 1021,
		Name: "申请加好友请求",
		Param: string(Param),
		Date: time.Now(),
		Msg: "",
	}
	log.InsertLog()
	return SendRev.Send(1021,Data)
}
