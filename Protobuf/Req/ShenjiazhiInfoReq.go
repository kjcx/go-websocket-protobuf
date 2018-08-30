package Req

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
	"WebSocket/Protobuf"
)
//身价值明细请求请求 2049
func ShenjiazhiInfoReq(Uid int32) []byte {
	ShenjiazhiInfoReq := &AutoMsg.ShenjiazhiInfoReq{}
	proto.Marshal(ShenjiazhiInfoReq)
	Param,_ := json.Marshal(ShenjiazhiInfoReq)
	log := Mgo.Log{
		Uid:   Uid,
		MsgId: 2049,
		Name:  "身价值明细请求请求",
		Param: string(Param),
		Date:  time.Now(),
		Msg:   "无",
	}
	log.InsertLog()
	return SendRev.Send(2049,Data)
}
