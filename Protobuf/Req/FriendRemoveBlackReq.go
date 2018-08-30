package Req

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
	"WebSocket/Protobuf"
)

type FriendRemoveBlack struct {
	RoleId string
}
// 移除黑名单好友 1027
func FriendRemoveBlackReq(Uid int32,FriendRemoveBlack *FriendRemoveBlack) []byte{
	FriendRemoveBlackReq := &AutoMsg.FriendRemoveBlackReq{
		RoleId: FriendRemoveBlack.RoleId,
	}
	Data,_ := proto.Marshal(FriendRemoveBlackReq)
	Param,_ := json.Marshal(FriendRemoveBlackReq)
	log := Mgo.Log{
		Uid:   Uid,
		MsgId: 1027,
		Name:  "移除黑名单好友",
		Param: string(Param),
		Date:  time.Now(),
		Msg:   "无",
	}
	log.InsertLog()
	return SendRev.Send(1027,Data)
}
