package Req

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
	"WebSocket/Protobuf"
)
//随机请求庄园 2013
func RandManorReq(Uid int32) []byte{
	RandManorReq := &AutoMsg.RandManorReq{}
	Data,_ := proto.Marshal(RandManorReq)
	Param,_ := json.Marshal(RandManorReq)
	log := Mgo.Log{
		Uid:   Uid,
		MsgId: 2013,
		Name:  "随机请求庄园",
		Param: string(Param),
		Date:  time.Now(),
		Msg:   "无",
	}
	log.InsertLog()
	return SendRev.Send(2013,Data)
}
