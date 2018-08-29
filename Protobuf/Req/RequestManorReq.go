package Req

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"WebSocket/Protobuf"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
)
type RequestManor struct {
	UserId	string
}
//1053 请求庄园
func RequestManorReq(Uid int32,RequestManor *RequestManor) []byte{
	RequestManorReq := &AutoMsg.RequestManorReq{
		UserId: RequestManor.UserId,
	}
	Data,_ := proto.Marshal(RequestManorReq)
	Param,_ := json.Marshal(RequestManorReq)
	log := Mgo.Log{
		Uid:   Uid,
		MsgId: 1053,
		Name:  "请求庄园",
		Param: string(Param),
		Date:  time.Now(),
		Msg:   "无",
	}
	log.InsertLog()
	return SendRev.Send(1053,Data)
}