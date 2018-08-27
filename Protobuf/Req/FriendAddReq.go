package Req

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"WebSocket/Protobuf"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
)

type FriendAdd struct {
	RoleIds []string
}
//1022 通过好友请求
func FriendAddReq(Uid int32,FriendAdd *FriendAdd) []byte{
	FriendAddReq := &AutoMsg.FriendAddReq{
		RoleIds: FriendAdd.RoleIds,
	}
	Data,_ := proto.Marshal(FriendAddReq)
	Param,_ := json.Marshal(FriendAddReq)
	//插入日志
	log := &Mgo.Log{
		Uid:Uid,
		MsgId: 1022,
		Name: "通过好友请求",
		Param: string(Param),
		Date: time.Now(),
		Msg: "",
	}
	log.InsertLog()
	return SendRev.Send(1022,Data)
}
