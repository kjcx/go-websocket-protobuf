package Req

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
	"WebSocket/Protobuf"
)

type UseCompost struct {
	Id int32
}
//施肥 2011
func UseCompostReq(Uid int32,UseCompost *UseCompost) []byte {
	UseCompostReq := &AutoMsg.UseCompostReq{
		Id: UseCompost.Id,
	}
	Data,_ := proto.Marshal(UseCompostReq)
	Param,_ :=json.Marshal(UseCompostReq)
	log := Mgo.Log{
		Uid:   Uid,
		MsgId: 2011,
		Name:  "施肥",
		Param: string(Param),
		Date:  time.Now(),
		Msg:   "无",
	}
	log.InsertLog()
	return SendRev.Send(2011,Data)
}
