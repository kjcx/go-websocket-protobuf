package Req

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
	"WebSocket/Protobuf"
)

type UpgradeLandLevel struct {
	Id int32
}
//升级土地请求 1125
func UpgradeLandLevelReq(Uid int32,UpgradeLandLevel *UpgradeLandLevel)[]byte  {
	UpgradeLandLevelReq:= &AutoMsg.UpgradeLandLevelReq{
		Id: UpgradeLandLevel.Id,
	}
	Data,_ := proto.Marshal(UpgradeLandLevelReq)
	Param,_ := json.Marshal(UpgradeLandLevelReq)
	log := Mgo.Log{
		Uid:   Uid,
		MsgId: 1125,
		Name:  "升级土地请求",
		Param: string(Param),
		Date:  time.Now(),
		Msg:   "无",
	}
	log.InsertLog()
	return SendRev.Send(1125,Data)
}
