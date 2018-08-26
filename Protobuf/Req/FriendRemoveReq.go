package Req

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"WebSocket/Protobuf"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
)

type FriendRemove struct {
	RoleId string
}
//删除好友请求 1023
func FriendRemoveReq(Uid int32,FriendRemove *FriendRemove) []byte{
	FriendRemoveReq := &AutoMsg.FriendRemoveReq{
		RoleId: FriendRemove.RoleId,
	}
	Data,_ := proto.Marshal(FriendRemoveReq)
	Param,_ := json.Marshal(FriendRemoveReq)
	//插入日志
	log := &Mgo.Log{
		Uid:Uid,
		MsgId: 1023,
		Name: "删除好友请求请求",
		Param: string(Param),
		Date: time.Now(),
		Msg: "",
	}
	log.InsertLog()
	return SendRev.Send(1023,Data)
}
