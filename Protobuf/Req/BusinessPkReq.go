package Req

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
	"WebSocket/Protobuf"
)
//商战请求 2043
func BusinessPkReq(Uid int32) []byte {
	BusinessPkReq := &AutoMsg.BusinessPkReq{}
	Data,_ := proto.Marshal(BusinessPkReq)
	Param,_ := json.Marshal(BusinessPkReq)
	log := Mgo.Log{
		Uid:   Uid,
		MsgId: 2043,
		Name:  "商战请求",
		Param: string(Param),
		Date:  time.Now(),
		Msg:   "无",
	}
	log.InsertLog()
	return SendRev.Send(2043,Data)
}
