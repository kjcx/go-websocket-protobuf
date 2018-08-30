package Req

import (
	"github.com/golang/protobuf/proto/test_proto"
	"encoding/json"
	"github.com/golang/protobuf/proto"
	"WebSocket/Protobuf"
	"WebSocket/Mgo"
	"time"
)

type FriendAddBlack struct {
	RoleId string
}
//好友拉入黑名单 1026
func FriendAddBlackReq(Uid int32,FriendAddBlack *FriendAddBlack)[]byte  {
	FriendAddBlackReq := &AutoMsg.FriendAddBlackReq{
		RoleId: FriendAddBlack.RoleId,
	}
	Data,_ := proto.Marshal(FriendAddBlackReq)
	Param,_ := json.Marshal(FriendAddBlackReq)
	log := Mgo.Log{
		Uid:   Uid,
		MsgId: 1026,
		Name:  "好友拉入黑名单",
		Param: string(Param),
		Date:  time.Now(),
		Msg:   "无",
	}
	log.InsertLog()
	return SendRev.Send(1026,Data)
}
