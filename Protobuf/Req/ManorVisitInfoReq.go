package Req

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
	"WebSocket/Protobuf"
)
//拜访庄园1148
func ManorVisitInfoReq(Uid int32) []byte{
	ManorVisitInfoReq := &AutoMsg.ManorVisitInfoReq{}
	Data,_ := proto.Marshal(ManorVisitInfoReq)
	Param,_ := json.Marshal(ManorVisitInfoReq)
	log := Mgo.Log{
		Uid:   Uid,
		MsgId: 1148,
		Name:  "庄园拜访记录",
		Param: string(Param),
		Date:  time.Now(),
		Msg:   "无",
	}
	log.InsertLog()
	return SendRev.Send(1148,Data)
}
