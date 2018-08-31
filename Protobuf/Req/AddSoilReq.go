package Req

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
	"WebSocket/Protobuf"
)
//庄园购买地块请求1078
func AddSoilReq(Uid int32,AddSoil *AutoMsg.AddSoilReq) []byte {
	AddSoilReq := &AutoMsg.AddSoilReq{
		Id: AddSoil.Id,
	}
	Data,_ := proto.Marshal(AddSoilReq)
	Param,_:= json.Marshal(AddSoilReq)
	log := Mgo.Log{
		Uid:   Uid,
		MsgId: 1078,
		Name:  "庄园购买地块请求",
		Param: string(Param),
		Date:  time.Now(),
		Msg:   "无",
	}
	log.InsertLog()
	return SendRev.Send(1078,Data)
}
