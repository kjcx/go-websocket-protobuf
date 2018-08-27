package Req

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"WebSocket/Protobuf"
	"WebSocket/Mgo"
	"time"
	"encoding/json"
)
//1130 人才市场列表请求
func GetTalentListReq(Uid int32) []byte{
	GetTalentListReq := &AutoMsg.GetTalentListReq{}
	Data,_ := proto.Marshal(GetTalentListReq)
	Param,_ := json.Marshal(GetTalentListReq)
	log := &Mgo.Log{
		Uid:   Uid,
		MsgId: 1130,
		Name:  "人才市场列表请求",
		Param: string(Param),
		Date:  time.Now(),
		Msg:   "无",
	}
	log.InsertLog()
	return SendRev.Send(1130,Data)
}