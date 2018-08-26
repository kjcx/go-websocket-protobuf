package Req

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"WebSocket/Protobuf"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
)

type ChangeAvatar struct {
	Id []int32
}
//改变时装请求 1002
func ChangeAvatarReq(Uid int32,ChangeAvatar *ChangeAvatar) []byte {
	ChangeAvatarReq := &AutoMsg.ChangeAvatarReq{
		Id: ChangeAvatar.Id,
	}
	Data,_ := proto.Marshal(ChangeAvatarReq)
	Param,_ := json.Marshal(ChangeAvatarReq)
	//插入日志
	log := &Mgo.Log{
		Uid:Uid,
		MsgId: 1002,
		Name: "改变时装请求请求",
		Param: string(Param),
		Date: time.Now(),
		Msg: "",
	}
	log.InsertLog()
	return SendRev.Send(1002,Data)
}
