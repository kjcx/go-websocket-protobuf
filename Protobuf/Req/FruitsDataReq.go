package Req

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"WebSocket/Protobuf"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
)
//水果机列表请求 1031
func FruitsDataReq(Uid int32) []byte{
	FruitsDataReq := &AutoMsg.FruitsDataReq{}
	Data,_ := proto.Marshal(FruitsDataReq)
	Param,_ := json.Marshal(FruitsDataReq)
	log := Mgo.Log{
		Uid:   Uid,
		MsgId: 1031,
		Name:  "水果机列表请求",
		Param: string(Param),
		Date:  time.Now(),
		Msg:   "无",
	}
	log.InsertLog()
	return SendRev.Send(1031,Data)
}
