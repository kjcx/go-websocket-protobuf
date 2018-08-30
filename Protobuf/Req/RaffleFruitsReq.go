package Req

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
	"WebSocket/Protobuf"
)

type RaffleFruits struct {
	Id int32
}
//水果机请求
func RaffleFruitsReq(Uid int32,RaffleFruits *RaffleFruits) []byte{
	RaffleFruitsReq := &AutoMsg.RaffleFruitsReq{
		Id: RaffleFruits.Id,
	}
	Data,_ := proto.Marshal(RaffleFruitsReq)
	Param,_ := json.Marshal(RaffleFruitsReq)
	log := Mgo.Log{
		Uid:   Uid,
		MsgId: 1035,
		Name:  "水果机请求",
		Param: string(Param),
		Date:  time.Now(),
		Msg:   "无",
	}
	log.InsertLog()
	return SendRev.Send(1035,Data)
}
